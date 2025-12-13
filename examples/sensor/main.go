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

	// Create a thermometer
	s, err := c.Sensor(
		&haconnect.EntityConfig{
			ID:   "mythermometer",
			Name: "My Thermometer",
		},
		&haconnect.SensorConfig{
			DeviceClass:               haconnect.SensorTemperature,
			UnitOfMeasurement:         haconnect.SensorDegreesCelsius,
			SuggestedDisplayPrecision: 1,
		},
	)
	if err != nil {
		panic(err)
	}

	// Read interactive commands from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("New value for temperature in Celsius? (EOF to quit)")
		if !scanner.Scan() {
			return
		}
		s.SetValue(scanner.Text())
	}
}
