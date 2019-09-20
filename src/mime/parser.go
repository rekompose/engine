package mime

import (
	"encoding/base64"
	"errors"
	"regexp"

	"rekompose.com/engine/types"
)

// ParseBase64 will return the parsed mail from a raw message encoded in Base64
func ParseBase64(message string) (*types.Message, error) {
	message = addMissingPaddings(message)

	raw, err := base64.StdEncoding.DecodeString(message)
	if err == nil {
		return Parse(raw)
	}

	return nil, errors.New("Could not parse, potential encoding error")
}

// Parse will return the parsed mail from a raw message
func Parse(message []byte) (*types.Message, error) {
	match, _ := regexp.Match(`MIME\-Version`, message)

	if !match {
		return nil, errors.New("Could not parse the message")
	}

	return &types.Message{Subject: "N/A", HTML: false}, nil
}

func addMissingPaddings(message string) string {
	mod := len(message) % 4

	if mod != 0 {
		for i := 0; i < mod; i++ {
			message = message + "="
		}
	}

	return message
}
