package robot

const (
	M_TEXT msgType = iota
	M_LINK
	M_MARKDOWN
	M_ACTIONCARD
	M_FEEDCARD
)

type IMsg interface {
	GetType() string
}

type msgType uint8

func (r msgType) String() string {
	switch r {
	case M_TEXT:
		return "text"
	case M_LINK:
		return "link"
	case M_MARKDOWN:
		return "markdown"
	case M_ACTIONCARD:
		return "actionCard"
	case M_FEEDCARD:
		return "feedCard"
	default:
		return "Unknown"
	}
}
