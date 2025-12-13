package haconnect

import (
	"encoding/json"
	"fmt"
)

// EntityConfig provides basic configuration for all entities.
type EntityConfig struct {

	// ID provides the unique ID of the entity.
	ID string `json:"unique_id"`

	// Name provides the name of the entity.
	Name string `json:"name"`

	// Unavailable makes the entity unavailable until explicitly set.
	Unavailable bool `json:"-"`
}

type haconnectEntityConfig struct {
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
}

func (c *Conn) buildEntityConfig(id string) *haconnectEntityConfig {
	return &haconnectEntityConfig{
		AvailabilityTemplate: fmt.Sprintf(
			"{{ value_json['%s'] | default('%s') }}",
			id,
			payloadOffline,
		),
		AvailabilityTopic: c.availabilityTopic,
	}
}

type iEntity interface {
	init(*Conn, *EntityConfig) error
}

func (c *Conn) initEntity(e iEntity, cfg *EntityConfig) error {
	if err := e.init(c, cfg); err != nil {
		return err
	}
	return nil
}

// Entity provides base functionality for all entities.
type Entity struct {
	conn *Conn
	id   string
}

func (e *Entity) init(c *Conn, cfg *EntityConfig) error {
	e.conn = c
	e.id = cfg.ID
	return e.SetAvailability(!cfg.Unavailable)
}

// SetAvailability indicates whether the entity is available or not.
func (e *Entity) SetAvailability(availability bool) error {
	payload := payloadOffline
	if availability {
		payload = payloadOnline
	}
	v, err := func() (string, error) {
		e.conn.mutex.Lock()
		defer e.conn.mutex.Unlock()
		e.conn.availability[e.id] = payload
		b, err := json.Marshal(e.conn.availability)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}()
	if err != nil {
		return err
	}
	return e.conn.publishSafe(
		e.conn.availabilityTopic,
		v,
	)
}
