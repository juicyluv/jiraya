package postgres

import (
	"encoding/json"
	"errors"
)

const (
	codeInternal = iota - 1
	codeSuccess
	codeFailure
)

var (
	ErrInternal = errors.New("internal error")
)

func handleQueryError(queryError []byte) error {
	if queryError == nil {
		return nil
	}

	e := QueryError{}

	err := json.Unmarshal(queryError, &e)

	if err != nil {
		return err
	}

	if e.Code == codeSuccess {
		return nil
	}

	if e.Details == nil || e.Details.Msg == nil {
		return ErrInternal
	}

	return errors.New(*e.Details.Msg)
}
