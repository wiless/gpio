package main

import (
	"fmt"
	"time"

	"github.com/wiless/gpio"
)

func main() {
	fmt.Println("Watching PIN 18")
	watcher := gpio.NewWatcher()
	watcher.AddPin(18)
	defer watcher.Close()
	var lastPressed time.Time
	var nullTime time.Time
	func() {
		for {
			pin, value := watcher.Watch()
			if value == 1 && pin == 18 {
				lastPressed = time.Now()
				fmt.Printf("\nPressed GPIO %d", pin)
			} else {
				if lastPressed != nullTime {

					dur := time.Since(lastPressed)
					fmt.Printf("... for %v", dur)
				}
			}
		}
	}()
}
