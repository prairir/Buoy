# Buoy

**WARNING: these instructions work best on a unix based system**

(Logo Pending)

We are using [water](https://github.com/songgao/water) because its really good.

## Setup

Run this to run
``` go
sudo go run main.go
```

You **HAVE** to do this currently. We are working on a programatic way of bringing the interface up.
``` sh
sudo ip addr add <local ip range not in use> dev by1
sudo ip link set dev by1 up
```

## Terminology

Buoy has a cool sea theme so we are gonna stick to it.

**NOTE: This will be in the format of `<what we call it> - <term name>**
* fleet - network
* sonar - distrubted peer discovery protocol
* shanty - cmd

## Contributing
You can read all about contributing to this project in `CONTRIBUTING.md`

## Architecture
You can read about it in `ARCHITECTURE.md`
