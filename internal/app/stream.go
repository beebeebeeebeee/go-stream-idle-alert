package app

import (
	"fmt"
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
