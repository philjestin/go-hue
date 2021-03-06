package main

import (
	"fmt"
	"log"
	"net"

	"github.com/philjestin/go-hue/hue"
	"github.com/philjestin/go-hue/utils"
	"github.com/thatisuday/commando"
)

func main() {
	// Configure commando
	commando.
		SetExecutableName("go-hue").
		SetVersion("v1.0.0").
		SetDescription("This CLI tool helps you configure and manage Philips Hue Lights")

	// Configure the root-command
	// $ go-hue <category> --verbose | -V, --version| -V, --help | -h
	commando.
		Register(nil).
		AddArgument("commands", "commands that the cli tool provide.", "").
		AddFlag("verbose,V", "display log information ", commando.Bool, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// Print Arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// Configure the root-command
	// $ go-hue <category> --verbose | -V, --version| -V, --help | -h
	commando.
		Register("discover").
		SetShortDescription("discover the hue bridge on your network.").
		SetDescription("this command is used to discover the hue bridge on your network using Hue's nupnp service.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			bridges := utils.FindBridge()
			fmt.Printf("Your Philips Hue Bridge can be found at: %s\n", bridges[0].Host)
		})

	commando.
		Register("config-get").
		SetShortDescription("get your current go-hue configuration").
		SetDescription("this command is used to get your current go-hue configuration").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			printConfig()
		})

	commando.
		Register("config-set").
		SetShortDescription("configure the IP address and user token for authentication with the Philips Hue Bridge.").
		SetDescription("this command is used to configure the IP address and the user token for authentication with the Philips Hue Bridge.").
		AddFlag("ip,i", "The IP Address of the Philips Hue bridge on your local network", commando.String, nil).
		AddFlag("user-token,u", "The User Token generated by your Philips Hue bridge for authentication", commando.String, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			addr := flags["ip"].Value
			userToken := flags["user-token"].Value
			var hueIP net.IP

			if str, ok := addr.(string); ok {
				/* act on str */
				hueIP = net.ParseIP(str)
			} else {
				/* not string */
				log.Fatalln("Invalid IP Address")
			}

			if userTokenStr, userTokenOk := userToken.(string); userTokenOk {
				config(hueIP, userTokenStr)
			} else {
				log.Fatalln("Invalid user token, user token provided was not a string")
			}
		})

	// Configure the lights command
	// $ go-hue get-lights
	commando.
		Register("get-lights").
		SetShortDescription("displays detailed information about all lights connected to your Philips Hue Bridge.").
		SetDescription("this command displays more information about the state of all of the lights connected to your Philips Hue Bridge.").
		AddFlag("light,l", "A specific light you want the state of", commando.Int, 0).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			configData := readFromConfig()
			hue.GetLightState(configData.HueIP, configData.HueUser)
		})

	// Configure the groups command
	// $ go-hue get-groups
	commando.
		Register("get-groups").
		SetShortDescription("displays detailed information about all groups connected to your Philips Hue Bridge.").
		SetDescription("this command displays more information about the state of all of the lights connected to your Philips Hue Bridge.").
		AddFlag("group,g", "A specific group you want the state of", commando.String, "default").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			configData := readFromConfig()
			hue.GetGroups(configData.HueIP, configData.HueUser)
		})

	// Configure the scenes command
	// $ go-hue get-scenes
	commando.
		Register("get-scenes").
		SetShortDescription("displays detailed information about all scenes on your Philips Hue Bridge.").
		SetDescription("this command displays more information about all of the scenes on your Philips Hue Bridge.").
		AddFlag("names-only,n", "Return just the name of scenes", commando.Bool, "false").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			configData := readFromConfig()

			fNames := flags["names-only"].Value
			var namesOnly bool
			if names, ok := fNames.(bool); ok {
				namesOnly = names
			}

			if namesOnly {
				hue.GetScenesNames(configData.HueIP, configData.HueUser)
			} else {
				hue.GetScenes(configData.HueIP, configData.HueUser)
			}
		})

	// Configure the set-groups command
	// $ go-hue set-groups
	commando.
		Register("set-groups").
		SetShortDescription("Set values of a hue groups state").
		SetDescription("this command sets the state of the specified groups lights connected to your Philips Hue Bridge.").
		AddFlag("group,g", "A specific group you want so set the state of", commando.String, nil).
		AddFlag("value,v", "Value to set the lights 'on' or 'off'", commando.String, "on").
		AddFlag("brightness,b", "Value to set the brightness to, between 1 and 254", commando.Int, 0).
		AddFlag("saturation,s", "Value to set the saturation to, betwen 1 and 254. 254 is the most saturated (colored) and 0 is the least saturated (white)", commando.Int, 0).
		AddFlag("hue,h", "Value to set the hue to. The hue value to set light to.The hue value is a wrapping value between 0 and 65535. Both 0 and 65535 are red, 25500 is green and 46920 is blue.", commando.Int, 0).
		AddFlag("effect,e", "The dynamic effect of the light. Currently ???none??? and ???colorloop??? are supported. Other values will generate an error of type 7.Setting the effect to colorloop will cycle through all hues using the current brightness and saturation settings", commando.String, "potato").
		AddFlag("scene,c", "The scene that you would like to set the light to", commando.String, "potato").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fValue := flags["value"].Value
			fGroup := flags["group"].Value
			fBright := flags["brightness"].Value
			fSat := flags["saturation"].Value
			fHue := flags["hue"].Value
			fEffect := flags["effect"].Value
			fScene := flags["scene"].Value

			var setHue int
			value := new(bool)
			var group string
			var bright int
			var sat int
			var effect string
			var scene string

			if valueStr, ok := fValue.(string); ok {
				/* act on str */
				if valueStr == "on" {
					*value = true
				} else if valueStr == "off" {
					*value = false
				} else {
					log.Fatalln("the value of 'value' must be either 'on' or 'off'")
				}
			} else {
				/* not string */
				log.Fatalln("Invalid value provided.")
			}

			if groupStr, ok := fGroup.(string); ok {
				group = groupStr
			} else {
				log.Fatalln("Invalid value provided to 'group'")
			}

			if brightInt, ok := fBright.(int); ok {
				bright = brightInt
			}

			if satInt, ok := fSat.(int); ok {
				sat = satInt
			}

			if hueInt, ok := fHue.(int); ok {
				setHue = hueInt
			}

			if sceneStr, ok := fScene.(string); ok {
				scene = sceneStr
			}

			if effectString, ok := fEffect.(string); ok {
				effect = effectString
			}

			configData := readFromConfig()

			params := hue.LightsAuthAndBody{
				Auth: hue.EndpointParams{
					HueIP:   configData.HueIP,
					HueUser: configData.HueUser,
					Item:    group,
				},
				Body: hue.LightsBodyOptions{
					On:         value,
					Brightness: uint8(bright),
					Saturation: uint8(sat),
					Hue:        uint16(setHue),
					Effect:     effect,
					Scene:      scene,
				},
			}

			hue.SetGroup(params)
		})

	// Configure the set-lights command
	// $ go-hue set-lights
	commando.
		Register("set-lights").
		SetShortDescription("Set values of a hue lights state").
		SetDescription("this command sets the state of the specified light connected to your Philips Hue Bridge.").
		AddFlag("light,l", "A specific light you want so set the state of", commando.String, nil).
		AddFlag("value,v", "Value to set the light 'on' or 'off'", commando.String, "on").
		AddFlag("brightness,b", "Value to set the brightness to, between 1 and 254", commando.Int, 0).
		AddFlag("saturation,s", "Value to set the saturation to, betwen 1 and 254. 254 is the most saturated (colored) and 0 is the least saturated (white)", commando.Int, 0).
		AddFlag("hue,h", "Value to set the hue to. The hue value to set light to.The hue value is a wrapping value between 0 and 65535. Both 0 and 65535 are red, 25500 is green and 46920 is blue.", commando.Int, -1).
		AddFlag("effect,e", "The dynamic effect of the light. Currently ???none??? and ???colorloop??? are supported. Other values will generate an error of type 7.Setting the effect to colorloop will cycle through all hues using the current brightness and saturation settings", commando.String, "none").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fValue := flags["value"].Value
			fLight := flags["light"].Value
			fBright := flags["brightness"].Value
			fSat := flags["saturation"].Value
			fHue := flags["hue"].Value
			fEffect := flags["effect"].Value

			value := new(bool)
			var light string
			var bright int
			var sat int
			var setHue int
			var effect string

			if valueStr, ok := fValue.(string); ok {
				/* act on str */
				if valueStr == "on" {
					*value = true
				} else if valueStr == "off" {
					*value = false
				} else {
					log.Fatalln("the value of 'value' must be either 'on' or 'off'")
				}
			} else {
				/* not string */
				log.Fatalln("Invalid value provided.")
			}

			if lightStr, ok := fLight.(string); ok {
				light = lightStr
			} else {
				log.Fatalln("Invalid value provided to 'group'")
			}

			if brightInt, ok := fBright.(int); ok {
				bright = brightInt
			}

			if satInt, ok := fSat.(int); ok {
				sat = satInt
			}

			if hueInt, ok := fHue.(int); ok {
				setHue = hueInt
			}

			if effectString, ok := fEffect.(string); ok {
				effect = effectString
			}

			configData := readFromConfig()

			params := hue.LightsAuthAndBody{
				Auth: hue.EndpointParams{
					HueIP:   configData.HueIP,
					HueUser: configData.HueUser,
					Item:    light,
				},
				Body: hue.LightsBodyOptions{
					On:         value,
					Brightness: uint8(bright),
					Saturation: uint8(sat),
					Hue:        uint16(setHue),
					Effect:     effect,
				},
			}

			hue.ToggleLight(params)
		})

	// parse command-line arguments from the STDIN
	commando.Parse(nil)
}
