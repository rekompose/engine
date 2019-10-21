package mime

import (
	"bytes"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/jhillyerd/enmime"
	"rekompose.com/engine/types"
)

// ParseBase64 will return the parsed mail from a raw message encoded in Base64
func ParseBase64(raw string) (types.Message, error) {
	b, err := base64.URLEncoding.DecodeString(addMissingPaddings(raw))
	if err == nil {
		return Parse(b)
	}

	return types.Message{}, errors.New("Could not parse, potential encoding error")
}

// Parse will return the parsed mail from a raw message
func Parse(raw []byte) (types.Message, error) {
	reader := bytes.NewReader(raw)
	env, err := enmime.ReadEnvelope(reader)
	if err != nil {
		return types.Message{}, errors.New("Could not parse the message")
	}

	return types.Message{
		MimeVersion: env.GetHeader("Mime-Version"),
		MessageID:   env.GetHeader("Message-Id"),
		Date:        env.GetHeader("Date"),
		From:        env.GetHeader("From"),
		To:          extractEmailAddresses(env.GetHeader("To")),
		Cc:          extractEmailAddresses(env.GetHeader("Cc")),
		Bcc:         extractEmailAddresses(env.GetHeader("Bcc")),
		Subject:     env.GetHeader("Subject"),
		Text:        env.Text,
		HTML:        env.HTML,
	}, nil
}

// this might be useful later once integrated with various mail providers
func addMissingPaddings(raw string) string {
	if m := len(raw) % 4; m != 0 {
		raw += strings.Repeat("=", 4-m)
	}

	return raw
}

// extract email addresses from comma separated list
func extractEmailAddresses(s string) []string {
	var list []string

	list = strings.Split(s, ",")
	return list
}
