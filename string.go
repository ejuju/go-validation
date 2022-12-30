package validation

import (
	"fmt"
	"unicode/utf8"
)

func IsNotEmpty(in string) CheckFunc {
	return func() error {
		if in == "" {
			return fmt.Errorf("string is empty")
		}
		return nil
	}
}

func IsUTF8(in string) CheckFunc {
	return func() error {
		if !utf8.ValidString(in) {
			return fmt.Errorf("string is not UTF-8 compliant")
		}
		return nil
	}
}

func IsLongerThanString(in string, min int) CheckFunc {
	return func() error {
		gotlen := utf8.RuneCountInString(in)
		if gotlen < min {
			return fmt.Errorf(
				"string misses %d characters to have minimum length %d", min-gotlen, min,
			)
		}
		return nil
	}
}

// Maximum is NOT included.
func IsStrictlyShorterThanString(in string, max int) CheckFunc {
	return func() error {
		gotlen := utf8.RuneCountInString(in)
		if gotlen >= max {
			return fmt.Errorf(
				"string is %d characters longer than maximum length %d", gotlen-max+1, max,
			)
		}
		return nil
	}
}

func IsWithinMinMaxString(in string, min, max int) CheckFunc {
	return func() error {
		err := CheckAll(
			IsNotEmpty(in)(),
			IsUTF8(in)(),
			IsLongerThanString(in, min)(),
			IsStrictlyShorterThanString(in, min)(),
		)
		if err != nil {
			return err
		}
		return nil
	}
}

func IsNotEmptyStringSlice(in []string) CheckFunc {
	return func() error {
		if len(in) == 0 {
			return fmt.Errorf("string slice is empty")
		}
		return nil
	}
}

func IsNotStrictlyLongerThanStringSlice(maxlen int, in []string) CheckFunc {
	return func() error {
		if len(in) >= maxlen {
			return fmt.Errorf("string is strictly longer than maximum length %d", maxlen)
		}
		return nil
	}
}
