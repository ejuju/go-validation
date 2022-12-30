package main

import "github.com/ejuju/go-validation"

func main() {
	emailAddr := "hehe@.fr"
	phoneNumber := "06 112 ss"
	age := -2

	err := validation.CheckAll(
		validation.Check("email is valid",
			validation.IsEmailAddress(emailAddr),
			validation.IsPhoneNumber(phoneNumber),
			validation.IntIsAboveMin(age, 0, "age"),
		),
	)
	if err != nil {
		panic(err)
	}
}
