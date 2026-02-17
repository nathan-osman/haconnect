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

	// Identifiers provides a list of IDs that uniquely identify the device,
	// such as its MAC address(es).
	Identifiers []string

	// Manufacturer provides the manufacturer of the device.
	Manufacturer string

	// Model provides the model of the device.
	Model string

	// ModelID provides the model identifier of the device.
	ModelID string

	// HWVersion provides the hardware version of the device.
	HWVersion string

	// SWVersion provides the software version of the device.
	SWVersion string

	// SerialNumber provides the serial number of the device.
	SerialNumber string

	// SuggestedArea provides the area where the device is located.
	SuggestedArea string
}
