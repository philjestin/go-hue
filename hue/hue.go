package hue

import "net"

// ToggleParams struct
type ToggleParams struct {
	HueIP   net.IP
	HueUser string
	Item    string
	OnValue bool
}
