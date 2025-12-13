package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/nathan-osman/haconnect"
)

const (
	Press = "press"
	Hold  = "hold"
)

func main() {
	var (
		addr     = flag.String("addr", "", "MQTT broker address")
		username = flag.String("username", "", "Username for authentication")
		password = flag.String("password", "", "Password for authentication")
	)

	// Parse CLI arguments
	flag.Parse()

	// Connect to the MQTT broker
	c, err := haconnect.New(&haconnect.Config{
		Addr:     *addr,
		Username: *username,
		Password: *password,
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Create an event for a doorbell
	e, err := c.Event(
		&haconnect.EntityConfig{
			ID:   "mydoorbell",
			Name: "My Doorbell",
		},
		&haconnect.EventConfig{
			DeviceClass: haconnect.EventDoorbell,
			EventTypes: []string{
				Press,
				Hold,
			},
		},
	)
	if err != nil {
		panic(err)
	}

	// Read interactive commands from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Command? (press / hold / quit)")
		if !scanner.Scan() {
			return
		}
		switch scanner.Text() {
		case Press:
			e.Send(Press)
		case Hold:
			e.Send(Hold)
		case "quit":
			return
		}
	}
}
