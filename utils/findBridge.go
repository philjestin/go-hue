package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Bridge struct {
	Host string `json:"internalipaddress,omitempty"`
	User string
	ID   string `json:"id,omitempty"`
}

func FindBridge() []Bridge {
	req, err := http.NewRequest(http.MethodGet, "https://discovery.meethue.com", nil)

	if err != nil {
		log.Fatalln("Couldn't connect to Hue nupnp service.", err)
	}

	client := http.DefaultClient
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln("Couldn't connect to Hue nupnp service.", err)
	}

	res := ReadResponseBody(resp, err)

	var bridges []Bridge

	err = json.Unmarshal(res, &bridges)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(res)
	log.Println(sb)

	return bridges
}
