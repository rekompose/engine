package mime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"unicode"
)

func TestParseBase64ShouldReturnErrorOnInvalidEncoding(t *testing.T) {
	raw := "invalid encoding"

	_, err := ParseBase64(raw)

	if err == nil {
		t.Errorf("Invalid encoding should result in error!")
	}
}

func TestParseBase64ShouldReturnMessageWhenValidEncodingGiven(t *testing.T) {
	raw := "TUlNRS1WZXJzaW9uCg==" // "MIME-Version" in base64

	result, _ := ParseBase64(raw)
	resultType := fmt.Sprintf("%T", result)

	if resultType != "types.Message" {
		t.Errorf("Invalid encoding should result in error!")
	}
}

func TestParseShouldReturnErrorOnInvalidRawMessages(t *testing.T) {
	raw := "invalid message"

	_, err := Parse([]byte(raw))

	if err == nil {
		t.Errorf("Invalid messages should result in error!")
	}
}

func TestParseBase64ShouldReadSubject(t *testing.T) {
	mail := loadRawMessage("./data/simple.json")

	message, _ := ParseBase64(mail)

	if message.Subject == "" {
		t.Errorf("Should parse `Subject` from the mime!")
	}
}

func TestParseBase64ShouldReadFrom(t *testing.T) {
	mail := loadRawMessage("./data/simple.json")

	message, _ := ParseBase64(mail)

	if message.From.Address == "" {
		t.Errorf("Should parse `From` from the mime when sender and email address present!")
	}
}

func TestParseBase64ShouldReadFromWhenOnlyEmailAddressPresent(t *testing.T) {
	mail := loadRawMessage("./data/no-sender-in-from.json")

	message, _ := ParseBase64(mail)

	if message.From.Address == "" {
		t.Errorf("Should parse `From` from the mime when only email address present!")
	}
}

func TestParseBase64ShouldReadTo(t *testing.T) {
	mail := loadRawMessage("./data/simple.json")

	message, _ := ParseBase64(mail)

	if len(message.To) == 0 {
		t.Errorf("Should parse `To` from the mime!")
	}
}

func TestParseBase64ShouldReadMultipleToAddresses(t *testing.T) {
	mail := loadRawMessage("./data/multiple-to-addresses.json")

	message, _ := ParseBase64(mail)

	if len(message.To) < 2 {
		t.Errorf("Should parse multiple `To` from the mime!")
	}
}

func TestParseBase64ShouldReadCc(t *testing.T) {
	mail := loadRawMessage("./data/no-sender-in-from.json")

	message, _ := ParseBase64(mail)

	if len(message.Cc) == 0 {
		t.Errorf("Should parse `Cc` from the mime!")
	}
}

func TestParseBase64ShouldReadMultipleCcAddresses(t *testing.T) {
	mail := loadRawMessage("./data/multiple-cc-addresses.json")

	message, _ := ParseBase64(mail)

	if len(message.Cc) < 2 {
		t.Errorf("Should parse multiple `Cc` from the mime!")
	}
}

func TestParseBase64ShouldReadDate(t *testing.T) {
	mail := loadRawMessage("./data/simple.json")

	message, _ := ParseBase64(mail)

	if message.Date == "" {
		t.Errorf("Should parse `Date` from the mime!")
	}
}

func TestParseBase64ShouldReadText(t *testing.T) {
	mail := loadRawMessage("./data/simple.json")

	message, _ := ParseBase64(mail)

	if message.Text == "" {
		t.Errorf("Should parse `Text` from the mime!")
	}
}

func TestParseBase64ShouldReadAttachments(t *testing.T) {
	mail := loadRawMessage("./data/no-sender-in-from.json")

	message, _ := ParseBase64(mail)

	if len(message.Attachments) == 0 {
		t.Errorf("Should parse all the `attachments` from the mime!")
	}
}

func TestParseBase64ShouldReadAttachedPDF(t *testing.T) {
	mail := loadRawMessage("./data/no-sender-in-from.json")

	message, _ := ParseBase64(mail)
	pdf := message.Attachments[0]

	if !strings.HasSuffix(strings.ToLower(pdf.FileName), "pdf") {
		t.Errorf("Should parse all the `attachments` including PDF ones from the mime!")
	}
}

func TestParseBase64ShouldEmbedImagesInHTMLContent(t *testing.T) {
	mail := loadRawMessage("./data/embed-images-to-html.json")

	message, _ := ParseBase64(mail)

	if strings.Contains(message.HTML, "cid:") {
		t.Errorf(message.HTML)
		t.Errorf("Should embed all the images in HTML within the mime parts!")
	}
}

func loadRawMessage(filename string) string {
	json := loadJSONFile(filename)
	return strings.TrimRightFunc(json["raw"].(string), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func loadJSONFile(filename string) map[string]interface{} {
	file, _ := os.Open(filename)
	defer file.Close()

	mailContent, _ := ioutil.ReadAll(file)
	var result map[string]interface{}
	json.Unmarshal([]byte(mailContent), &result)

	return result
}
