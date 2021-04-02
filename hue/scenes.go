package hue

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/philjestin/go-hue/utils"
)

type Scene struct {
	Name        string      `json:"id,omitempty"`
	SceneName   string      `json:"name,omitempty"`
	Type        string      `json:"type,omitempty"`
	Group       string      `json:"group,omitempty"`
	Lights      []string    `json:"lights,omitempty"`
	Owner       string      `json:"owner,omitempty"`
	Recycle     bool        `json:"recycle,omitempty"`
	Locked      bool        `json:"locked,omitempty"`
	AppData     interface{} `json:"appdata,omitempty"`
	Picture     string      `json:"picture,omitempty"`
	LastUpdated string      `json:"lastupdated,omitempty"`
	Version     int         `json:"version,omitempty"`
}

func getSceneData(hueIP net.IP, hueUser string) []byte {
	stateURL := fmt.Sprintf("https://%s/api/%s/scenes", hueIP, hueUser)
	client := utils.GetClient()
	resp, err := client.Get(stateURL)

	res := utils.ReadResponseBody(resp, err)

	return res
}

// GetScenes retrieves a list of all scenes currently stored in the bridge
func GetScenes(hueIP net.IP, hueUser string) {

	res := getSceneData(hueIP, hueUser)

	//Convert the body to type string
	sb := string(res)
	log.Println(sb)
}

func GetScenesNames(hueIP net.IP, hueUser string) {
	var m map[string]Scene
	res := getSceneData(hueIP, hueUser)

	err := json.Unmarshal(res, &m)

	if err != nil {
		log.Fatalln("PROBLEM")
	}

	// Duplice scene names exist because there are scenes per light.
	// Turn the slice of scenes into a set.
	var set []string
	for _, g := range m {
		if !contains(set, g.SceneName) {
			set = append(set, g.SceneName)
		}
	}

	for _, g := range set {
		fmt.Println(g)
	}

}

func contains(slice []string, s string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}
