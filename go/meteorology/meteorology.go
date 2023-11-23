package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (u TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[u]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (u SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[u]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%v %v", s.magnitude, s.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (d MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", d.location, d.temperature.String(), d.windDirection, d.windSpeed.String(), d.humidity)
}
