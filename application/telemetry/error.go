package telemetry

import (
	"errors"
	"fmt"
)

type Error struct {
	Err error
}

func (e *Error) Error() string {
	return fmt.Sprintf("err %v", e.Err)
}

func NewAppError(msg string) *Error {
	err := new(Error)
	err.Err = errors.New(msg)
	return err
}
