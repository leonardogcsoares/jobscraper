package wlg

import (
	"fmt"
	"job-scraper/manager/backend"
	"job-scraper/manager/model"

	"github.com/PuerkitoBio/goquery"

	"gopkg.in/mgo.v2/bson"
)

const (
	provider        = "wlg"
	welovegolangURL = "http://www.welovegolang.com"
)

// Wlg TODO
type Wlg struct{}

// Run TODO
func (w Wlg) Run(facts chan model.Item, done chan bool) {
	doc, err := goquery.NewDocument(welovegolangURL)
	if err != nil {
		return
	}

	checkNewItems(doc, facts)
	fmt.Println("WeLoveGolang - Sleeping 1 day")
}

// Checks for any new items and sends them to the email channel
func checkNewItems(doc *goquery.Document, facts chan model.Item) {
	var lastItem model.LItem
	backend.LastItem.Find(bson.M{"provider": provider}).One(&lastItem)
	upToDate := false

	firstNode := doc.Find("div.stream-item.stream-job").First()
	if guid, exists := firstNode.Find(".media-heading a").Attr("href"); exists && guid != lastItem.GUID {
		backend.LastItem.Update(bson.M{"provider": provider},
			bson.M{"$set": bson.M{"guid": guid}})
	}

	doc.Find("div.stream-item.stream-job").Each(func(pos int, s *goquery.Selection) {
		var item model.Item
		guid, exists := s.Find(".media-heading a").Attr("href")
		if exists && guid == lastItem.GUID {
			upToDate = true
		}

		if !upToDate {
			item.GUID = guid
			item.Title = s.Find(".media-heading a > span").Text()
			item.Description = s.Find(".summary").Text()
			item.Link = fmt.Sprintf("%s%s", welovegolangURL, guid)
			facts <- item
		}
	})

	return
}
