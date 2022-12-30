package validation

import "fmt"

type CheckFunc func() error

func Check(id string, checks ...CheckFunc) error {
	for _, checkfn := range checks {
		err := checkfn()
		if err != nil {
			return fmt.Errorf("%q validation failed: %w", id, err)
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
