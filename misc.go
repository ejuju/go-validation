package validation

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
)

func IsNotNil(in any) CheckFunc {
	return func() error {
		if in == nil {
			return errors.New("should not be nil")
		}
		return nil
	}
}

func IsEmailAddress(in string) CheckFunc {
	return func() error {
		_, err := mail.ParseAddress(in)
		if err != nil {
			return fmt.Errorf("invalid email address: %s", in)
		}
		return nil
	}
}

func IsEmailAddresses(in []string) CheckFunc {
	return func() error {
		// Ensure string slice is not empty
		err := IsNotEmptyStringSlice(in)()
		if err != nil {
			return err
		}
		for i, emailaddr := range in {
			err = IsEmailAddress(emailaddr)()
			if err != nil {
				return fmt.Errorf("at index %d: %w", i, err)
			}
		}
		return nil
	}
}

func IsPhoneNumber(in string) CheckFunc {
	rgx := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	return func() error {
		if !rgx.MatchString(in) {
			return fmt.Errorf("input is not a phone number")
		}
		return nil
	}
}
