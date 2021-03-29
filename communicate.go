package main

import (
	"fmt"
	"log"
	"net"

	"github.com/philjestin/go-hue/utils"
)

// getLightState retrieves light state from the Hue Bridge
func getLightState(hueIP net.IP, hueUser string) {
	// State url
	stateURL := fmt.Sprintf("https://%s/api/%s/lights", hueIP, hueUser)
	client := utils.GetClient()
	resp, err := client.Get(stateURL)

	res := utils.ReadResponseBody(resp, err)

	//Convert the body to type string
	sb := string(res)
	log.Printf(sb)
}

// getGroups retrieves data about groups from the Hue Bridge
func getGroups(hueIP net.IP, hueUser string) {
	// Groups URL
	groupsURL := fmt.Sprintf("https://%s/api/%s/groups", hueIP, hueUser)
	client := utils.GetClient()
	resp, err := client.Get(groupsURL)
	res := utils.ReadResponseBody(resp, err)

	//Convert the body to type string
	sb := string(res)
	log.Printf(sb)
}

func toggleLight(hueIP net.IP, hueUser string, group string, onValue bool) {
	on := utils.On{
		On: onValue,
	}

	update := utils.UpdateObject{
		URL:    fmt.Sprintf("https://%s/api/%s/lights/%s/action", hueIP, hueUser, group),
		Client: utils.GetClient(),
		On:     on,
	}

	res := utils.UpdateItem(update)

	//Convert the body to type string
	sb := string(res)
	log.Printf(sb)
}

func toggleGroup(hueIP net.IP, hueUser string, group string, onValue bool) {
	on := utils.On{
		On: onValue,
	}

	update := utils.UpdateObject{
		URL:    fmt.Sprintf("https://%s/api/%s/groups/%s/action", hueIP, hueUser, group),
		Client: utils.GetClient(),
		On:     on,
	}

	res := utils.UpdateItem(update)
	//Convert the body to type string
	sb := string(res)
	log.Printf(sb)
}
