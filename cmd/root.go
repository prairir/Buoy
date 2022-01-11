package shanty

import (
	"os"
	"time"

	"github.com/prairir/Buoy/pkg/config"
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

	pFlags.BoolP("log-cli", "l", true, "log output to cli or /var/log/buoy.log")
	viper.BindPFlag("log-cli", pFlags.Lookup("log-cli"))

	pFlags.StringP("fleet", "f", "", "CIDR of network to join")
	viper.BindPFlag("fleet", pFlags.Lookup("fleet"))

	pFlags.StringP("interface", "i", "by1", "The internal network interface name")
	viper.BindPFlag("interface", pFlags.Lookup("interface"))

	pFlags.StringP("password", "p", "", "Encryption password for VPN")
	viper.BindPFlag("password", pFlags.Lookup("password"))
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

	log.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())

	// unmarshal viper to config.InitConfig
	err = config.UnmarshalInitConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Couldn't unmarshal config")
	}

	// set to info level first
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// if debug true then set to debug level
	if config.InitConfig.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// if LogCli true then log to stderr with pretty formatting
	if config.InitConfig.LogCli {
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
}

func Root(cmd *cobra.Command, args []string) {
	log.Info().Msg("HI")
}
