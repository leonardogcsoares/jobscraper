package sender

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"job-scraper/manager/model"
	"job-scraper/messenger/bot"
	"net/http"
)

const (
	userLeoID = "1008475132592764"
)

// Run TODO
func Run(facts chan model.Item) {
	fmt.Println("Running Messenger thread")
	URL := fmt.Sprintf("%s/%s/messages?access_token=%s", bot.MessengerAPI, bot.BotPageID, bot.AccessToken)

	for item := range facts {

		msg := bot.MountJobMessage(item, userLeoID)
		req, _ := http.NewRequest("POST",
			URL,
			bytes.NewBuffer(msg),
		)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// Parse response from body and print.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println("Response body: ", string(body))

	}
}
