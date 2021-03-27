package main

import (
	"fmt"
	"log"
	"net"
)

func findBridge() {
	list, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
	}

	for i, iface := range list {
		fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatalln(err)
		}
		for j, addr := range addrs {
			fmt.Printf(" %d %v\n", j, addr)
		}
	}
}
