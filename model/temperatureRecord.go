package model

type TemperatureRecord struct {
	Location    string  `json:"location"`
	DeviceId    string  `json:"deviceId"`
	Temperature float32 `json:"temperature"`
}
