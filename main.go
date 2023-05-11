package main

import (
	"time"

	C "github.com/parsifal-m/pinggo/monitor"
)

func main() {
	// Websites we want to monitor

	websites := []string{
		"http://google.com",
	}

	for _, website := range websites {
		go func(website string) {
			for {
				C.CheckWebsite(website)
				time.Sleep(60 * time.Second)
			}
		}(website)
	}

	select {}

}
