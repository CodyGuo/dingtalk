package dingtalk

import "encoding/json"

type Error struct {
	Op              string `json:"Op,omitempty"`
	URL             string `json:"URL,omitempty"`
	Code            int64  `json:"errcode,omitempty"`
	Msg             string `json:"errmsg,omitempty"`
	ApplicationHost string `json:"application_host,omitempty"`
	ServiceHost     string `json:"service_host,omitempty"`
	Err             error  `json:"Err,omitempty"`
}

func newError(op, url string, err error) *Error {
	return &Error{
		Op:  op,
		URL: url,
		Err: err,
	}
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
