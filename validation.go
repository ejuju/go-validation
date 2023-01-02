package validation

import (
	"fmt"
	"net/mail"
	"unicode/utf8"
)

func Check(errmsg string, checks ...bool) error {
	for i, c := range checks {
		if !c {
			return fmt.Errorf("at check %d: %s", i, errmsg)
		}
	}
	return nil
}

func CheckAll(errs ...error) error {
	for i, err := range errs {
		if err != nil {
			return fmt.Errorf("at index %d: %w", i, err)
		}
	}
	return nil
}

func IsEmailAddr(in string) bool {
	_, err := mail.ParseAddress(in)
	return err == nil
}

func CheckEmailAddress(in string) error {
	return CheckAll(
		CheckStringLength("email", in, 1, 300),
		Check(
			fmt.Sprintf("invalid email address: %s", in),
			IsEmailAddr(in),
		),
	)
}

func CheckStringLength(name, in string, min, max int) error {
	return Check(
		fmt.Sprintf(
			"%s is not within length [%d;%d]: %d chars",
			name,
			min,
			max,
			utf8.RuneCountInString(in),
		),
		utf8.RuneCountInString(in) >= min,
		utf8.RuneCountInString(in) <= max,
	)
}
