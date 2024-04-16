package app

import (
	"log"
	"slices"
	"time"
)

type App struct {
	timerMap map[string]*AlertTimer
}

func NewApp() *App {
	return &App{
		timerMap: make(map[string]*AlertTimer),
	}
}

func (a *App) Run() {
	var settings []AlertSetting
	go (func() {
		for {
			settings = GetSetting()
			log.Printf("New settings: %v\n", settings)

			for channel, timer := range a.timerMap {
				idx := slices.IndexFunc(settings, func(setting AlertSetting) bool {
					return setting.Channel == channel && setting.Enabled
				})

				if idx == -1 {
					log.Printf("New settings deleted timer: %s\n", channel)
					timer.Timer.Stop()
					delete(a.timerMap, channel)
				}
			}

			for _, setting := range settings {
				if !setting.Enabled {
					continue
				}

				channel := setting.Channel
				duration := time.Duration(setting.Timeout) * time.Millisecond
				if _, ok := a.timerMap[channel]; !ok {
					log.Printf("New settings created timer: %s\n", channel)
					timer := time.AfterFunc(duration, func() {
						log.Printf("Timeout: %s\n", channel)
						a.timerMap[channel].IsDone = true
					})
					a.timerMap[channel] = &AlertTimer{
						Channel: channel,
						Timer:   timer,
						IsDone:  false,
					}
				}
			}

			time.Sleep(5 * time.Second)
		}
	})()

	cb := func(data StreamData) error {
		log.Printf("message recieved: %v\n", data)

		if timer, ok := a.timerMap[data.Channel]; ok {
			log.Printf("Reset timer: %s\n", data.Channel)
			timer.Timer.Reset(time.Duration(settings[0].Timeout) * time.Millisecond)
		}

		return nil
	}

	if err := Stream(cb); err != nil {
		log.Println(err)
	}
}
