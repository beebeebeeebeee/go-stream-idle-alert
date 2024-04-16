package app

type StreamData struct {
	Channel string
	Message string
}

type AlertSetting struct {
	Channel string
	Timeout int
	Enabled bool
}
