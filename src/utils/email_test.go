package utils

import (
	"fmt"
	"testing"
)

func TestExtractSenderInformation(t *testing.T) {
	name := "John Doe"
	test := fmt.Sprintf("%s <john@doe.org>", name)

	result, _ := Extract(test)

	if result.Sender != name {
		t.Errorf("Should extract sender information from email string")
	}
}

func TestExtractShouldTrimQuotes(t *testing.T) {
	before := "\"'John Doe\"'"
	after := "John Doe"
	test := fmt.Sprintf("%s <john@doe.org>", before)

	result, _ := Extract(test)

	if result.Sender != after {
		t.Errorf("Should trim all the quotes from the sender information in the email string")
	}
}

func TestExtractEmailAddress(t *testing.T) {
	emailAddress := "john@doe.org"
	test := fmt.Sprintf("John Doe <%s>", emailAddress)

	result, _ := Extract(test)

	if result.Address != emailAddress {
		t.Errorf("Should extract sender email address from email string")
	}
}

func TestExtractShouldDisplayFullEmailAddress(t *testing.T) {
	address := "John Doe <john@doe.org>"
	test := fmt.Sprintf(address)

	result, _ := Extract(test)

	if result.Text() != address {
		t.Errorf("Should display full email address from email string")
	}
}

func TestExtractShouldDisplayFullEmailAddressWhenMissingSenderInformation(t *testing.T) {
	address := "<john@doe.org>"
	test := fmt.Sprintf(address)

	result, _ := Extract(test)

	if result.Text() != address {
		t.Errorf("Should display full email address from email string")
	}
}

func TestExtractShouldDisplayFullEmailAddressWhenGibberishValidStringGiven(t *testing.T) {
	address := "<CALnn9HgYRR6CN=v9+CSeH8hwZaJZ2pF+cbQR5=YnzdcEJTrpig@mail.gmail.com>"
	test := fmt.Sprintf(address)

	result, _ := Extract(test)

	if result.Text() != address {
		t.Errorf("Should display full email address from gibberish, but valid email string")
	}
}

func TestExtractShouldDisplayFullEmailAddressWhenEmailStringWithoutAngleBracketsGiven(t *testing.T) {
	address := "random@gmail.com"

	result, _ := Extract(address)

	if result.Address != address {
		t.Errorf("Should display full email address from valid email string without angle brackets")
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
