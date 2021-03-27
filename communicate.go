package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// On to turn on a item the request body
type On struct {
	On bool `json:"on"`
}

func getClient() *http.Client {
	// No valid TLS Cert on Philips Hue Bridge
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return client
}

func readResponseBody(resp *http.Response, err error) []byte {
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

// getLightState retrieves light state from the Hue Bridge
func getLightState(hueIP net.IP, hueUser string) {
	// State url
	stateURL := fmt.Sprintf("https://%s/api/%s/lights", hueIP, hueUser)
	client := getClient()
	resp, err := client.Get(stateURL)

	res := readResponseBody(resp, err)

	//Convert the body to type string
	sb := string(res)
	log.Printf(sb)
}

// getGroups retrieves data about groups from the Hue Bridge
func getGroups(hueIP net.IP, hueUser string) {
	// Groups URL
	groupsURL := fmt.Sprintf("https://%s/api/%s/groups", hueIP, hueUser)
	client := getClient()
	resp, err := client.Get(groupsURL)
	res := readResponseBody(resp, err)

	//Convert the body to type string
	sb := string(res)
	log.Printf(sb)
}

func toggleGroup(hueIP net.IP, hueUser string, group string, onValue bool) {
	// Groups URL
	groupsURL := fmt.Sprintf("https://%s/api/%s/groups/%s/action", hueIP, hueUser, group)
	client := getClient()

	on := On{
		On: onValue,
	}

	// marshal User to json
	json, err := json.Marshal(on)
	if err != nil {
		log.Fatalln(err)
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, groupsURL, bytes.NewBuffer(json))
	if err != nil {
		log.Fatalln(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	res := readResponseBody(resp, err)
	//Convert the body to type string
	sb := string(res)
	log.Printf(sb)
}
