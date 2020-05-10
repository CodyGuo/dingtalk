package dingtalk

type Error struct {
	Op   string `json:"Op"`
	URL  string `json:"URL"`
	Body string `json:"Body"`
	Err  error  `json:"Err"`
}

func newError(op, url, body string, err error) *Error {
	return &Error{
		Op:   op,
		URL:  url,
		Body: body,
		Err:  err,
	}
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Error() string {
	return e.Op + " " + e.URL + " " + e.Body + ": " + e.Err.Error()
}
