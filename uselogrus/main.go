package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var lf *os.File

func main() {
	lf, err := os.OpenFile("./log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return
	}
	defer lf.Close()
	//logrus.SetOutput(lf)

	logrus.Debug("Useful debugging information.")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).WithTime(time.Now()).Info("A walrus appears")
}

func init() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	logrus.SetOutput(os.Stdout)
	//设置最低loglevel
	logrus.SetLevel(logrus.DebugLevel)

}
