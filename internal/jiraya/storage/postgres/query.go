package postgres

import (
	"encoding/json"
	"errors"
)

func handleQueryError(queryError []byte) error {
	e := QueryError{}

	err := json.Unmarshal(queryError, &e)

	if err != nil {
		return err
	}

	if e.Error == nil {
		return nil
	}

	return errors.New(*e.Error)
}
