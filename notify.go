package haconnect

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// NotifyConfig provides configuration for Notify.
type NotifyConfig struct {

	// NotifyCallback is invoked when a new notification is received.
	NotifyCallback func(payload string) `json:"-"`
}

type haconnectNotify struct {
	*haconnectEntityConfig
	*EntityConfig
	*NotifyConfig
	Device       *haconnectDevice `json:"device"`
	CommandTopic string           `json:"command_topic"`
}

// Notify represents an entity that receives and processes notifications.
type Notify struct {
	Entity
}

// Notify creates a new notify entity.
func (c *Conn) Notify(
	entityCfg *EntityConfig,
	cfg *NotifyConfig,
) (*Notify, error) {
	cmdTopic := c.cmdTopic(entityCfg.ID, "notify")
	if err := c.publishSafeJSON(
		c.cfgTopic(entityCfg.ID, "notify"),
		&haconnectNotify{
			haconnectEntityConfig: c.buildEntityConfig(entityCfg.ID),
			EntityConfig:          entityCfg,
			NotifyConfig:          cfg,
			Device:                c.device,
			CommandTopic:          cmdTopic,
		},
	); err != nil {
		return nil, err
	}
	if t := c.client.Subscribe(
		cmdTopic,
		0,
		func(client mqtt.Client, msg mqtt.Message) {
			cfg.NotifyCallback(string(msg.Payload()))
		},
	); t.Wait() && t.Error() != nil {
		return nil, t.Error()
	}
	n := &Notify{}
	if err := c.initEntity(n, entityCfg); err != nil {
		return nil, err
	}
	return n, nil
}
