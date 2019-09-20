package mime

import (
	"errors"
	"regexp"

	"rekompose.com/engine/types"
)

// Parse will return the parsed mail from raw message
func Parse(message []byte) (*types.Message, error) {
	match, _ := regexp.Match(`MIME\-Version`, message)

	if !match {
		return nil, errors.New("Could not parse the message")
	}

	return &types.Message{Subject: "N/A", HTML: false}, nil
}
