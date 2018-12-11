package log

import (
	"context"
	"time"
	"log/conf"
	"log/plugins/zaplog"
	//"log/plugins/logrus"
	"testing"
)

func TestSetLogger(t *testing.T) {
	//设置为当前目录下 //设置级别
	SetLogger(zaplog.New(
		conf.WithLogName(time.Now().Format("2006-01-02")),
		conf.WithLogPath("logs"),
		conf.WithLogLevel("info"),
	))
	for i := 0; i < 10000; i++ {
		Debug("hello", context.Background())
		Infof("hello %s", "world", context.Background())
		Infof("hello %s", "world")
		Infof("hello %s,%d", "world", 2018, context.Background())
		Errorf("hello %s,%d", "world", 2018)
	}

/*	l2 := logrus.New(conf.WithLogPath("tmp"), conf.WithLogName("logrus"), conf.WithProjectName("logrus test"))
	SetLogger(l2)
	for i := 0; i < 1000000; i++ {
		Debug("hello", context.Background())
		Infof("hello %s", "world", context.Background())
		Infof("hello %s", "world")
		Infof("hello %s,%d", "world", 2018, context.Background())
		Errorf("hello %s,%d", "world", 2018)
	}*/

	time.Sleep(time.Second * 5)
}
