package utils

import (
	"errors"
	"regexp"
	"strings"

	"rekompose.com/engine/types"
)

// Extract will extract sender information and sender email address from an email string
func Extract(s string) (types.Email, error) {
	var result types.Email
	re := regexp.MustCompile(`(?P<Sender>.*)\s*<(?P<Address>[a-zA-Z0-9_\.\+\-]+@[a-zA-Z0-9\-]+\.[a-zA-Z0-9\-\.]+)>`)

	matches := re.FindStringSubmatch(s)
	if matches == nil || len(matches) < 2 {
		return result, errors.New("not a standard email string pattern")
	}

	result.Sender = strings.Trim(strings.TrimSpace(matches[1]), "'|\"")
	result.Address = strings.Trim(strings.TrimSpace(matches[2]), "'|\"")

	return result, nil
}
