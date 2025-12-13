package haconnect

const (
	BinarySensorBattery         = "battery"
	BinarySensorBatteryCharging = "battery_charging"
	BinarySensorCarbonMonoxide  = "carbon_monoxide"
	BinarySensorCold            = "cold"
	BinarySensorConnectivity    = "connectivity"
	BinarySensorDoor            = "door"
	BinarySensorGarageDpor      = "garage_door"
	BinarySensorGas             = "gas"
	BinarySensorHeat            = "heat"
	BinarySensorLight           = "light"
	BinarySensorLock            = "lock"
	BinarySensorMoisture        = "moisture"
	BinarySensorMotion          = "motion"
	BinarySensorMoving          = "moving"
	BinarySensorOccupancy       = "occupancy"
	BinarySensorOpening         = "opening"
	BinarySensorPlug            = "plug"
	BinarySensorPower           = "power"
	BinarySensorPresence        = "presence"
	BinarySensorProblem         = "problem"
	BinarySensorRunning         = "running"
	BinarySensorSafety          = "safety"
	BinarySensorSmoke           = "smoke"
	BinarySensorSound           = "sound"
	BinarySensorTamper          = "tamper"
	BinarySensorUpdate          = "update"
	BinarySensorVibration       = "vibration"
	BinarySensorWindow          = "window"
)

// BinarySensorConfig provides configuration for BinarySensor.
type BinarySensorConfig struct {

	// State provides the initial state of the binary sensor.
	State bool `json:"-"`

	// DeviceClass categorizes the type of data reported by the sensor.
	DeviceClass string `json:"device_class,omitempty"`
}

type haconnectBinarySensor struct {
	*haconnectEntityConfig
	*EntityConfig
	*BinarySensorConfig
	Device     *haconnectDevice `json:"device"`
	Platform   string           `json:"platform"`
	StateTopic string           `json:"state_topic"`
}

// BinarySensor represents a sensor with binary state, such as a door being
// open or a switch being on.
type BinarySensor struct {
	Entity
	stateTopic string
}

// SetValue updates the binary sensor's value.
func (b *BinarySensor) SetValue(state bool) error {
	return b.conn.publishSafeState(b.stateTopic, state)
}

// BinarySensor creates a new binary sensor entity.
func (c *Conn) BinarySensor(
	entityCfg *EntityConfig,
	cfg *BinarySensorConfig,
) (*BinarySensor, error) {
	stateTopic := c.stateTopic(entityCfg.ID)
	if err := c.publishSafeState(stateTopic, cfg.State); err != nil {
		return nil, err
	}
	if err := c.publishSafeJSON(
		c.cfgTopic(entityCfg.ID, "binary_sensor"),
		&haconnectBinarySensor{
			haconnectEntityConfig: c.buildEntityConfig(entityCfg.ID),
			EntityConfig:          entityCfg,
			BinarySensorConfig:    cfg,
			Device:                c.device,
			Platform:              "binary_sensor",
			StateTopic:            stateTopic,
		},
	); err != nil {
		return nil, err
	}
	b := &BinarySensor{
		stateTopic: stateTopic,
	}
	if err := c.initEntity(b, entityCfg); err != nil {
		return nil, err
	}
	return b, nil
}
