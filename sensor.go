package hamqtt

const (
	SensorAbsoluteHumidity              = "absolute_humidity"
	SensorApparentPower                 = "apparent_power"
	SensorAQI                           = "aqi"
	SensorArea                          = "area"
	SensorAtmosphericPressure           = "atmospheric_pressure"
	SensorBattery                       = "battery"
	SensorBloodGlucoseConcentration     = "blood_glucose_concentration"
	SensorCarbonDioxide                 = "carbon_dioxide"
	SensorCarbonMonoxide                = "carbon_monoxide"
	SensorCurrent                       = "current"
	SensorDataRate                      = "data_rate"
	SensorDataSize                      = "data_size"
	SensorDate                          = "date"
	SensorDistance                      = "distance"
	SensorDuration                      = "duration"
	SensorEnergy                        = "energy"
	SensorEnergyDistance                = "energy_distance"
	SensorEnergyStorage                 = "energy_storage"
	SensorEnum                          = "enum"
	SensorFrequency                     = "frequency"
	SensorGas                           = "gas"
	SensorHumidity                      = "humidity"
	SensorIlluminance                   = "illuminance"
	SensorIrradiance                    = "irradiance"
	SensorMoisture                      = "moisture"
	SensorMonetary                      = "monetary"
	SensorNitrogenDioxide               = "nitrogen_dioxide"
	SensorNitrogenMonoxide              = "nitrogen_monoxide"
	SensorNitrousOxide                  = "nitrous_oxide"
	SensorOzone                         = "ozone"
	SensorPH                            = "ph"
	SensorPM1                           = "pm1"
	SensorPM25                          = "pm25"
	SensorPM4                           = "pm4"
	SensorPM10                          = "pm10"
	SensorPowerFactor                   = "power_factor"
	SensorPower                         = "power"
	SensorPrecipitation                 = "precipitation"
	SensorPrecipitationIntensity        = "precipitation_intensity"
	SensorPressure                      = "pressure"
	SensorReactiveEnergy                = "reactive_energy"
	SensorReactivePower                 = "reactive_power"
	SensorSignalStrength                = "signal_strength"
	SensorSoundPressure                 = "sound_pressure"
	SensorSpeed                         = "speed"
	SensorSulphurDioxide                = "sulphur_dioxide"
	SensorTemperature                   = "temperature"
	SensorTimestamp                     = "timestamp"
	SensorVolatileOrganicCompounds      = "volatile_organic_compounds"
	SensorVolatileOrganicCompoundsParts = "volatile_organic_compounds_parts"
	SensorVoltage                       = "voltage"
	SensorVolume                        = "volume"
	SensorVolumeFlowRate                = "volume_flow_rate"
	SensorVolumeStorage                 = "volume_storage"
	SensorWater                         = "water"
	SensorWeight                        = "weight"
	SensorWindDirection                 = "wind_direction"
	SensorWindSpeed                     = "wind_speed"

	SensorMeasurement      = "measurement"
	SensorMeasurementAngle = "measurement_angle"
	SensorTotal            = "total"
	SensorTotalIncreasing  = "total_increasing"

	SensorAcre                          = "ac"
	SensorAmpere                        = "A"
	SensorAWeightedDecibel              = "dBA"
	SensorBar                           = "bar"
	SensorBit                           = "bit"
	SensorBitsPerSecond                 = "bit/s"
	SensorBTUPerHourPerSquareFoot       = "BTU/(h⋅ft²)"
	SensorByte                          = "B"
	SensorBytesPerSecond                = "B/s"
	SensorCalorie                       = "cal"
	SensorCentibar                      = "cbar"
	SensorCentimeter                    = "cm"
	SensorCubicFeetPerMinute            = "ft³/min"
	SensorCubicFoot                     = "ft³"
	SensorCubicMeter                    = "m³"
	SensorCubicMetersPerHour            = "m³/h"
	SensorCubicMetersPerMinute          = "m³/min"
	SensorCubicMetersPerSecond          = "m³/s"
	SensorDay                           = "d"
	SensorDecibel                       = "dB"
	SensorDecibelMilliwatt              = "dBm"
	SensorDegrees                       = "°"
	SensorDegreesCelsius                = "°C"
	SensorDegreesFahrenheit             = "°F"
	SensorExabyte                       = "EB"
	SensorExbibyte                      = "EiB"
	SensorFeetPerSecond                 = "ft/s"
	SensorFluidOunce                    = "fl. oz."
	SensorGallon                        = "gal"
	SensorGallonsPerHour                = "gal/h"
	SensorGallonsPerMinute              = "gal/min"
	SensorGibibyte                      = "GiB"
	SensorGibibytesPerSecond            = "GiB/s"
	SensorGigabit                       = "Gbit"
	SensorGigabitsPerSecond             = "Gbit/s"
	SensorGigabyte                      = "GB"
	SensorGigabytesPerSecond            = "GB/s"
	SensorGigacalorie                   = "Gcal"
	SensorGigahertz                     = "GHz"
	SensorGigajoule                     = "GJ"
	SensorGigawatt                      = "GW"
	SensorGigawattHour                  = "GWh"
	SensorGram                          = "g"
	SensorGramsPerCubicMeter            = "g/m³"
	SensorHectare                       = "ha"
	SensorHectopascal                   = "hPa"
	SensorHertz                         = "Hz"
	SensorHour                          = "h"
	SensorHundredCubicFeet              = "CCF"
	SensorInch                          = "in"
	SensorInchesOfMercury               = "inHg"
	SensorInchesOfWater                 = "inH₂O"
	SensorInchesPerDay                  = "in/d"
	SensorInchesPerHour                 = "in/h"
	SensorInchesPerSecond               = "in/s"
	SensorJoule                         = "J"
	SensorKelvin                        = "K"
	SensorKibibyte                      = "KiB"
	SensorKibibytesPerSecond            = "KiB/s"
	SensorKilobit                       = "kbit"
	SensorKilobitsPerSecond             = "kbit/s"
	SensorKilobyte                      = "kB"
	SensorKilobytesPerSecond            = "kB/s"
	SensorKilocalorie                   = "kcal"
	SensorKilogram                      = "kg"
	SensorKilohertz                     = "kHz"
	SensorKilojoule                     = "kJ"
	SensorKilometer                     = "km"
	SensorKilometersPerHour             = "km/h"
	SensorKilometersPerKilowattHour     = "km/kWh"
	SensorKilopascal                    = "kPa"
	SensorKilovar                       = "kvar"
	SensorKilovarHour                   = "kvarh"
	SensorKilovolt                      = "kV"
	SensorKilovoltAmpere                = "kVA"
	SensorKilowatt                      = "kW"
	SensorKilowattHour                  = "kWh"
	SensorKilowattHoursPer100Kilometers = "kWh/100km"
	SensorKnot                          = "kn"
	SensorLiter                         = "L"
	SensorLitersPerHour                 = "L/h"
	SensorLitersPerMinute               = "L/min"
	SensorLitersPerSecond               = "L/s"
	SensorLux                           = "lx"
	SensorMebibyte                      = "MiB"
	SensorMebibytesPerSecond            = "MiB/s"
	SensorMegabit                       = "Mbit"
	SensorMegabitsPerSecond             = "Mbit/s"
	SensorMegabyte                      = "MB"
	SensorMegabytesPerSecond            = "MB/s"
	SensorMegacalorie                   = "Mcal"
	SensorMegahertz                     = "MHz"
	SensorMegajoule                     = "MJ"
	SensorMegavar                       = "mvar"
	SensorMegavolt                      = "MV"
	SensorMegawatt                      = "MW"
	SensorMegawattHour                  = "MWh"
	SensorMeter                         = "m"
	SensorMetersPerSecond               = "m/s"
	SensorMicrogram                     = "µg"
	SensorMicrogramsPerCubicCentimeter  = "µg/m³"
	SensorMicrogramsPerCubicMeter       = "µg/m³"
	SensorMicrosecond                   = "µs"
	SensorMicrosiemensPerCentimeter     = "µS/cm"
	SensorMicrovolt                     = "µV"
	SensorMile                          = "mi"
	SensorMilesPerHour                  = "mph"
	SensorMilesPerKilowattHour          = "mi/kWh"
	SensorMilliampere                   = "mA"
	SensorMillibar                      = "mbar"
	SensorMilligram                     = "mg"
	SensorMilligramsPerCubicMeter       = "mg/m³"
	SensorMilligramsPerDeciliter        = "mg/dL"
	SensorMilliliter                    = "mL"
	SensorMillilitersPerSecond          = "mL/s"
	SensorMillimeter                    = "mm"
	SensorMillimetersOfMercury          = "mmHG"
	SensorMillimetersPerDay             = "mm/d"
	SensorMillimetersPerHour            = "mm/h"
	SensorMillimetersPerSecond          = "mm/s"
	SensorMillimolesPerLiter            = "mmol/L"
	SensorMillisecond                   = "ms"
	SensorMillisiemensPerCentimeter     = "mS/cm"
	SensorMillivolt                     = "mV"
	SensorMillivoltAmpere               = "mVA"
	SensorMilliwatt                     = "mW"
	SensorMilliwattHour                 = "mWh"
	SensorMinute                        = "min"
	SensorNauticalMile                  = "nmi"
	SensorOunce                         = "oz"
	SensorPartsPerBillion               = "ppb"
	SensorPartsPerMillion               = "ppm"
	SensorPascal                        = "Pa"
	SensorPebibyte                      = "PiB"
	SensorPercentage                    = "%"
	SensorPetabyte                      = "PB"
	SensorPound                         = "lb"
	SensorPoundsPerSquareInch           = "psi"
	SensorSecond                        = "s"
	SensorSiemensPerCentimeter          = "S/cm"
	SensorSquareCentimeter              = "cm²"
	SensorSquareFoot                    = "ft²"
	SensorSquareInch                    = "in²"
	SensorSquareKilometer               = "km²"
	SensorSquareMeter                   = "m²"
	SensorSquareMile                    = "mi²"
	SensorSquareYard                    = "yd²"
	SensorStone                         = "st"
	SensorTebibyte                      = "TiB"
	SensorTerabyte                      = "TB"
	SensorTerawatt                      = "TW"
	SensorTerawattHour                  = "TWh"
	SensorThousandCubicFeet             = "MCF"
	SensorVolt                          = "V"
	SensorVoltAmpere                    = "VA"
	SensorVoltAmpereReactive            = "var"
	SensorVoltAmpereReactiveHour        = "varh"
	SensorWatt                          = "W"
	SensorWattHour                      = "Wh"
	SensorWattHoursPerKilometer         = "Wh/km"
	SensorWattsPerSquareMeter           = "W/m²"
	SensorYard                          = "yd"
	SensorYobibyte                      = "YiB"
	SensorYottabyte                     = "YB"
	SensorZebibyte                      = "ZiB"
	SensorZettabyte                     = "ZB"
)

