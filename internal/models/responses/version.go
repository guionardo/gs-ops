package responses

import "time"

type VersionResponse struct {
	AppName string    `json:"appnane"`
	Version string    `json:"version"`
	Date    time.Time `json:"date"`
}
