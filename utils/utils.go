package utils

import "fmt"

func AppendNewErr(initErr, newErr error) error {
	if initErr == nil {
		return newErr
	}
	return fmt.Errorf("%v, %w", initErr, newErr)
}
