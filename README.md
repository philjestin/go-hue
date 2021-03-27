# go-hue

A command-line tool for controlling Philips Hue smart lights that are connected to a Philips Hue Bridge via the CLI.

## Setup

```
$ mkdir config
$ touch ./config/config.json
$ go-hue config-set --ip <your-bridge-ip> --user-token <your-bridge-generated-token>
```

### Verify it set the config file

```
$ go-hue config-get
```

You should see the values you just set.

### Test it tout

```
$ go-hue lights
```

## Usage

```
This CLI tool helps you configure and manage Philips Hue Lights

Usage:
   go-hue <commands> {flags}
   go-hue <command> {flags}

Commands:
   config-get                    get your current go-hue configuration
   config-set                    configure the IP address and user token for authentication with the Philips Hue Bridge.
   get-groups                    displays detailed information about all lights connected to your Philips Hue Bridge.
   help                          displays usage information
   lights                        displays detailed information about all lights connected to your Philips Hue Bridge.
   set-groups                    Set values of a hue groups state
   version                       displays version number

Arguments:
   commands                      commands that the cli tool provide.

Flags:
   -h, --help                    displays usage information of the application or a command (default: false)
   -V, --verbose                 display log information (default: false)
   -v, --version                 displays version number (default: false)
```
