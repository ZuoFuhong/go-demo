// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"

	"github.com/golang/glog"
)

/*
	一、介绍：
		golang/glog 是 C++ 版本 google/glog 的 Go 版本实现，基本实现了原生 glog 的日志格式。在 Kuberntes 中，glog 是默认日志库。

    二、日志级别
		INFO：普通日志；
		WARNING：告警日志；
		ERROR：错误日志；
		FATAL：严重错误日志，打印完日志后程序将会推出（os.Exit()）

	三、使用示例：
		$ go run glog.go -log_dir="./" -v=3 -alsologtostderr=true -stderrthreshold=-1 -logtostderr=true
	 	其中：
		-log_dir：指定log files的目录，默认是os.TempDir()
		-v：是vLog是用户自定义的log级别
		-alsologtostderr：同时输出到os.Stderr和log files
		-stderrthreshold：大于等于该severity的log，会输出到os.Stderr
		-logtostderr：只输出到os.Stderr

	四、参考资料
		Golang glog使用详解：https://www.cnblogs.com/sunsky303/p/11081165.html
*/
func main() {
	flag.Parse()
	defer glog.Flush()

	glog.V(3).Info("hello")
}
