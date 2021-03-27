package main

import (
	"net"
	"encoding/json"
	"io/ioutil"
	"log"
)

// Configuration defines whats required to communicate
// with your Philips Hue Bridge
type Configuration struct {
	HueIP   net.IP
	HueUser string
}

func config(hueIp net.IP, hueUser string) {
	data := Configuration{
		HueIP: hueIp,
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