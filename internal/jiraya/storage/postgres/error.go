package postgres

type QueryError struct {
	Error *string `json:"error"`
}
