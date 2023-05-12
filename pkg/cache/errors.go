package cache

import (
	"fmt"
)

type InvalidTypeError struct {
	Type string
}

func (c InvalidTypeError) Error() string {
	return fmt.Sprintf("%s is not supported as cache type.", c.Type)
}
