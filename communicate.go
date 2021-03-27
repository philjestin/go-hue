package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// getLightState retrieves light state from the Hue Bridge
func getLightState(hueIP net.IP, hueUser string) {
	// No valid TLS Cert on Philips Hue Bridge
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	hueAddr := fmt.Sprintf("https://%s/api/%s/lights", hueIP, hueUser)
	resp, err := client.Get(hueAddr)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
