package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	// Create a button entity
	if _, err := c.Button(
		&haconnect.EntityConfig{
			ID:   "mybutton",
			Name: "My Restart Button",
		},
		&haconnect.ButtonConfig{
			DeviceClass: haconnect.ButtonRestart,
			PressCallback: func() {
				fmt.Println("Restart button was pressed!")
			},
		},
	); err != nil {
		panic(err)
	}

	// Wait for SIGINT or SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
