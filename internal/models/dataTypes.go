package models

type ExperimentItem struct {
	Id   int    `json:"id,omitempty" db:"id"`
	RSSI int    `json:"rssi" db:"RSSI"`
	SSID string `json:"ssid" db:"SSID"`
	Time string `json:"time" db:"CurrentTime"`
}

type ExperimentData struct {
	Data []ExperimentItem `json:"data,omitempty"`
}

type SuccessCauseResponse struct {
	Success bool   `json:"success,omitempty"`
	Cause   string `json:"cause,omitempty"`
}
