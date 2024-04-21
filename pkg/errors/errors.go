package errors

import "fmt"

func WrapError(mes string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", mes, err)
	}
	return nil
}
