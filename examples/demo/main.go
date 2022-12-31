package main

import (
	"fmt"

	"github.com/ejuju/go-validation"
)

func main() {
	emailAddr := "hehe@.fr"
	age := -2

	err := validation.CheckAll(
		validation.Check(
			fmt.Sprintf("invalid email address: %s", emailAddr),
			validation.IsEmailAddressString(emailAddr),
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
