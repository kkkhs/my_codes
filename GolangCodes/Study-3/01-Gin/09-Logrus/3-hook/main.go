package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type MyHook struct {
}

func (hook MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook MyHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = "khs"
	fmt.Println(entry.Level)
	return nil
}

func main() {
	logrus.AddHook(&MyHook{})

	logrus.Warnln("你好")
	logrus.Errorf("你好")
}
