package types

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rekompose.com/engine/utils/logging"
)

var log = logging.NewLogger()

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
	Attachments []Attachment
	Inlines     []string
}

// Email is an proper email type as in John Doe <john@doe.org>
type Email struct {
	Sender  string
	Address string
}

// Attachment represents the attached file of a mime part
type Attachment struct {
	ContentType string
	FileName    string
	Content     []byte
}

// Text will give you full blown email address as text
func (e *Email) Text() string {
	return strings.TrimSpace(fmt.Sprintf("%s <%s>", e.Sender, e.Address))
}

// Save will save the content of the attachment with the filename of the attachment
func (a *Attachment) Save(path string) error {
	_ = os.Mkdir(path, 0700)
	file := filepath.Join(path, a.FileName)

	f, err := os.Create(file)
	if err != nil {
		log.Error("Cannot open attachment file to the disk ", file)
		return errors.New("cannot open attachment file to the disk")
	}
	defer f.Close()

	if _, err := f.Write(a.Content); err != nil {
		log.Error("Cannot write attachment file to the disk ", file)
		return errors.New("cannot write attachment file to the disk")
	}
	if err := f.Sync(); err != nil {
		log.Error("Cannot flush attachment file to the disk ", file)
		return errors.New("cannot flush attachment file to the disk")
	}

	return nil
}
