package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

// Configuration defines whats required to communicate
// with your Philips Hue Bridge
type Configuration struct {
	HueIP   net.IP
	HueUser string
}

func readFromConfig() Configuration {
	file, err := ioutil.ReadFile("./config/config.json")

	if err != nil {
		log.Fatalln("'go-hue config --ip <your-bridges-ip> --user-token <bridge-generated-token>' must be ran first.")
	}

	data := Configuration{}

	_ = json.Unmarshal([]byte(file), &data)

	return data
}

// config is used to set the configuration settings in your ./config/config.json file
func config(hueIP net.IP, hueUser string) {
	data := Configuration{
		HueIP:   hueIP,
		HueUser: hueUser,
	}

	file, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("./config/config.json", file, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func printConfig() {
	data := readFromConfig()

	fmt.Printf("Hue Bridge IP: %v\n", data.HueIP)
	fmt.Printf("Hue Bridge User Token: %v\n", data.HueUser)
}
