package dingtalk

import "dingtalk/pkg/robot"

func (dt *DingTalk) RobotSendText(text string, options ...robot.SendOption) error {
	msg := robot.Text{Content: text}
	return dt.Request(robot.NewSend(msg, options...))
}

func (dt *DingTalk) RobotSendLink(title, text, messageURL, picURL string, options ...robot.SendOption) error {
	msg := robot.Link{
		Title:      title,
		Text:       text,
		MessageURL: messageURL,
		PicURL:     picURL,
	}
	return dt.Request(robot.NewSend(msg))
}

func (dt *DingTalk) RobotSendMarkdown(title, text string, options ...robot.SendOption) error {
	msg := robot.Markdown{
		Title: title,
		Text:  text,
	}
	return dt.Request(robot.NewSend(msg, options...))
}

func (dt *DingTalk) RobotSendEntiretyActionCard(title, text, singleTitle, singleURL, btnOrientation string, options ...robot.SendOption) error {
	msg := robot.ActionCard{
		Title:          title,
		Text:           text,
		SingleTitle:    singleTitle,
		SingleURL:      singleURL,
		BtnOrientation: btnOrientation,
	}
	return dt.Request(robot.NewSend(msg, options...))
}

func (dt *DingTalk) RobotSendIndependentActionCard(title, text, btnOrientation string, btns map[string]string, options ...robot.SendOption) error {
	var rBtns []robot.Btn
	for title, actionURL := range btns {
		btn := robot.Btn{
			Title:     title,
			ActionURL: actionURL,
		}
		rBtns = append(rBtns, btn)
	}
	msg := robot.ActionCard{
		Title:          title,
		Text:           text,
		Btns:           rBtns,
		BtnOrientation: btnOrientation,
	}
	return dt.Request(robot.NewSend(msg, options...))
}

func (dt *DingTalk) RobotSendFeedCard(links []robot.FeedCardLink, options ...robot.SendOption) error {
	msg := robot.FeedCard{
		Links: links,
	}
	return dt.Request(robot.NewSend(msg, options...))
}
