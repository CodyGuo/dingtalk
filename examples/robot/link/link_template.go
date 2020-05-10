package main

import (
	"github.com/CodyGuo/dingtalk"
	"github.com/CodyGuo/glog"
)

const linkTemplate = `{{.}}
https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI
https://cdn.pixabay.com/photo/2020/05/05/08/05/butterfly-5131967_960_720.jpg
这个即将发布的新版本，创始人xx称它为“红树林”。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？
`

func main() {
	glog.SetFlags(glog.LglogFlags)
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=xx"
	secret := "xxx"
	dt := dingtalk.New(webHook, dingtalk.WithSecret(secret))

	// go 模板
	err := dt.RobotSendLinkWithTemplate(linkTemplate, "通知")
	if err != nil {
		glog.Fatal(err)
	}

	// go 模板文件
	err = dt.RobotSendLinkWithFile("link.tmpl", "通知")
	if err != nil {
		glog.Fatal(err)
	}
	glog.Info("send message success")
}
