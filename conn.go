package haconnect

import (
	"fmt"
	"os"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	payloadOnline  = "online"
	payloadOffline = "offline"

	emptyAvailability = "{}"
)

type haconnectDevice struct {
	IDs          string `json:"identifiers"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Model        string `json:"model,omitempty"`
}

// Conn maintains a connection to an MQTT broker.
type Conn struct {
	mutex             sync.Mutex
	client            mqtt.Client
	discoveryPrefix   string
	id                string
	availabilityTopic string
	availability      map[string]string
	device            *haconnectDevice
}

// New creates a new Conn instance with the provided configuration.
func New(cfg *Config) (*Conn, error) {

	// If DiscoveryPrefix is empty, use the default
	discoveryPrefix := cfg.DiscoveryPrefix
	if discoveryPrefix == "" {
		discoveryPrefix = "homeassistant"
	}

	// If ID and / or Name was not provided, use the hostname
	var (
		id   = cfg.ID
		name = cfg.Name
	)
	if id == "" {
		v, err := os.Hostname()
		if err != nil {
			return nil, err
		}
		id = v
	}
	if name == "" {
		name = id
	}

	// Generate the availability topic
	availabilityTopic := fmt.Sprintf("%s/availability", id)

	// Connect to the MQTT broker
	client := mqtt.NewClient(
		mqtt.NewClientOptions().
			AddBroker(fmt.Sprintf("tcp://%s", cfg.Addr)).
			SetUsername(cfg.Username).
			SetPassword(cfg.Password).
			SetClientID(id).
			SetAutoReconnect(true).
			SetCleanSession(false).
			SetConnectRetry(true).
			SetWill(
				availabilityTopic,
				emptyAvailability,
				0,
				true,
			),
	)

	// Ensure the connection was successful
	if t := client.Connect(); t.Wait() && t.Error() != nil {
		return nil, t.Error()
	}

	// Create the Conn
	c := &Conn{
		client:            client,
		discoveryPrefix:   discoveryPrefix,
		id:                id,
		availabilityTopic: availabilityTopic,
		availability:      make(map[string]string),
		device: &haconnectDevice{
			IDs:          id,
			Name:         name,
			Manufacturer: cfg.Manufacturer,
			Model:        cfg.Model,
		},
	}

	// Initialize availablity
	if err := c.publishSafeJSON(
		c.availabilityTopic,
		emptyAvailability,
	); err != nil {
		c.client.Disconnect(0)
		return nil, err
	}

	// Return the Conn
	return c, nil
}

// Close shuts down the connection.
func (c *Conn) Close() {

	// Ignore any errors returned here since the LWT will do the same thing
	c.publishFast(
		c.availabilityTopic,
		emptyAvailability,
	)
	c.client.Disconnect(1000)
}
