package database

import "fmt"

type UnknownDbTypeError struct {
	Type string
}

func (udt UnknownDbTypeError) Error() string {
	return fmt.Sprintf("unknown data type, db-type: %s", udt.Type)
}
