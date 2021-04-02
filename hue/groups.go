package hue

import (
	"fmt"
	"log"
	"net"

	"github.com/philjestin/go-hue/utils"
)

// GetGroups retrieves data about groups from the Hue Bridge
func GetGroups(hueIP net.IP, hueUser string) {
	// Groups URL
	groupsURL := fmt.Sprintf("https://%s/api/%s/groups", hueIP, hueUser)
	client := utils.GetClient()
	resp, err := client.Get(groupsURL)
	res := utils.ReadResponseBody(resp, err)

	//Convert the body to type string
	sb := string(res)
	log.Println(sb)
}

// SetGroup is used to set the state values for a group.
func SetGroup(params LightsAuthAndBody) {

	if params.Body.Effect != "none" && params.Body.Effect != "colorloop" {
		params.Body.Effect = ""
	}

	if params.Body.Scene == "potato" {
		params.Body.Scene = ""
	}

	fmt.Println(params)

	update := utils.UpdateObject{
		URL:    fmt.Sprintf("https://%s/api/%s/groups/%s/action", params.Auth.HueIP, params.Auth.HueUser, params.Auth.Item),
		Client: utils.GetClient(),
		Body: utils.LightsBodyOptions{
			On:         params.Body.On,
			Brightness: params.Body.Brightness,
			Saturation: params.Body.Saturation,
			Hue:        params.Body.Hue,
			Effect:     params.Body.Effect,
		}}

	res := utils.UpdateItem(update)
	//Convert the body to type string
	sb := string(res)
	log.Println(sb)
}
