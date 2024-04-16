package app

import "time"

type StreamData struct {
	Channel string
	Message string
}

type AlertSetting struct {
	Channel string
	Timeout int
	Enabled bool
}

type AlertTimer struct {
	Channel string
	Timer   *time.Timer
	IsDone  bool
}
