package main

import "github.com/ejuju/go-validation"

func main() {
	emailAddr := "hehe@.fr"
	phoneNumber := "06 112 ss"
	age := -2

	err := validation.CheckAll(
		validation.Check("user data is valid",
			validation.IsEmailAddress(emailAddr),
			validation.IsPhoneNumber(phoneNumber),
			validation.IntIsAboveMin(age, 0, "age"),
		),
		validation.Check("some other checks",
			validation.IsLongerThanString("i'm short", 10),
			validation.IntIsBelowMax(0, 10, "some number"),
		),
	)
	if err != nil {
		panic(err)
	}
}
