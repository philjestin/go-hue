# go-hue

![go-version](https://img.shields.io/github/go-mod/go-version/philjestin/go-hue?label=Go%20Version) &nbsp;
![release](https://github.com/philjestin/go-hue/workflows/release/badge.svg?style=flat-square)

A command-line tool for controlling Philips Hue smart lights that are connected to a Philips Hue Bridge via the CLI.

## Installation

```
$ GO111MODULE=on go get -u "github.com/philjestin/go-hue"
```

## Setup

```
$ mkdir config
$ touch ./config/config.json
$ go-hue config-set --ip <your-bridge-ip> --user-token <your-bridge-generated-token>
```

### Verify it set the config file

```
$ go-hue config-get
Hue Bridge IP: 0.0.0.0
Hue Bridge User Token: potatopotato
```

You should see the values you just set.

### Test it tout

```
$ go-hue lights
2021/03/27 14:35:31 {"1":{"state":{"on":false,"bri":254,"hue":8402,"sat":140,"effect":"none","xy":[0.4575,0.4099],"ct":366,"alert":"select","colormode":"xy","mode":"homeautomation","reachable":true},"swupdate":{"state":"noupdates","lastinstall":"2020-10-05T18:36:31"},"type":"Extended color light","name":"Desk light 1","modelid":"LCT024","manufacturername":"Signify Netherlands B.V.","productname":"Hue play","capabilities":{"certified":true,"control":{"mindimlevel":100,"maxlumen":540,"colorgamuttype":"C","colorgamut":[[0.6915,0.3083],[0.1700,0.7000],[0.1532,0.0475]],"ct":{"min":153,"max":500}},"streaming":{"renderer":true,"proxy":true}},"config":{"archetype":"hueplay","function":"decorative","direction":"upwards","startup":{"mode":"safety","configured":true}},"uniqueid":"00:17:88:01:04:fa:03:c0-0b","swversion":"1.50.2_r30933","swconfigid":"949259E6","productid":"3241-3127-7871-LS00"},"2":{"state":{"on":false,"bri":254,"hue":56016,"sat":102,"effect":"none","xy":[0.3937,0.2952],"ct":266,"alert":"select","colormode":"xy","mode":"homeautomation","reachable":true},"swupdate":{"state":"noupdates","lastinstall":"2020-10-04T16:14:30"},"type":"Extended color light","name":"Desk light 2","modelid":"LCT024","manufacturername":"Signify Netherlands B.V.","productname":"Hue play","capabilities":{"certified":true,"control":{"mindimlevel":100,"maxlumen":540,"colorgamuttype":"C","colorgamut":[[0.6915,0.3083],[0.1700,0.7000],[0.1532,0.0475]],"ct":{"min":153,"max":500}},"streaming":{"renderer":true,"proxy":true}},"config":{"archetype":"hueplay","function":"decorative","direction":"upwards","startup":{"mode":"safety","configured":true}},"uniqueid":"00:17:88:01:04:fa:1e:0e-0b","swversion":"1.50.2_r30933","swconfigid":"949259E6","productid":"3241-3127-7871-LS00"},"3":{"state":{"on":false,"bri":254,"hue":8402,"sat":140,"effect":"none","xy":[0.4575,0.4099],"ct":366,"alert":"select","colormode":"xy","mode":"homeautomation","reachable":true},"swupdate":{"state":"noupdates","lastinstall":"2020-10-05T18:36:41"},"type":"Extended color light","name":"Tv light 1","modelid":"LCT024","manufacturername":"Signify Netherlands B.V.","productname":"Hue play","capabilities":{"certified":true,"control":{"mindimlevel":100,"maxlumen":540,"colorgamuttype":"C","colorgamut":[[0.6915,0.3083],[0.1700,0.7000],[0.1532,0.0475]],"ct":{"min":153,"max":500}},"streaming":{"renderer":true,"proxy":true}},"config":{"archetype":"hueplay","function":"decorative","direction":"upwards","startup":{"mode":"safety","configured":true}},"uniqueid":"00:17:88:01:04:fa:5a:41-0b","swversion":"1.50.2_r30933","swconfigid":"949259E6","productid":"3241-3127-7871-LS00"},"4":{"state":{"on":false,"bri":254,"hue":47625,"sat":48,"effect":"none","xy":[0.3403,0.3160],"ct":192,"alert":"select","colormode":"xy","mode":"homeautomation","reachable":true},"swupdate":{"state":"noupdates","lastinstall":"2020-10-05T18:36:36"},"type":"Extended color light","name":"Tv light 2","modelid":"LCT024","manufacturername":"Signify Netherlands B.V.","productname":"Hue play","capabilities":{"certified":true,"control":{"mindimlevel":100,"maxlumen":540,"colorgamuttype":"C","colorgamut":[[0.6915,0.3083],[0.1700,0.7000],[0.1532,0.0475]],"ct":{"min":153,"max":500}},"streaming":{"renderer":true,"proxy":true}},"config":{"archetype":"hueplay","function":"decorative","direction":"upwards","startup":{"mode":"safety","configured":true}},"uniqueid":"00:17:88:01:04:fa:5a:84-0b","swversion":"1.50.2_r30933","swconfigid":"949259E6","productid":"3241-3127-7871-LS00"}}
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
