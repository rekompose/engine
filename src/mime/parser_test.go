package mime

import (
	"testing"
)

func TestParseBase64ShouldAddMissingPadding(t *testing.T) {
	raw := "TUlNRS1WZXJzaW9uCg" //"MIME-Version" in base64 without the paddings

	_, err := ParseBase64(raw)

	if err != nil {
		t.Errorf("Missing padding should have been added!")
	}
}

func TestParseBase64ShouldReturnErrorOnInvalidEncoding(t *testing.T) {
	raw := "invalid encoding"

	_, err := ParseBase64(raw)

	if err == nil {
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
