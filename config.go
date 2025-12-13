package haconnect

// Config provides configuration for instantiating Conn instances.
type Config struct {

	// Addr provides the address of the MQTT broker to connect to.
	Addr string

	// Username provides the username to use for authenticating to the broker.
	Username string

	// Password provides the password to use for authenticating to the broker.
	Password string

	// DiscoveryPrefix provides the MQTT discovery prefix to use. If left
	// empty, this defaults to "homeassistant".
	DiscoveryPrefix string

	// ID provides the unique ID of the device. If left empty, the current
	// hostname (after removing invalid characters) is used.
	ID string

	// Name provides the name of the device. If left empty, the value of ID is
	// used.
	Name string

	// Manufacturer provides the manufacturer of the device.
	Manufacturer string

	// Model provides the model of the device
	Model string
}
