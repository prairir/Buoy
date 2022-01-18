# Buoy

**WARNING: these instructions work best on a unix based system**

(Logo Pending)

We are using [water](https://github.com/songgao/water) because its really good.

## Configuration

Buoy uses YAML as the configuration langauge.

You can use `--config` to specify a configuration file. If you do not specify a location, it will look for `/etc/buoy/buoy.yaml`

All the configuration options can be specified with flags. You still need to point to a configuration file even if you specify with only flags.

We provide an example configuration called `buoy.example.yaml`

## Logging

By default, it logs to `/var/log/buoy.log`. It uses structured logging so it may be a bit unreadable.

To have nicer CLI output, use the `--log-cli` or `-l` flags

Both of these logging solutions are thread safe so you will drop any logs

## Building

To build, run

```sh
earthly +build
```

This will create a binary called `buoy`

## Run

Run with sudo permissions. It needs to create and manage network interfaces to run.

```sh
sudo ./buoy
```

run with printing

```sh
sudo ./buoy --log-cli
```

run with specified config

```sh
sudo ./buoy --config="buoy.yaml"
```

## Testing

We run all our tests inside earthly. This creates reproducible tests which are system agnostic.

To run all tests, run

```sh
earthly --allow-privileged +test
```

We need privileged container because some tests require it. Like the test `TestTunNew` requires a privileged container.

## Terminology

Buoy has a cool sea theme so we are gonna stick to it.

**NOTE: This will be in the format of `<what we call it> - <term name>**

- fleet - network
- sonar - distrubted peer discovery protocol
- shanty - cmd

## Contributing

You can read all about contributing to this project in `CONTRIBUTING.md`

## Architecture

You can read about it in `ARCHITECTURE.md`
