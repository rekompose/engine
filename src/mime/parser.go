package mime

import (
	"fmt"
	"errors"

	"rekompose.com/engine/types"
)

// Parse will return the parsed mail from raw message
func Parse(message []byte) (*types.Message, error) {
	fmt.Printf("%v", message)
	// return &types.Message{Subject: "N/A", HTML: false}, nil
	return nil, errors.New("Could not parse the message")
}
