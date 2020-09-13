package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

/*
	一、Go性能优化
		Go语言项目中的性能优化主要有以下几个方面：
		- CPU profile：报告程序的 CPU 使用情况，按照一定频率去采集应用程序在 CPU 和寄存器上面的数据
		- Memory Profile（Heap Profile）：报告程序的内存使用情况
		- Block Profiling：报告 goroutines 不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈
		- Goroutine Profiling：报告 goroutines 的使用情况，有哪些 goroutine，它们的调用关系是怎样的
	二、采集性能数据
		pprof是golang官方提供的性能评测工具，包括以下两个标准库：
		- runtime/pprof：采集工具型应用运行数据进行分析
		- net/http/pprof：采集服务型应用运行时数据进行分析

		pprof开启后，每隔一段时间（10ms）就会收集下当前的堆栈信息，获取格格函数占用的CPU以及内存资源；最后通过对这些采样数据进行分析，形成一个性能分析报告。
		注意，我们只应该在性能测试的时候才在代码中引入pprof。
*/

/*
	三、工具型应用
		如果你的应用程序是运行一段时间就结束退出类型。那么最好的办法是在应用退出的时候把 profiling 的报告保存到文件中，进行分析。
		对于这种情况，可以使用runtime/pprof库。 首先在代码中导入runtime/pprof工具：

		运行后，得到采样数据，使用 go tool pprof 工具进行分析：
		$ go tool pprof cpu.pprof
		$ go tool pprof memory.pprof
*/
func CPUProfile() {
	file, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	// 开启CPU性能分析
	err = pprof.StartCPUProfile(file)
	if err != nil {
		panic(err)
	}
	// 停止CPU性能分析
	defer pprof.StopCPUProfile()

	var data []byte
	for n := 0; n < 10000000; n++ {
		data = append(data, 12)
	}
}

func MemoryProfile() {
	file, err := os.Create("memory.pprof")
	if err != nil {
		panic(err)
	}

	var data []byte
	for n := 0; n < 10000000; n++ {
		data = append(data, 12)
	}
	// 内存性能分析
	err = pprof.WriteHeapProfile(file)
	if err != nil {
		panic(err)
	}
}

/*
	四、服务型应用
		如果你的应用程序是一直运行的，比如 web 应用，那么可以使用net/http/pprof库，它能够在提供 HTTP 服务进行分析。
		1.首先导入包：_ "net/http/pprof"，注意利用下划线"_"导入，只需要运行该包的init()函数，该包会自动完成信息采集并保存在内存中。
		2.通过Web界面：http://127.0.0.1:8080/debug/pprof/查看服务的运行情况。
		3.通过交互式终端使用
			$ go tool pprof http://127.0.0.1:8080/debug/pprof/profile?seconds=10
			$ go tool pprof http://127.0.0.1:8080/debug/pprof/heap
		4.相关的数据指标：
			- cpu：$host/debug/pprof/profile，默认进行30s的CPU	Profiling，得到一个分析用的profile文件。
			- block：$host/debug/pprof/block，查看导致阻塞同步的堆栈跟踪。
			- goroutine：$host/debug/pprof/goroutine，查看当前所有运行的goroutines堆栈跟踪。
			- heap：$host/debug/pprof/heap，查看活动对象的内存分配情况
			- mutex：$host/debug/pprof/mutex，查看导致互斥锁的竞争持有者的堆栈跟踪。
*/
func WebServer() {
	go func() {
		var data []byte
		for {
			data = append(data, 12)
		}
	}()
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}

/*
	五、go tool pprof命令
		不管是工具型应用还是服务型应用，我们使用相应的pprof库获取数据之后，下一步的都要对这些数据进行分析，我们可以使用go tool pprof命令行工具。
		$ go tool pprof [binary] [source]
		其中：
		- binary 是应用的二进制文件，用来解析各种符号；
		- source 表示 profile 数据的来源，可以是本地的文件，也可以是 http 地址。

		注意事项：获取的 Profiling 数据是动态的，要想获得有效的数据，请保证应用处于较大的负载（比如正在生成中运行的服务，或者通过其他工具模拟访问压力）。
		否则如果应用处于空闲状态，得到的结果可能没有任何意义。

	六、图形化
		通过svg图的方式查看程序中详细的CPU占用情况，需要安装graphviz图形化工具。
		$ brew install graphviz
		安装完成后在 go tool pprof 的命令行中输入web即可生成一个svg格式的文件，将其用浏览器打开即可。

		关于图形的说明： 每个框代表一个函数，理论上框的越大表示占用的CPU资源越多。 方框之间的线条代表函数之间的调用关系。 线条上的数字表示函数调用的次数。
		方框中的第一行数字表示当前函数占用CPU的百分比，第二行数字表示当前函数累计占用CPU的百分比。
	七、火焰图
		$ go-torch -u http://127.0.0.1:8080  --seconds 60 -f cpu.svg
		$ go-torch  http://127.0.0.1:8080/debug/pprof/heap --colors mem  -f mem.svg
	八、参考资料
		Go pprof性能调优：https://www.cnblogs.com/nickchen121/p/11517452.html
		Go代码调优利器-火焰图：https://lihaoquan.me/2017/1/1/Profiling-and-Optimizing-Go-using-go-torch.html
*/
func main() {
	WebServer()
}
