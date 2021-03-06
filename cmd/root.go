package shanty

import (
	"errors"
	"net"
	"os"
	"time"

	"github.com/prairir/Buoy/pkg/config"
	"github.com/prairir/Buoy/pkg/ethrouter"
	"github.com/prairir/Buoy/pkg/tun"
	"github.com/prairir/Buoy/pkg/tunrouter"
	"golang.org/x/sync/errgroup"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Buoy",
	Short: "A mesh VPN",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: Root,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	pFlags := rootCmd.PersistentFlags()
	pFlags.StringVar(&cfgFile, "config", "", "config file (default is /etc/buoy/buoy.yaml)")

	pFlags.BoolP("debug", "d", false, "debug mode")
	viper.BindPFlag("debug", pFlags.Lookup("debug"))

	pFlags.BoolP("log-cli", "l", false, "log output to cli or /var/log/buoy.log")
	viper.BindPFlag("log-cli", pFlags.Lookup("log-cli"))

	pFlags.StringP("fleet", "f", "", "CIDR of network to join")
	viper.BindPFlag("fleet", pFlags.Lookup("fleet"))

	pFlags.StringP("interface", "i", "by1", "The internal network interface name")
	viper.BindPFlag("interface", pFlags.Lookup("interface"))

	pFlags.StringP("password", "p", "", "Encryption password for VPN")
	viper.BindPFlag("passwordStr", pFlags.Lookup("password"))

	pFlags.String("listen-port", "31337", "Port to listen on")
	viper.BindPFlag("listen-port", pFlags.Lookup("listen-port"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// search config in defualt directory(/etc/buoy)
		// with name `buoy.yaml`
		viper.AddConfigPath("/etc/buoy/")
		viper.SetConfigType("yaml")
		viper.SetConfigName("buoy")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// set time stamp to unix epoch in ms
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Couldn't read in config")
	}

	// unmarshal viper to config.InitConfig
	err = config.UnmarshalConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Couldn't unmarshal config")
	}

	config.Config.FleetAddr, config.Config.FleetNet, err = net.ParseCIDR(config.Config.FleetStr)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Couldn't parse CIDR")
	}

	config.Config.Password = []byte(config.Config.PasswordStr)

	// set to info level first
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// if debug true then set to debug level
	if config.Config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// if LogCli true then log to stderr with pretty formatting
	if config.Config.LogCli {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123Z})
	} else {
		// this is thread safe because the underlying file write sys
		// calls are blocking and so they are thread safe
		f, err := os.OpenFile("/var/log/buoy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Couldn't open file /var/log/buoy.log")
		}

		log.Logger = log.Output(f)
	}

	if len(config.Config.FleetMates) < 1 {
		log.Fatal().
			Err(errors.New("Require at least 1 mate")).
			Msg("no mates given")
	}

	mates := map[string]net.UDPAddr{}
	for key, value := range config.Config.FleetMates {
		addr, err := net.ResolveUDPAddr("udp", value)
		if err != nil {
			log.Fatal().
				Str("address", value).
				Err(err).
				Msg("couldn't parse address")
		}
		mates[key] = *addr
	}

	tunrouter.FleetList = mates

	log.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())

	log.Debug().
		Bool("Debug", config.Config.Debug).
		Bool("Log-Cli", config.Config.LogCli).
		Str("Iname", config.Config.IName).
		Str("Password", config.Config.PasswordStr).
		Str("Listen-Port", config.Config.ListenPort).
		Str("Fleet Network", config.Config.FleetNet.String()).
		Str("Fleet Address", config.Config.FleetAddr.String()).
		Msg("Config")
}

func Root(cmd *cobra.Command, args []string) {
	// TODO look into creating interface in tunrouter
	inf, err := tun.New(config.Config.IName,
		config.Config.FleetAddr, config.Config.FleetNet)
	if err != nil {
		log.Fatal().Err(err).Msg("Tun Creation Error")
	}

	eg := new(errgroup.Group)

	eth2TunQ := make(chan []byte, 1)
	tun2EthQ := make(chan ethrouter.Packet, 1)

	eg.Go(func() error {
		return ethrouter.Run(eg, tun2EthQ, eth2TunQ)
	})

	eg.Go(func() error {
		return tunrouter.Run(eg, inf, tun2EthQ, eth2TunQ) //TODO verify order of egress and ingress
	})

	if err = eg.Wait(); err != nil {
		log.Fatal().Err(err).Msg("Buoy Run Failed")
	}
}
