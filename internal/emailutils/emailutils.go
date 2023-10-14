package emailutils

import (
	"errors"
	"fmt"
	"regexp"
)

const EmailRegex = `[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]+`

func ValidateEmailValid(s string) error {
	var emailRegex = regexp.MustCompile(fmt.Sprintf("^%s$", EmailRegex))
	if !emailRegex.MatchString(s) {
		return errors.New("Email must be of a valid format")
	}
	return nil
}
