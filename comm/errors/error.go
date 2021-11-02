package errors

import "encoding/json"

// Error defined TODO
type Error struct {
	Code   int32
	Status string
	Detail string
}

// Error defined TODO
func (e Error) Error() string {
	errByte, _ := json.Marshal(e)
	return string(errByte)
}

// New defined TODO
func New(code int32, desc string, err error) error {
	return &Error{Code: code, Detail: err.Error(), Status: desc}
}

// Format defined TODO
func Format(err error) string {
	msg, er := err.Error(), Error{}
	json.Unmarshal([]byte(msg), &er)
	if er.Code != 0 {
		return err.Error()
	}
	er.Code = 500
	er.Detail = err.Error()
	er.Status = "internal error"
	return er.Error()
}
