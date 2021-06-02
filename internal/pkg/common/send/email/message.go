package email

// Message defines a message to be sent.
type Message struct {
	subject, body string
	contentType   string
}

// NewTextMessage creates message in plain text.
func NewTextMessage(subject, body string) Message {
	return Message{subject, body, ""}
}

// NewHTMLMessage creates message in HTML format.
func NewHTMLMessage(subject, body string) Message {
	return Message{subject, body, "text/html"}
}
