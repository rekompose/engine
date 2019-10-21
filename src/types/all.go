package types

import (
	"fmt"
	"strings"
)

// Message is a parsed mail
type Message struct {
	MimeVersion string
	MessageID   string
	Date        string
	From        Email
	To          []Email
	Cc          []Email
	Bcc         []Email
	Subject     string
	HTML        string
	Text        string
	Attachments []string
	Inlines     []string
}

// Email is an proper email type as in John Doe <john@doe.org>
type Email struct {
	Sender  string
	Address string
}

// Text will give you full blown email address as text
func (e *Email) Text() string {
	return strings.TrimSpace(fmt.Sprintf("%s <%s>", e.Sender, e.Address))
}
