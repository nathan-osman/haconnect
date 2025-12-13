package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/nathan-osman/haconnect"
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

	// Create a binary sensor (for a door)
	s, err := c.BinarySensor(
		&haconnect.EntityConfig{
			ID:   "mydoorsensor",
			Name: "My Door",
		},
		&haconnect.BinarySensorConfig{
			DeviceClass: haconnect.BinarySensorDoor,
		},
	)
	if err != nil {
		panic(err)
	}

	// Read interactive commands from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Command? (open / close / avail / unavail / quit)")
		if !scanner.Scan() {
			return
		}
		switch scanner.Text() {
		case "open":
			s.SetValue(true)
		case "close":
			s.SetValue(false)
		case "avail":
			s.SetAvailability(true)
		case "unavail":
			s.SetAvailability(false)
		case "quit":
			return
		}
	}
}
