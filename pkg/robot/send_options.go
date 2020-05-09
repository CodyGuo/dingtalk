package robot

type SendOption func(*Send)

func SendWithAtMobiles(atMobiles []string) SendOption {
	return func(s *Send) {
		s.at = &At{
			AtMobiles: atMobiles,
		}
	}
}

func SendWithIsAtAll(isAtAll bool) SendOption {
	return func(s *Send) {
		s.at = &At{
			IsAtAll: isAtAll,
		}
	}
}
