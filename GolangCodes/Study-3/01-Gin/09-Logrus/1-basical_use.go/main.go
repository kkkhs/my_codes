package main

import "github.com/sirupsen/logrus"

func PrintColor() {

}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetFormatter(&logrus.JSONFormatter{}) //更改字段显示方式
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,                  //设置颜色输出
		FullTimestamp:   true,                  //设置输出显示时间
		TimestampFormat: "2006-01-02 15:04:05", //设置时间格式
	})
	logrus.Error("出错")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("debug")
	logrus.Println("\033[31m打印\033[0m") //自定义颜色输出

	//设置特定字段
	log := logrus.WithField("app", "study").WithField("service", "logrus")

	log = log.WithFields(logrus.Fields{
		"user": "khs",
		"sex":  "man",
		"ip":   "127.0.0.1",
	})

	log.Errorf("你好")

}
