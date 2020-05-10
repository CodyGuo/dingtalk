package dingtalk

import (
	"bufio"
	"strings"

	"github.com/CodyGuo/dingtalk/utils"

	"github.com/CodyGuo/dingtalk/pkg/robot"
)

// RobotSendText text类型的消息
func (dt *DingTalk) RobotSendText(text string, options ...robot.SendOption) error {
	msg := robot.Text{Content: text}
	return dt.Request(robot.NewSend(msg, options...))
}

// RobotSendTextWithTemplate text类型的消息
// template
func (dt *DingTalk) RobotSendTextWithTemplate(text string, data interface{}, options ...robot.SendOption) error {
	out, err := utils.TemplateParse("text", text, data)
	if err != nil {
		return err
	}
	return dt.RobotSendText(string(out), options...)
}

// RobotSendTextWithFile text类型的消息
// template file
func (dt *DingTalk) RobotSendTextWithFile(filename string, data interface{}, options ...robot.SendOption) error {
	out, err := utils.TemplateParseFile(filename, data)
	if err != nil {
		return err
	}
	return dt.RobotSendText(string(out), options...)
}

// RobotSendLink link类型的消息
func (dt *DingTalk) RobotSendLink(title, text, messageURL, picURL string, options ...robot.SendOption) error {
	msg := robot.Link{
		Title:      title,
		Text:       text,
		MessageURL: messageURL,
		PicURL:     picURL,
	}
	return dt.Request(robot.NewSend(msg, options...))
}

// RobotSendLinkWithTemplate link类型的消息
// template:
//     第一行: title
//     第二行: messageURL
//     第三行: picURL
//     其他行: text
func (dt *DingTalk) RobotSendLinkWithTemplate(text string, data interface{}, options ...robot.SendOption) error {
	out, err := utils.TemplateParse("link", text, data)
	if err != nil {
		return err
	}
	msg := parseLinkWithTemplate(string(out), data)
	return dt.Request(robot.NewSend(msg, options...))
}

// RobotSendLinkWithFile link类型的消息
// template file:
//     第一行: title
//     第二行: messageURL
//     第三行: picURL
//     其他行: text
func (dt *DingTalk) RobotSendLinkWithFile(filename string, data interface{}, options ...robot.SendOption) error {
	out, err := utils.TemplateParseFile(filename, data)
	if err != nil {
		return err
	}
	msg := parseLinkWithTemplate(string(out), data)
	return dt.Request(robot.NewSend(msg, options...))
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

// parseLinkWithTemplate 解析template
// template:
//     第一行: title
//     第二行: messageURL
//     第三行: picURL
//     其他行: text
func parseLinkWithTemplate(text string, data interface{}) *robot.Link {
	n := 0
	b := []byte{}
	msg := &robot.Link{}
	buf := bufio.NewScanner(strings.NewReader(text))
	for buf.Scan() {
		line := buf.Text()
		switch n {
		case 0:
			msg.Title = line
		case 1:
			msg.MessageURL = line
		case 2:
			msg.PicURL = line
		default:
			b = append(b, line...)
			b = append(b, '\n')
		}
		n++
	}
	msg.Text = string(b)
	return msg
}
