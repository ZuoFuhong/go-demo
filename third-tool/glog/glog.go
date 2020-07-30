// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/golang/glog"
)

// 使用示例：go run glog.go -log_dir="./" -v=3 -alsologtostderr=true -stderrthreshold=-1 -logtostderr=true
//  -log_dir：指定log files的目录，默认是os.TempDir()
//  -v：是vLog是用户自定义的log级别
//  -alsologtostderr：同时输出到os.Stderr和log files
//  -stderrthreshold：大于等于该severity的log，会输出到os.Stderr
//  -logtostderr：只输出到os.Stderr，
func main() {
	flag.Parse()
	defer glog.Flush()

	glog.V(3).Info("hello")
}
