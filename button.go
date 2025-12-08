package hamqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	ButtonIdentify = "identify"
	ButtonRestart  = "restart"
	ButtonUpdate   = "update"
)

// ButtonConfig provides configuration for button entities.
type ButtonConfig struct {

	// DeviceClass categorizes the button.
	DeviceClass string `json:"device_class,omitempty"`

	// PressCallback is invoked when the button is pressed.
	PressCallback func() `json:"-"`
}

type hamqttButton struct {
	*hamqttEntityConfig
	*EntityConfig
	*ButtonConfig
	Device       *hamqttDevice `json:"device"`
	Platform     string        `json:"platform"`
	CommandTopic string        `json:"command_topic"`
}

// Button provides methods for interacting with a button.
type Button struct {
	Entity
}

// Button creates a new button entity with the provided configuration.
func (c *Conn) Button(
	entityCfg *EntityConfig,
	cfg *ButtonConfig,
) (*Button, error) {
	cmdTopic := c.cmdTopic(entityCfg.ID, "button")
	if err := c.publishSafeJSON(
		c.cfgTopic(entityCfg.ID, "button"),
		&hamqttButton{
			hamqttEntityConfig: c.buildEntityConfig(entityCfg.ID),
			EntityConfig:       entityCfg,
			ButtonConfig:       cfg,
			Device:             c.device,
			Platform:           "button",
			CommandTopic:       cmdTopic,
		},
	); err != nil {
		return nil, err
	}
	if t := c.client.Subscribe(
		cmdTopic,
		0,
		func(client mqtt.Client, msg mqtt.Message) {
			if string(msg.Payload()) == "PRESS" {
				cfg.PressCallback()
			}
		},
	); t.Wait() && t.Error() != nil {
		return nil, t.Error()
	}
	b := &Button{}
	if err := c.initEntity(b, entityCfg); err != nil {
		return nil, err
	}
	return b, nil
}
