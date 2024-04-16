package app

import (
	"fmt"
	"log"
	"slices"
	"time"
)

func Stream(cb func(StreamData) error) error {
	// Simulate stream data
	for i := 0; i < 10; i++ {
		data := StreamData{
			Channel: fmt.Sprintf("Channel %d", i%2),
			Message: "Hello",
		}
		if err := cb(data); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	// Wait forever
	for {
		time.Sleep(1 * time.Second)
	}
}

type App struct {
	timerMap map[string]*time.Timer
}

func NewApp() *App {
	return &App{
		timerMap: make(map[string]*time.Timer),
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
					log.Printf("New settings delete timer: %s\n", channel)
					timer.Stop()
					delete(a.timerMap, channel)
				}
			}

			time.Sleep(5 * time.Second)
		}
	})()

	cb := func(data StreamData) error {
		log.Printf("message recieved: %v\n", data)

		idx := slices.IndexFunc(settings, func(setting AlertSetting) bool {
			return setting.Channel == data.Channel && setting.Enabled
		})
		if idx == -1 {
			return nil
		}
		setting := settings[idx]
		channel := setting.Channel
		duration := time.Duration(setting.Timeout) * time.Millisecond

		if timer, ok := a.timerMap[channel]; ok {
			log.Printf("Reset timer: %s\n", channel)
			timer.Reset(duration)
		} else {
			log.Printf("Create timer: %s\n", channel)
			timer := time.AfterFunc(duration, func() {
				log.Printf("Timeout: %s\n", channel)
				delete(a.timerMap, channel)
			})
			a.timerMap[channel] = timer
		}

		return nil
	}

	if err := Stream(cb); err != nil {
		log.Println(err)
	}
}
