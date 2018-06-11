/*
Wrapper for the google wifi api.

*/
package gw_api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Status Object from Google Wifi router
type StatusData struct {
	Software SoftwareData `json:"software"`
	System   SystemData   `json:"system"`
	Wan      WanData      `json:"wan"`
}

// Software section of Status
type SoftwareData struct {
	SoftwareVersion  string  `json:"softwareVersion"`
	UpdateChannel    string  `json:"updateChannel"`
	UpdateNewVersion string  `json:"updateNewVersion"`
	UpdateProgress   float32 `json:"updateProgress"`
	UpdateRequired   bool    `json:"updateRequired"`
	UpdateStatus     string  `json:"updateStatus"`
}

// System section from Status
type SystemData struct {
	CountryCode string `json:"countryCode"`
	DeviceId    string `json:"deviceId"`
	GroupRole   string `json:"groupRole"`
	HardwareId  string `json:"hardwareId"`
	Lan0Link    bool   `json:"lan0Link"`
	ModelId     string `json:"modelId"`
	Update      int    `json:"uptime"`
}

// Wan section from Status
type WanData struct {
	CaptivePortal        bool     `json:"captivePortal"`
	EthernetLink         bool     `json:"ethernetLink"`
	GatewayIPAddress     string   `json:"gatewayIpAddress"`
	InvalidCredentials   bool     `json:"invalidCredentials"`
	IPAddress            bool     `json:"ipAddress"`
	IPMethod             string   `json:"ipMethod"`
	IPPrefixLength       int      `json:"ipPrefixLength"`
	LeaseDurationSeconds int      `json:"leaseDurationSeconds"`
	LocalIPAddress       string   `json:"localIpAddress"`
	NameServers          []string `json:"nameServers"`
	Online               bool     `json:"online"`
	PppoeDetected        bool     `json:"pppoeDetected"`
}

// Gets the gateway IP from Google Wifi
func GetStatus() (*StatusData, error) {
	var status StatusData
	rsp, err := http.Get("http://192.168.86.1/api/v1/status")
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &status)
	if err != nil {
		return nil, err
	}

	return &status, nil

}
