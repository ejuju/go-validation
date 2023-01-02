package main

import (
	"github.com/ejuju/go-validation"
)

func main() {
	emailAddr := "hehe@.fr"

	err := validation.CheckAll(
		validation.CheckEmailAddress(emailAddr),                               // ensure valid email
		validation.CheckStringLength("phone number", "06 12 34 45 65", 0, 20), // optional or shorter than 20
		validation.CheckStringLength("message", "bla bla bla", 1, 20_000),     // required and between 1 and 20_000 chars
	)
	if err != nil {
		panic(err)
	}
}
