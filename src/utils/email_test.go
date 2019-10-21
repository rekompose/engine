package utils

import (
	"fmt"
	"testing"
)

func TestExtractSenderFromString(t *testing.T) {
	name := "John Doe"
	test := fmt.Sprintf("%s <john@doe.org>", name)

	result, _ := Extract(test)

	if result.Sender != name {
		t.Errorf("Should extract sender information from email string")
	}
}

func TestExtractShouldRemoveQuotesFromString(t *testing.T) {
	before := "\"'John Doe\"'"
	after := "John Doe"
	test := fmt.Sprintf("%s <john@doe.org>", before)

	result, _ := Extract(test)

	if result.Sender != after {
		t.Errorf("Should remove all the quotes from the sender information in the email string")
	}
}

func TestExtractEmailAddressFromString(t *testing.T) {
	emailAddress := "john@doe.org"
	test := fmt.Sprintf("John Doe <%s>", emailAddress)

	result, _ := Extract(test)

	if result.Address != emailAddress {
		t.Errorf("Should extract sender email address from email string")
	}
}

func TestExtractShouldDisplayFullEmailAddressFromString(t *testing.T) {
	address := "John Doe <john@doe.org>"
	test := fmt.Sprintf(address)

	result, _ := Extract(test)

	if result.Text() != address {
		t.Errorf("Should display full email address from email string")
	}
}

func TestExtractShouldDisplayFullEmailAddressFromStringWhenMissingSenderInformation(t *testing.T) {
	address := "<john@doe.org>"
	test := fmt.Sprintf(address)

	result, _ := Extract(test)

	if result.Text() != address {
		t.Errorf("Should display full email address from email string")
	}
}

func TestExtractReturnErrorWhenNoEmailAddressFound(t *testing.T) {
	address := "Random Gibberish <abc>"
	test := fmt.Sprintf(address)

	_, err := Extract(test)

	if err == nil {
		t.Errorf("Should throw an error when no email address detected")
	}
}
