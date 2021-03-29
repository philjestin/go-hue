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
	log.Printf(sb)
}

// ToggleGroup is used to toggle a given groups on/off values
func ToggleGroup(params ToggleParams) {
	on := utils.On{
		On: params.OnValue,
	}

	update := utils.UpdateObject{
		URL:    fmt.Sprintf("https://%s/api/%s/groups/%s/action", params.HueIP, params.HueUser, params.Item),
		Client: utils.GetClient(),
		On:     on,
	}

	res := utils.UpdateItem(update)
	//Convert the body to type string
	sb := string(res)
	log.Printf(sb)
}
