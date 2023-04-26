package models

type ExperimentItem struct {
	Id   int    `json:"id,omitempty" db:"id"`
	RSSI int    `json:"rssi" db:"RSSI"`
	SSID string `json:"ssid" db:"ssid"`
	Time string `json:"time" db:"CurrentTime"`
}

type ExperimentData struct {
	Data []ExperimentItem `json:"data,omitempty"`
}