// SensorConfig provides configuration for sensors.
type SensorConfig struct {

	// DeviceClass categorizes the type of data reported by the sensor.
	DeviceClass string `json:"device_class,omitempty"`

	// StateClass categorizes the nature of the data reported.
	StateClass string `json:"state_class,omitempty"`

	// Options is a list of allowed values if DeviceClass is SensorEnum.
	Options []string `json:"options,omitempty"`

	// UnitOfMeasurement indicates the unit for the sensor's values.
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`

	// SuggestedDisplayPrecision indicates the number of decimal places that
	// should be used to display the sensor's value.
	SuggestedDisplayPrecision int `json:"suggested_display_precision,omitempty"`
}

type hamqttSensor struct {
	*hamqttEntityConfig
	*EntityConfig
	*SensorConfig
	Device     *hamqttDevice `json:"device"`
	Platform   string        `json:"platform"`
	StateTopic string        `json:"state_topic"`
}

// Sensor provides methods for indicating changes to the sensor.
type Sensor struct {
	Entity
	stateTopic string
}

// Set updates the sensor's value.
func (s *Sensor) Set(value string) error {
	return s.conn.publishSafe(s.stateTopic, value)
}

// Sensor creates a new entity that represents a sensor.
func (c *Conn) Sensor(
	entityCfg *EntityConfig,
	cfg *SensorConfig,
) (*Sensor, error) {
	stateTopic := c.stateTopic(entityCfg.ID)
	if err := c.publishSafeJSON(
		c.cfgTopic(entityCfg.ID, "sensor"),
		&hamqttSensor{
			hamqttEntityConfig: c.buildEntityConfig(entityCfg.ID),
			EntityConfig:       entityCfg,
			SensorConfig:       cfg,
			Device:             c.device,
			Platform:           "sensor",
			StateTopic:         stateTopic,
		},
	); err != nil {
		return nil, err
	}
	s := &Sensor{
		stateTopic: stateTopic,
	}
	if err := c.initEntity(s, entityCfg); err != nil {
		return nil, err
	}
	return s, nil
}
