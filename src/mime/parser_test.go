package mime

import (
	"testing"
)

func TestParseShouldReturnErrorOnInvalidRawMessages(t *testing.T) {
	raw := "invalid message"

	_, err := Parse([]byte(raw))

	if err == nil {
		t.Errorf("Invalid messages should result in error!")
	}
}
