package bot

import (
	"encoding/json"
	"fmt"
	"job-scraper/manager/model"
)

// MountJobMessage takes the Job item and creates the approriate response message
func MountJobMessage(job model.Item, recipientID string) []byte {

	message := GenericMessage{
		Recipient: Recipient{
			ID: recipientID,
		},
		Message: Message{
			Attachment: Attachment{
				Type: "template",
				Payload: Payload{
					TemplateType: "generic",
					Elements: []Element{
						Element{
							Title:    job.Title,
							ImageURL: "https://golang.org/doc/gopher/gopherbw.png",
							Subtitle: fmt.Sprintf("%s...", job.Description[:100]),
							Buttons: []Button{
								Button{
									Type:  "web_url",
									Title: "Open Url",
									URL:   job.Link,
								},
							},
						},
					},
				},
			},
		},
	}

	response, _ := json.Marshal(message)

	return response
}
