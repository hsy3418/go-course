package hsy3418-puppy

import "fmt"

type Error struct {
	Message string
	Code    int
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code:%d,Error:%s", e.Code, e.Message)
}
