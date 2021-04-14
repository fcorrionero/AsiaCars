package http

import (
	"fmt"
)

type Error struct {
	Err string `json:"error"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("error %v", e.Err)
}

func NewHttpError(msg string) *Error {
	err := new(Error)
	err.Err = msg
	return err
}
