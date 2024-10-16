package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file, _ := os.OpenFile("./info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logrus.SetReportCaller(true)
	// writes := []io.Writer{
	// 	file,
	// 	os.Stdout,
	// }
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))

	logrus.Infof("你好")
	logrus.Infoln("你好")
	logrus.Errorf("出错了")
	logrus.Errorln("出错了")
}
