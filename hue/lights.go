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
	log.Printf(sb)
}

// ToggleLight is used to toggle a lights on/off value
func ToggleLight(params ToggleParams) {
	on := utils.On{
		On: params.OnValue,
	}

	update := utils.UpdateObject{
		URL:    fmt.Sprintf("https://%s/api/%s/lights/%s/state", params.HueIP, params.HueUser, params.Item),
		Client: utils.GetClient(),
		On:     on,
	}

	res := utils.UpdateItem(update)

	//Convert the body to type string
	sb := string(res)
	log.Printf(sb)
}
