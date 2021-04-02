package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type LightsBodyOptions struct {
	On         *bool  `json:"on,omitempty" bson:",omitempty"`
	Brightness uint8  `json:"bri,omitempty" bson:",omitempty"`
	Hue        uint16 `json:"hue,omitempty" bson:",omitempty"`
	Saturation uint8  `json:"sat,omitempty" bson:",omitempty"`
	Effect     string `json:"effect,omitempty" bson:",omitempty"`
	Scene      string `json:"scene,omitempty" bson:",omitempty"`
}

// UpdateObject required for updating an item
type UpdateObject struct {
	URL    string
	Client *http.Client
	Body   LightsBodyOptions
}

// GetClient returns a http Client with insecure skip verify set to true
func GetClient() *http.Client {
	// No valid TLS Cert on Philips Hue Bridge
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return client
}

// ReadResponseBody is a helper func to read the response from the hue api
func ReadResponseBody(resp *http.Response, err error) []byte {
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

// UpdateItem is a helper func used for turning lights on and off
func UpdateItem(updateItem UpdateObject) []byte {
	// marshal User to json
	json, err := json.Marshal(updateItem.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, updateItem.URL, bytes.NewBuffer(json))
	if err != nil {
		log.Fatalln(err)
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := updateItem.Client.Do(req)

	res := ReadResponseBody(resp, err)

	return res
}
