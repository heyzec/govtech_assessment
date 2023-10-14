package emailutils_test

import (
	"testing"

	"github.com/heyzec/govtech-assignment/internal/helpers/emailutils"
	"github.com/stretchr/testify/require"
)

func TestValidEmails(t *testing.T) {
	validEmails := []string{
		"guy@gmail.com",
		"simple@example.com",
		"abc@example.co.uk",
		"a@b.com",
		"other.email-with-hyphen@example.com",
		"more-hyphens@strange-hyphens.com",
		"0123457@numbers.com",
		"plus+plus@example.com",
		"_______@underscores.com",
	}
	for _, email := range validEmails {
		require.Nil(t, emailutils.ValidateEmailValid(email))
	}
}

func TestInvalidEmails(t *testing.T) {
	invalidEmails := []string{
		"guy@com",
		"guygmailcom",
		"@guy@gmail.com",
		"@guy@gmail.com",
		"#@%^%#$@#$@#.com",
		"email@example.com (Joe Smith)",
	}
	for _, email := range invalidEmails {
		require.NotNil(t, emailutils.ValidateEmailValid(email))
	}
}
