package dingtalk

type Option func(*DingTalk)

func WithSecret(secret string) Option {
	return func(dt *DingTalk) {
		dt.secret = secret
	}
}
