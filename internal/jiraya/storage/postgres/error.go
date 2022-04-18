package postgres

type QueryError struct {
	Code    int32          `json:"code"`
	Details map[string]int `json:"details"`
}
