package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/ejuju/go-validation"
)

func main() {
	emailAddr := "hehe@.fr"
	maxEmailAddrLength := 300
	age := -2

	err := validation.CheckAll(
		validation.Check("email address is empty", emailAddr == ""),
		validation.Check(
			fmt.Sprintf(
				"email address too long: %d chars (max length is %d)",
				utf8.RuneCountInString(emailAddr),
				maxEmailAddrLength,
			),
			utf8.RuneCountInString(emailAddr) <= maxEmailAddrLength,
		),
		validation.Check(
			fmt.Sprintf("invalid email address: %s", emailAddr),
			validation.IsEmailAddr(emailAddr),
		),
		validation.Check(
			fmt.Sprintf("age %d is not within 0 and 150", age),
			age >= 0,
			age <= 150,
		),
	)
	if err != nil {
		panic(err)
	}
}
