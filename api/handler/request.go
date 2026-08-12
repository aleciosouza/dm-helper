package handler

import "fmt"

func ErrParamIsRequired(name string, paramType string) error {
	return fmt.Errorf("param: %s (%s) is required", name, paramType)
}
