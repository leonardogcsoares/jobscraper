package scraper

import (
	"job-scraper/plugins/upwork"
	"job-scraper/plugins/welovegolang"
)

// Executor represents the processor for the each job provider
type Executor struct {
	Plugin Scraper
}

// GetAllPlugins returns all the executable plugins
func GetAllPlugins() []Executor {
	return []Executor{
		Executor{
			Plugin: upwork.Upwork{},
		},
		Executor{
			Plugin: wlg.Wlg{},
		},
	}
}
