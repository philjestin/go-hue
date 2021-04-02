package hue

import (
	"net"
	"net/http"
)

// Endpoint Params
type EndpointParams struct {
	HueIP   net.IP
	HueUser string
	Item    string
}

// ToggleParams struct
type ToggleParams struct {
	Auth    EndpointParams
	OnValue bool
}

// LightsBodyOptions
type LightsBodyOptions struct {
	On         *bool  `json:"on,omitempty" bson:",omitempty"`
	Brightness uint8  `json:"bri,omitempty" bson:",omitempty"`
	Hue        uint16 `json:"hue,omitempty" bson:",omitempty"`
	Saturation uint8  `json:"sat,omitempty" bson:",omitempty"`
	Effect     string `json:"effect,omitempty" bson:",omitempty"`
	Scene      string `json:"scene,omitempty" bson:",omitempty"`
}

type LightsAuthAndBody struct {
	Auth EndpointParams
	Body LightsBodyOptions
}

type UpdateObject struct {
	URL    string
	Client *http.Client
	Body   LightsBodyOptions
}
