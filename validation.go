package validation

import (
	"errors"
	"net/mail"
)

func Check(errmsg string, checks ...bool) error {
	for _, c := range checks {
		if !c {
			return errors.New(errmsg)
		}
	}
	return nil
}

func CheckAll(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func IsEmailAddr(in string) bool {
	_, err := mail.ParseAddress(in)
	return err == nil
}
