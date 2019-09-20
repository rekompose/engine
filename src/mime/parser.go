package mime

import (
	"fmt"
	"rekompose.com/engine/types"
)

// Parse will return the parsed mail from raw message
func Parse(message []byte) types.Message {
	fmt.Printf("%v", message)
	return types.Message{Subject:"N/A", HTML:false}
}
