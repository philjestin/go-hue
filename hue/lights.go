package hue

import (
	"fmt"
	"log"
	"net"

	"github.com/philjestin/go-hue/utils"
)

// GetLightState retrieves light state from the Hue Bridge
func GetLightState(hueIP net.IP, hueUser string) {
	// State url
	stateURL := fmt.Sprintf("https://%s/api/%s/lights", hueIP, hueUser)
	client := utils.GetClient()
	resp, err := client.Get(stateURL)

	res := utils.ReadResponseBody(resp, err)

	//Convert the body to type string
	sb := string(res)
	log.Println(sb)
}

// ToggleLight is used to toggle a lights on/off value
func ToggleLight(params LightsAuthAndBody) {
	update := utils.UpdateObject{
		URL:    fmt.Sprintf("https://%s/api/%s/lights/%s/state", params.Auth.HueIP, params.Auth.HueUser, params.Auth.Item),
		Client: utils.GetClient(),
		Body: utils.LightsBodyOptions{
			On:         params.Body.On,
			Brightness: params.Body.Brightness,
			Saturation: params.Body.Saturation,
			Hue:        params.Body.Hue,
			Effect:     params.Body.Effect,
		},
	}

	res := utils.UpdateItem(update)

	//Convert the body to type string
	sb := string(res)
	log.Println(sb)
}
