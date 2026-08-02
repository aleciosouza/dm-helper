package handler

import "fmt"

func errParamIsRequired(name string, paramType string) error {
	return fmt.Errorf("param: %s (%s) is required", name, paramType)
}
