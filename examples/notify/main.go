package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/hamqtt"
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
	c, err := hamqtt.New(&hamqtt.Config{
		Addr:     *addr,
		Username: *username,
		Password: *password,
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Create a notify entity which writes to STDOUT when a notification is
	// received
	if _, err := c.Notify(
		&hamqtt.EntityConfig{
			ID:   "mynotify",
			Name: "My Notify",
		},
		&hamqtt.NotifyConfig{
			NotifyCallback: func(payload string) {
				fmt.Println(payload)
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
