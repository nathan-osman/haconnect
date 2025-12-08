package hamqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// LightConfig provides configuration for light entities.
type LightConfig struct {

	// State indicates the initial state of the light.
	State bool `json:"-"`

	// ChangeCallback is invoked when the light's value is changed. Returning
	// true will cause a corresponding change to the light's state.
	ChangeCallback func(bool) bool `json:"-"`
}

type hamqttLight struct {
	*hamqttEntityConfig
	*EntityConfig
	*LightConfig
	Device       *hamqttDevice `json:"device"`
	Platform     string        `json:"platform"`
	CommandTopic string        `json:"command_topic"`
	StateTopic   string        `json:"state_topic"`
}

// Light provides methods for controlling a light entity.
type Light struct {
	Entity
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
		&hamqttLight{
			hamqttEntityConfig: c.buildEntityConfig(entityCfg.ID),
			EntityConfig:       entityCfg,
			LightConfig:        cfg,
			Device:             c.device,
			Platform:           "light",
			CommandTopic:       cmdTopic,
			StateTopic:         stateTopic,
		},
	); err != nil {
		return nil, err
	}
	if t := c.client.Subscribe(
		cmdTopic,
		0,
		func(client mqtt.Client, msg mqtt.Message) {
			switch string(msg.Payload()) {
			case "ON":
				if cfg.ChangeCallback(true) {
					c.publishSafeState(stateTopic, true)
				}
			case "OFF":
				if cfg.ChangeCallback(false) {
					c.publishSafeState(stateTopic, false)
				}
			}
		},
	); t.Wait() && t.Error() != nil {
		return nil, t.Error()
	}
	l := &Light{}
	if err := c.initEntity(l, entityCfg); err != nil {
		return nil, err
	}
	return l, nil
}
