package postgres

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	codeInternal = iota - 1
	codeSuccess
	codeFailure
)

const (
	empty   int = iota + 1
	unknown     // TODO: придумать другую ошибку
	notFound
	alreadyExists
	invalidArguments
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

	switch e.Code {
	case codeSuccess:
		return nil
	case codeInternal:
		return ErrInternal
	case codeFailure:
		var errorStr string

		for k, v := range e.Details {
			var s string

			switch v {
			case empty:
				s = `empty`
			case unknown:
				s = `unknown error`
			case notFound:
				s = `not found`
			case alreadyExists:
				s = `already exists`
			case invalidArguments:
				s = `invalid argument`
			}

			errorStr += fmt.Sprintf("%s: %s", k, s)
		}

		return errors.New(errorStr)
	default:
		return fmt.Errorf("unknown error code. error: %v", string(queryError))
	}
}
