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
	mail := loadRawMessage("./data/mail1.json")

	message, _ := ParseBase64(mail)

	if message.Subject == "" {
		t.Errorf("Should parse `Subject` from the mime!")
	}
}

func TestParseBase64ShouldReadFrom(t *testing.T) {
	mail := loadRawMessage("./data/mail1.json")

	message, _ := ParseBase64(mail)

	if message.From.Address == "" {
		t.Errorf("Should parse `From` from the mime!")
	}
}

func TestParseBase64ShouldReadTo(t *testing.T) {
	mail := loadRawMessage("./data/mail1.json")

	message, _ := ParseBase64(mail)

	if len(message.To) == 0 {
		t.Errorf("Should parse `To` from the mime!")
	}
}

func TestParseBase64ShouldReadMultipleToAddresses(t *testing.T) {
	mail := loadRawMessage("./data/mail2.json")

	message, _ := ParseBase64(mail)

	if len(message.To) < 2 {
		t.Errorf("Should parse multiple `To` from the mime!")
	}
}

func TestParseBase64ShouldReadCc(t *testing.T) {
	mail := loadRawMessage("./data/mail3.json")

	message, _ := ParseBase64(mail)

	if len(message.Cc) == 0 {
		t.Errorf("Should parse `Cc` from the mime!")
	}
}

func TestParseBase64ShouldReadMultipleCcAddresses(t *testing.T) {
	mail := loadRawMessage("./data/mail4.json")

	message, _ := ParseBase64(mail)

	if len(message.Cc) < 2 {
		t.Errorf("Should parse multiple `Cc` from the mime!")
	}
}

func TestParseBase64ShouldReadDate(t *testing.T) {
	mail := loadRawMessage("./data/mail1.json")

	message, _ := ParseBase64(mail)

	if message.Date == "" {
		t.Errorf("Should parse `Date` from the mime!")
	}
}

func TestParseBase64ShouldReadText(t *testing.T) {
	mail := loadRawMessage("./data/mail1.json")

	message, _ := ParseBase64(mail)

	if message.Text == "" {
		t.Errorf("Should parse `Text` from the mime!")
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
