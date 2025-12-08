package hamqtt

import (
	"encoding/json"
	"fmt"
)

const (
	payloadOn  = "ON"
	payloadOff = "OFF"
)

func (c *Conn) cfgTopic(id, entityType string) string {
	return fmt.Sprintf(
		"%s/%s/%s/config",
		c.discoveryPrefix,
		entityType,
		id,
	)
}

func (c *Conn) cmdTopic(id, entityType string) string {
	return fmt.Sprintf(
		"%s/%s/%s",
		c.id,
		id,
		entityType,
	)
}

func (c *Conn) stateTopic(id string) string {
	return fmt.Sprintf(
		"%s/%s/state",
		c.id,
		id,
	)
}

// publishFast sends a message with QoS 0, which is immediate.
func (c *Conn) publishFast(topic, payload string) error {
	t := c.client.Publish(topic, 0, true, payload)
	return t.Error()
}

// publishSafe ensures at least one broker has acknowledged the message.
func (c *Conn) publishSafe(topic, payload string) error {
	if t := c.client.Publish(topic, 1, true, payload); t.Wait() && t.Error() != nil {
		return t.Error()
	}
	return nil
}

func (c *Conn) publishSafeJSON(topic string, payload any) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return c.publishSafe(topic, string(b))
}

func (c *Conn) publishSafeBool(
	topic string,
	value bool,
	payloadTrue string,
	payloadFalse string,
) error {
	payload := payloadFalse
	if value {
		payload = payloadTrue
	}
	return c.publishSafe(topic, payload)
}

func (c *Conn) publishSafeState(topic string, value bool) error {
	return c.publishSafeBool(topic, value, payloadOn, payloadOff)
}
