package upwork

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"job-scraper/manager/backend"
	"job-scraper/manager/model"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

const (
	provider  = "upwork"
	upworkURL = "https://www.upwork.com/ab/feed/jobs/rss?q=golang"
)

// Upwork plugin
type Upwork struct{}

// Run TODO
func (u Upwork) Run(facts chan model.Item, done chan bool) {
	fmt.Println("Running Upwork thread")
	res, err := http.Get(upworkURL)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	feed := model.Feed{}
	xml.Unmarshal(data, &feed)
	checkNewItems(feed.Channel.Items, facts)

}

// Checks for any new items and sends them to the email channel
func checkNewItems(items []model.Item, facts chan model.Item) {
	var lastItem model.LItem
	backend.LastItem.Find(bson.M{"provider": provider}).One(&lastItem)

	if items[0].GUID == lastItem.GUID {
		return
	}
	backend.LastItem.Update(bson.M{"provider": provider}, bson.M{"$set": bson.M{"guid": items[0].GUID}})
	facts <- items[0]

	for i := 1; i < len(items); i++ {
		if items[i].GUID == lastItem.GUID {
			break
		}

		facts <- items[i]
	}

	return
}
