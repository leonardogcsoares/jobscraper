package scraper

import "job-scraper/manager/model"

// Scraper represents the format every plugin must follow to be scraped
type Scraper interface {
	Run(chan model.Item, chan bool)
}
