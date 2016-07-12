package scraper

import (
	"job-scraper/plugins/upwork"
	"job-scraper/plugins/welovegolang"
	"job-scraper/scraper"
)

// Executor represents the processor for the each job provider
type Executor struct {
	Plugin Scraper
}

// GetAllPlugins returns all the executable plugins
func GetAllPlugins() []Executor {
	return []scraper.Executor{
		scraper.Executor{
			Plugin: upwork.Upwork{},
		},
		scraper.Executor{
			Plugin: wlg.Wlg{},
		},
	}
}
