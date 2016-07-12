package bot

// Sender TODO
type Sender struct {
	ID string `json:"id"`
}

// Recipient TODO
type Recipient struct {
	ID string `json:"id"`
}

// Messaging TODO
type Messaging struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	TS        int64     `json:"timestamp"`
	Message   Message   `json:"message,omitempty"`
	Postback  Postback  `json:"postback,omitempty"`
	Optin     Optin     `json:"optin,omitempty"`
}

// Optin TODO
type Optin struct {
	Ref string `json:"ref,omitempty"`
}

// Postback TODO
type Postback struct {
	Payload string `json:"payload,omitempty"`
}

// Entry TODO
type Entry struct {
	Messaging []Messaging `json:"messaging"`
}

// MsgEntry TODO
type MsgEntry struct {
	Entry []Entry `json:"entry"`
}

// SendMsgEntry TODO
type SendMsgEntry struct {
	Recipient Recipient `json:"recipient"`
	Message   Message   `json:"message"`
}

// GenericMessage Format
type GenericMessage struct {
	Recipient Recipient `json:"recipient"`
	Message   Message   `json:"message"`
}

// Message Format for Sending Button
type Message struct {
	Attachment Attachment `json:"attachment"`
}

// Attachment for sending button messages
type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

// Payload for sending button messages
type Payload struct {
	TemplateType string    `json:"template_type"`
	Text         string    `json:"text,omitempty"`
	Buttons      []Button  `json:"buttons,omitempty"`
	Elements     []Element `json:"elements,omitempty"`
}

// Element TODO
type Element struct {
	Title    string   `json:"title"`
	ImageURL string   `json:"image_url"`
	Subtitle string   `json:"subtitle"`
	Buttons  []Button `json:"buttons"`
}

// Button for sending button messages
type Button struct {
	Type    string `json:"type"`
	URL     string `json:"url,omitempty"`
	Title   string `json:"title"`
	Payload string `json:"payload,omitempty"`
}
