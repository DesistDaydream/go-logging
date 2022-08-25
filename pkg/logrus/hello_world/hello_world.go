package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// Logrus 共有七个日志级别，由高到底分别为：Trace、Debug、Info、Warning、Error、Fatal、Panic
	// 默认情况下，只有 Info 及以下级别可以正常输出。如果想要输出高级别日志，通过 SetLevel() 函数设置日志级别即可
	// SetLevel() 函数的实参可以通过 ParseLevel() 函数将字符串解析为对应级别
	// logrus.SetLevel(logrus.InfoLevel)

	// 输出 Info 级别的日志内容
	logrus.Info("Hello World")
}

// 输出内容如下：
// time="2021-09-20T11:58:36+08:00" level=info msg="Hello World"
