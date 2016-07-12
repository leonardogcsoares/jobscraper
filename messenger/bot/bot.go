package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handler is responsible for handling send/receive messages from user.
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/webhook" {
			if r.URL.Query().Get("hub.verify_token") == ValidationToken {
				fmt.Fprintf(w, "%s", r.URL.Query().Get("hub.challenge"))
				return
			}
			return
		}
	}

	if r.Method == "POST" {
		if r.URL.Path == "/webhook" {
			fmt.Println("Message received")
			d, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("bad request, error %v", err), http.StatusBadRequest)
				return
			}

			fbe := MsgEntry{}
			err = json.Unmarshal(d, &fbe)
			if err != nil {
				fmt.Println(err)
			}

			for _, entry := range fbe.Entry {
				for _, message := range entry.Messaging {
					fmt.Println(message.Sender.ID)
				}
			}

		}

	}
}
