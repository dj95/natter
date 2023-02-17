package main

import (
	"time"

	"github.com/dj95/natter/internal/config"
	"github.com/dj95/natter/internal/listen"
)

func init() {
	config.InitializeCommandFlags()
}

func main() {
    filter := config.Filter()

	for _, intrfc := range config.Interfaces() {
        listen.OnInterface(intrfc, filter)
	}

	for {
		time.Sleep(60 * time.Second)
	}
}

