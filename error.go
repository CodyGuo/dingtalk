package dingtalk

import "encoding/json"

type Err struct {
	Code            int64  `json:"errcode"`
	Msg             string `json:"errmsg"`
	ApplicationHost string `json:"application_host,omitempty"`
	ServiceHost     string `json:"service_host,omitempty"`
	Detail          error  `json:"detail,omitempty"`
}

func newErr(msg string, err error) *Err {
	return &Err{
		Code:   -1,
		Msg:    msg,
		Detail: err,
	}
}

func (e Err) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
