package emailutils

import (
	"fmt"
	"regexp"

	"github.com/heyzec/govtech-assignment/internal/errors"
)

const EmailRegex = `[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]+`

func ValidateEmailValid(s string) error {
	var emailRegex = regexp.MustCompile(fmt.Sprintf("^%s$", EmailRegex))
	if !emailRegex.MatchString(s) {
		return errors.NewInvalidEmailError(s)
	}
	return nil
}

func ValidateEmailsValid(arr []string) error {
	for _, s := range arr {
		if err := ValidateEmailValid(s); err != nil {
			return err
		}
	}
	return nil
}
