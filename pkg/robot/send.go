package robot

import (
	"encoding/json"
)

type Send struct {
	at  *At
	msg IMsg
}

func NewSend(msg IMsg, options ...SendOption) *Send {
	r := &Send{msg: msg}
	for _, option := range options {
		option(r)
	}
	return r
}

func (r *Send) GetMsgType() string {
	return r.msg.GetType()
}

func (r *Send) GetAt() *At {
	return r.at
}

func (r *Send) SetAt(at *At) {
	r.at = at
}

func (r *Send) GetMethod() string {
	return "POST"
}

func (r *Send) GetHeader() map[string]string {
	header := map[string]string{
		"Content-type":  "application/json;charset=UTF-8",
		"Cache-Control": "no-cache",
		"Connection":    "Keep-Alive",
		"User-Agent":    "ding talk robot send",
	}
	return header
}

func (r *Send) GetBody() ([]byte, error) {
	msg := make(map[string]interface{}, 3)
	msg["msgtype"] = r.msg.GetType()
	if r.at != nil {
		msg["at"] = r.at
	}
	name := r.msg.GetType()
	msg[name] = r.msg
	return json.Marshal(msg)
}

func (r *Send) GetSuccessCode() int64 {
	return 0
}

func (r *Send) GetApiName() string {
	return "dingtalk.oapi.robot.send"
}

// At @
type At struct {
	// atMobiles 被@人的手机号(在text内容里要有@手机号)
	AtMobiles []string `json:"atMobiles,omitempty"`

	// isAtAll   @所有人时：true，否则为：false
	IsAtAll bool `json:"isAtAll,omitempty"`
}

func (a *At) GetAtMobiles() []string {
	return a.AtMobiles
}

func (a *At) SetAtMobiles(atMobiles []string) {
	a.AtMobiles = atMobiles
}

func (a *At) GetIsAtAll() bool {
	return a.IsAtAll
}

func (a *At) SetIsAtAll(isAtAll bool) {
	a.IsAtAll = isAtAll
}
