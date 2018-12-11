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
		Debug("瑞波问题：1.fee！！！！！TBD2.queue状态的交易  ----！！！！！！！TBD  返回queue，塞入数据表“确认中”，queue中的交易状态还会成功吗？成功后会包含在下个ledger里吗？爬块时，表里不存在insert，表里存在update---TBD2.出块时间小组本周工作：1.瑞波和前端联调接口2.波场瑞波线上环境节点配置确认3.优化波场跟块服务的追块功能4.stellar预研下周计划：1.完成瑞波调试，提测，修bug2.研究跨链3.带柯杰熟悉代码，捋清接入stellar逻辑")
		Debug("hello", context.Background())
		Infof("hello %s", "world", context.Background())
		Infof("hello %s", "world")
		Infof("hello %s,%d", "world", 2018, context.Background())
		Errorf("hello %s,%d", "world", 2018)
	}

/*	l2 := logrus.New(conf.WithLogPath("tmp"), conf.WithLogName("logrus"), conf.WithProjectName("logrus test"))
	SetLogger(l2)
	for i := 0; i < 1000000; i++ {
		Debug("瑞波问题：1.fee！！！！！TBD2.queue状态的交易  ----！！！！！！！TBD  返回queue，塞入数据表“确认中”，queue中的交易状态还会成功吗？成功后会包含在下个ledger里吗？爬块时，表里不存在insert，表里存在update---TBD2.出块时间小组本周工作：1.瑞波和前端联调接口2.波场瑞波线上环境节点配置确认3.优化波场跟块服务的追块功能4.stellar预研下周计划：1.完成瑞波调试，提测，修bug2.研究跨链3.带柯杰熟悉代码，捋清接入stellar逻辑")
		Debug("hello", context.Background())
		Infof("hello %s", "world", context.Background())
		Infof("hello %s", "world")
		Infof("hello %s,%d", "world", 2018, context.Background())
		Errorf("hello %s,%d", "world", 2018)
	}*/

	time.Sleep(time.Second * 5)
}
