package dingtalk

import "github.com/CodyGuo/dingtalk/pkg/robot"

// RobotSendText text类型的消息
func (dt *DingTalk) RobotSendText(text string, options ...robot.SendOption) error {
	msg := robot.Text{Content: text}
	return dt.Request(robot.NewSend(msg, options...))
}

// RobotSendLink link类型的消息
func (dt *DingTalk) RobotSendLink(title, text, messageURL, picURL string, options ...robot.SendOption) error {
	msg := robot.Link{
		Title:      title,
		Text:       text,
		MessageURL: messageURL,
		PicURL:     picURL,
	}
	return dt.Request(robot.NewSend(msg))
}

// RobotSendMarkdown markdown类型的消息
func (dt *DingTalk) RobotSendMarkdown(title, text string, options ...robot.SendOption) error {
	msg := robot.Markdown{
		Title: title,
		Text:  text,
	}
	return dt.Request(robot.NewSend(msg, options...))
}

// RobotSendEntiretyActionCard 整体跳转ActionCard类型
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

// RobotSendIndependentActionCard 独立跳转ActionCard类型
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

// RobotSendFeedCard FeedCard类型
func (dt *DingTalk) RobotSendFeedCard(links []robot.FeedCardLink, options ...robot.SendOption) error {
	msg := robot.FeedCard{
		Links: links,
	}
	return dt.Request(robot.NewSend(msg, options...))
}
