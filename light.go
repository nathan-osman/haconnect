package haconnect

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// LightConfig provides configuration for light entities.
type LightConfig struct {

	// State indicates the initial state of the light.
	State bool `json:"-"`

	// ChangeCallback is invoked when the light is to be turned on or off.
	// Returning true will cause a corresponding change to the light's state.
	ChangeCallback func(bool) bool `json:"-"`
}

type haconnectLight struct {
	*haconnectEntityConfig
	*EntityConfig
	*LightConfig
	Device       *haconnectDevice `json:"device"`
	Platform     string           `json:"platform"`
	CommandTopic string           `json:"command_topic"`
	StateTopic   string           `json:"state_topic"`
}

// Light provides methods for controlling a light entity.
type Light struct {
	Entity
	stateTopic string
}

// SetValue indicates whether the light is on or off.
func (l *Light) SetValue(value bool) error {
	return l.conn.publishSafeState(l.stateTopic, value)
}

// Light creates a new light entity with the provided configuration.
func (c *Conn) Light(
	entityCfg *EntityConfig,
	cfg *LightConfig,
) (*Light, error) {
	var (
		cmdTopic   = c.cmdTopic(entityCfg.ID, "light")
		stateTopic = c.stateTopic(entityCfg.ID)
	)
	if err := c.publishSafeState(stateTopic, cfg.State); err != nil {
		return nil, err
	}
	if err := c.publishSafeJSON(
		c.cfgTopic(entityCfg.ID, "light"),
		&haconnectLight{
			haconnectEntityConfig: c.buildEntityConfig(entityCfg.ID),
			EntityConfig:          entityCfg,
			LightConfig:           cfg,
			Device:                c.device,
			Platform:              "light",
			CommandTopic:          cmdTopic,
			StateTopic:            stateTopic,
		},
	); err != nil {
		return nil, err
	}
	l := &Light{
		stateTopic: stateTopic,
	}
	if t := c.client.Subscribe(
		cmdTopic,
		0,
		func(client mqtt.Client, msg mqtt.Message) {
			switch string(msg.Payload()) {
			case "ON":
				if cfg.ChangeCallback(true) {
					l.SetValue(true)
				}
			case "OFF":
				if cfg.ChangeCallback(false) {
					l.SetValue(false)
				}
			}
		},
	); t.Wait() && t.Error() != nil {
		return nil, t.Error()
	}
	if err := c.initEntity(l, entityCfg); err != nil {
		return nil, err
	}
	return l, nil
}
