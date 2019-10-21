package types

// Message is a parsed mail
type Message struct {
	MimeVersion string
	MessageID   string
	Date        string
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	HTML        string
	Text        string
	Attachments []string
	Inlines     []string
}
