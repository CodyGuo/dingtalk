package dingtalk

import "time"

type Option func(*DingTalk)

func WithSecret(secret string) Option {
	return func(dt *DingTalk) {
		dt.secret = secret
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(dt *DingTalk) {
		dt.timeout = timeout
	}
}
