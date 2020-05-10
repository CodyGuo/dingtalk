package dingtalk

import "encoding/json"

type ResponseMsg struct {
	ErrCode         int64  `json:"errcode"`
	ErrMsg          string `json:"errmsg"`
	ApplicationHost string `json:"application_host,omitempty"`
	ServiceHost     string `json:"service_host,omitempty"`
}

func (r ResponseMsg) String() string {
	data, err := json.Marshal(&r)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
