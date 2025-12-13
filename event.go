package haconnect

const (
	EventButton   = "button"
	EventDoorbell = "doorbell"
	EventMotion   = "motion"
)

// EventConfig provides configuration for event entities.
type EventConfig struct {

	// DeviceClass categorizes the type of event.
	DeviceClass string `json:"device_class,omitempty"`

	// EventTypes is a list of valid events (i.e. "press", "hold", etc.).
	EventTypes []string `json:"event_types,omitempty"`
}

type haconnectEvent struct {
	*haconnectEntityConfig
	*EntityConfig
	*EventConfig
	Device     *haconnectDevice `json:"device"`
	Platform   string           `json:"platform"`
	StateTopic string           `json:"state_topic"`
}

// Event provides methods for sending events.
type Event struct {
	Entity
	stateTopic string
}

// Send sends the specified event.
func (e *Event) Send(eventType string) error {
	return e.conn.publishSafeJSON(
		e.stateTopic,
		map[string]any{
			"event_type": eventType,
		},
	)
}

// Event creates a new event entity with the provided configuration.
func (c *Conn) Event(
	entityCfg *EntityConfig,
	cfg *EventConfig,
) (*Event, error) {
	stateTopic := c.stateTopic(entityCfg.ID)
	if err := c.publishSafeJSON(
		c.cfgTopic(entityCfg.ID, "event"),
		&haconnectEvent{
			haconnectEntityConfig: c.buildEntityConfig(entityCfg.ID),
			EntityConfig:          entityCfg,
			EventConfig:           cfg,
			Device:                c.device,
			Platform:              "event",
			StateTopic:            stateTopic,
		},
	); err != nil {
		return nil, err
	}
	e := &Event{
		stateTopic: stateTopic,
	}
	if err := c.initEntity(e, entityCfg); err != nil {
		return nil, err
	}
	return e, nil
}
