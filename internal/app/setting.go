package app

import "fmt"

var settingVersion = 0

func GetSetting() []AlertSetting {
	settingVersion++
	return []AlertSetting{
		{
			Channel: fmt.Sprintf("Channel %d", settingVersion%2),
			Timeout: 3000,
			Enabled: true,
		},
	}
}
