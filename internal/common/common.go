package common

type ClientError struct {
	Code    int
	Message string
}

type ServerError struct {
	Code    int
	Message string
	Detail  string
}
