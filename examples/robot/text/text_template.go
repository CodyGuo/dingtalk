package main

import (
	"time"

	"github.com/CodyGuo/dingtalk"
	"github.com/CodyGuo/glog"
)

type Msg struct {
	Title    string
	Hostname string
	Address  string
	Event    string
	Time     string
}

func main() {
	glog.SetFlags(glog.LglogFlags)
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=xxx"
	secret := "xxx"
	dt := dingtalk.New(webHook, dingtalk.WithSecret(secret))

	// go 模板
	err := dt.RobotSendTextWithTemplate("{{.}}\ntemplate text message", "通知")
	if err != nil {
		glog.Fatal(err)
	}

	// go 模板文件
	msg := &Msg{
		Title:    "通知",
		Hostname: "centos",
		Address:  "192.168.1.1",
		Event:    "CPU超过80%, 当前CPU: 82.5%",
		Time:     time.Now().Format("2006-01-02 15:04:05.000"),
	}
	err = dt.RobotSendTextWithFile("text.tmpl", msg)
	if err != nil {
		glog.Fatal(err)
	}

	glog.Info("send message success")
}
