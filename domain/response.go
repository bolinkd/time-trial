package domain

import "github.com/paulmach/go.geojson"

type DevicesRequest struct {
	ChannelId         string                    `json:"channel_id"`
	FeatureCollection geojson.FeatureCollection `json:"feature_collection"`
}

type Response struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	ID       string `json:"id"`
}

type DevicesResponse struct {
	Name         string       `json:"name"`
	Devices      []string     `json:"devices"`
	CallbackInfo CallbackInfo `json:"cbinfo"`
}

type DeviceCountResponse struct {
	Name         string       `json:"name"`
	Count        int          `json:"count"`
	CallbackInfo CallbackInfo `json:"cbinfo"`
}

type CallbackInfo struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ID      string `json:"id"`
}
