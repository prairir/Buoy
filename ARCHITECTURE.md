# Architecture
**This document describes the high-level architecture of this project**

If you want to familiarize yourself with the code base and *generally* how it works, this is a good place to be.

## High Level TLDR
`main.go` loads `cmd/root.go` which runs the function `Root`. `Root` starts all the magic.

## Code Map

#### Code Map Legend

`<file name>` for a file name

`<folder name>/` for a folder

`<folder name>/<file name>` for a file within a folder

### `main.go`

Main function where everything gets started. 

### `cmd/`

Command line interface package where command line interface parsing stuff lives. We call this package `shanty` like sea shanty.

### `cmd/root.go`

Root command file which gets instantiated by `main.go`
