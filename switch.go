package haconnect

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	SwitchOutlet = "outlet"
	SwitchSwitch = "switch"
)

// SwitchConfig provides configuration for Switch.
type SwitchConfig struct {

	// State provides the initial state of the switch.
	State bool `json:"-"`

	// DeviceClass categorizes the type of switch.
	DeviceClass string `json:"device_class,omitempty"`

	// ChangeCallback is invoked when the switch is to be turned on or off.
	// Returning true will cause a corresponding change to the switch's state.
	// Alternatively, SetValue can be used to indicate state.
	ChangeCallback func(bool) bool `json:"-"`
}

type haconnectSwitch struct {
	*haconnectEntityConfig
	*EntityConfig
	*SwitchConfig
	Device       *haconnectDevice `json:"device"`
	Platform     string           `json:"platform"`
	CommandTopic string           `json:"command_topic"`
	StateTopic   string           `json:"state_topic"`
}

// Switch represents an entity that can be switched on and off.
type Switch struct {
	Entity
	stateTopic string
}

// SetValue indicates whether the switch is on or off.
func (s *Switch) SetValue(value bool) error {
	return s.conn.publishSafeState(s.stateTopic, value)
}

// Switch creates a new switch entity.
func (c *Conn) Switch(
	entityCfg *EntityConfig,
	cfg *SwitchConfig,
) (*Switch, error) {
	var (
		cmdTopic   = c.cmdTopic(entityCfg.ID, "switch")
		stateTopic = c.stateTopic(entityCfg.ID)
	)
	if err := c.publishSafeState(stateTopic, cfg.State); err != nil {
		return nil, err
	}
	if err := c.publishSafeJSON(
		c.cfgTopic(entityCfg.ID, "switch"),
		&haconnectSwitch{
			haconnectEntityConfig: c.buildEntityConfig(entityCfg.ID),
			EntityConfig:          entityCfg,
			SwitchConfig:          cfg,
			Device:                c.device,
			Platform:              "switch",
			CommandTopic:          cmdTopic,
			StateTopic:            stateTopic,
		},
	); err != nil {
		return nil, err
	}
	s := &Switch{
		stateTopic: stateTopic,
	}
	if t := c.client.Subscribe(
		cmdTopic,
		0,
		func(client mqtt.Client, msg mqtt.Message) {
			switch string(msg.Payload()) {
			case "ON":
				if cfg.ChangeCallback(true) {
					s.SetValue(true)
				}
			case "OFF":
				if cfg.ChangeCallback(false) {
					s.SetValue(false)
				}
			}
		},
	); t.Wait() && t.Error() != nil {
		return nil, t.Error()
	}
	if err := c.initEntity(s, entityCfg); err != nil {
		return nil, err
	}
	return s, nil
}
