package main

import (
	"job-scraper/manager/backend"
	"job-scraper/manager/model"
	"job-scraper/messenger/sender"
	"job-scraper/scraper"
	"time"
)

func main() {

	facts := make(chan model.Item, 4)
	done := make(chan bool)

	backend.Start()
	defer backend.Close()

	tasks := scraper.GetAllPlugins()

	for _, task := range tasks {
		go task.Plugin.Run(facts, done)
	}

	go sender.Run(facts)

	time.Sleep(time.Second * 60)
}
