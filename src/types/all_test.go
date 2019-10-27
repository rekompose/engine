package types

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestTextShouldCombineSenderInformationWithEmail(t *testing.T) {
	email := Email{Sender: "John Doe", Address: "john@doe.org"}
	expected := "John Doe <john@doe.org>"

	result := email.Text()

	if result != expected {
		t.Errorf("Should combine sender information with email string")
	}
}

func TestSaveShouldCreateAFileWithAttachmentContent(t *testing.T) {
	content := []byte("Content is the king")
	attachment := Attachment{
		FileName:    "attachment.txt",
		ContentType: "text/html",
		Content:     content,
	}

	dir, _ := ioutil.TempDir("", "test")
	defer os.RemoveAll(dir)
	_ = attachment.Save(dir)

	fileContent, _ := ioutil.ReadFile(filepath.Join(dir, attachment.FileName))
	if bytes.Compare(fileContent, content) != 0 {
		t.Errorf("Should create a file with the attachment content!")
	}
}
