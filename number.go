package validation

import "fmt"

func IntIsAboveMin(in, min int, intName string) CheckFunc {
	if intName == "" {
		intName = "integer"
	}
	return func() error {
		if in < min {
			return fmt.Errorf("%s (%d) is below minimum (%d)", intName, min-in, min)
		}
		return nil
	}
}

func IntIsBelowMax(in, max int, intName string) CheckFunc {
	if intName == "" {
		intName = "integer"
	}
	return func() error {
		if in >= max {
			return fmt.Errorf("%s (%d) is above maximum (%d)", intName, max, in-max)
		}
		return nil
	}
}
