package postgres

type QueryError struct {
	Code    int32    `json:"code"`
	Details *details `json:"details"`
}

type details struct {
	Msg *string `json:"msg"`
}
