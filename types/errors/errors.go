package errors

import "errors"

type cErr string

const (
	FileOnlyAccept cErr = "file is not accept types"
)

func Err(e cErr) error {
	return errors.New(string(e))
}
