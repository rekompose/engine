package mime

import (
	"bytes"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/jhillyerd/enmime"
	"rekompose.com/engine/types"
	"rekompose.com/engine/utils"
)

var log = utils.NewLogger()

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
		From:        extractSingleEmailAddress(env.GetHeader("From")),
		To:          extractMultipleEmailAddresses(env.GetHeader("To")),
		Cc:          extractMultipleEmailAddresses(env.GetHeader("Cc")),
		Bcc:         extractMultipleEmailAddresses(env.GetHeader("Bcc")),
		Subject:     env.GetHeader("Subject"),
		Text:        env.Text,
		HTML:        env.HTML,
	}, nil
}

// this might be useful later once integrated with various mail providers
func addMissingPaddings(raw string) string {
	if m := len(raw) % 4; m != 0 {
		raw += strings.Repeat("=", 4-m)
		log.Debug("Adding missing padding for base64 raw string!")
	}

	return raw
}

// extract single email address from a string
func extractSingleEmailAddress(s string) types.Email {
	address, err := utils.Extract(s)
	if err != nil {
		log.Error("invalid email adress detected", "address", s)
	}

	return address
}

// extract email addresses from comma separated list
func extractMultipleEmailAddresses(s string) []types.Email {
	log.Debug("extraction of emails ", "incoming text", s)
	var list []types.Email
	filter := func(c rune) bool {
		return c == ','
	}

	for _, email := range strings.FieldsFunc(s, filter) {
		match, err := utils.Extract(email)
		if err == nil {
			list = append(list, match)
		}
	}

	log.Debug("extracted ", "emails", list)
	return list
}
