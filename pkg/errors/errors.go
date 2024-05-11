package errors

import "fmt"

func WrapError(err error, mes string) error {
	if err != nil {
		return fmt.Errorf("%s: %w", mes, err)
	}
	return nil
}

func WrapFail(err error, mes string) error {
	return WrapError(err, "couldn't "+mes)
}
