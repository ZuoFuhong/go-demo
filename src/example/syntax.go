// 语法篇
package example

import (
	"container/list"
	"errors"
	"fmt"
	"math"
	"runtime"
	"strings"
	"sync"
	"time"
)

// ------------------------变量和数据类型------------------------------//

// 变量的语法
func variableSyntax()  {
	// 变量声明（自动对内存区域进行初始化操作，初始化默认值）
	var name string
	// 批量声明
	var (
		age int
		flag bool
	)
	// 声明变量时，赋予默认值
	var contury string = "china"

	// 编译器推导类型的格式
	var city = "wuhan"

	// 短变量（只能方法中使用）
	gender:= "wucang"

	fmt.Println(name, age, flag, contury, city, gender)

	// 多重赋值（变量的左值和右值按从左到右的顺序赋值）
	city, gender = gender, city
	fmt.Println("city: " + city, "gender: " + gender)

	// 匿名变量
	_ = time.Now().Day()
}

// 数据类型
func dataTypeSyntax () {
	// 整形
	var number int8 = 23
	fmt.Println("number: ", number)
	fmt.Println("int8 range：", math.MinInt8, "~", math.MaxInt8)
	var count uint8 = 24
	fmt.Println("count: ", count)
	fmt.Println("uint8 maxValue: ", math.MaxUint8)

	fmt.Println("*********************************")

	// 浮点型
	var price = math.Pi
	fmt.Println("pi：", price)
	// 2位精度输出
	fmt.Printf("pi：%.2f\n", price)

	fmt.Println("*********************************")


	// 布尔类型
	aBool := true
	fmt.Println("aBool: ", aBool)

	fmt.Println("*********************************")

	// 字符串
	str := "hello world"
	fmt.Println("str: " + str)
	fmt.Printf("str: %s\n", str)

	// 定义多行字符串，使用反引号
	str = `第一行
	第二行
第三行
`
	fmt.Println("mutiple line：", str)

	// 字符类型
	var a byte = 'a'
	// 可以发现，byte 类型的 a 变量，实际类型是 uint8，其值为 'a'，对应的 ASCII 编码为 97。
	fmt.Printf("a: %d, 变量的类型：%T\n", a, a)

	var b rune = '你'
	// rune 类型的 b 变量的实际类型是 int32，对应的 Unicode 码就是 20320。
	fmt.Printf("b：%d, 变量的类型：%T\n", b, b)

	fmt.Println("*********************************")

	// 数据类型转换
	var c = 3.22
	// 浮点数在转换为整型时，会将小数部分去掉，只保留整数部分。
	fmt.Println("c：", int8(c))
}

// 指针语法
func pointerSyntax()  {
	var hint string = "hello"

	// 对字符串取地址, ptr类型为*string
	var ptr *string = &hint

	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)

	// 打印ptr的指针地址，指针值带有0x的十六进制前缀。
	fmt.Println("ptr value：", ptr)

	// 对指针进行取值
	value := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value：%s\n", value)

	// 使用指针修改值
	*ptr = "world"
	fmt.Println("修改后的hint：", hint)

	// 使用new()创建指针
	str := new(string)
	fmt.Println("指向默认值（空字符串）：", *str)

	*str = "hello"
	fmt.Println("更新后的值：", *str)

	var a = "hello"
	var b = a
	// 输出结果：内存地址不一样
	fmt.Println("a 地址：", &a)
	fmt.Println("b 地址：", &b)
}

// 常量语法
func constantsSyntax()  {
	const size = 4
	const (
		pi = 3.141592
		e = 2.718281
	)
	fmt.Println("size:", size)
	fmt.Println("pi: ", pi, " e: ", e)

	fmt.Println("*******************************")

	// GoGo 语言中现阶段没有枚举，可以使用 const 常量配合 iota 模拟枚举

	// 将 int 定义为 Weapon 类型
	type Weapon int
	// 标识后，const 下方的常量可以是默认类型的，默认时，默认使用前面指定的类型作为常量类型。该行使用 iota 进行常量值自动生成。
	// iota 起始值为 0，一般情况下也是建议枚举从 0 开始，让每个枚举类型都有一个空值，方便业务和逻辑的灵活使用。

	// 一个 const 声明内的每一行常量声明，将会自动套用前面的 iota 格式，并自动增加。这种模式有点类似于电子表格中的单元格自动填充。
	// 只需要建立好单元格之间的变化关系，拖动右下方的小点就可以自动生成单元格的值。
	const (
		Arrow Weapon = iota    // 开始生成枚举值, 默认为0
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	fmt.Println("Arrow: ", Arrow)
	fmt.Println("Blower: ", Blower)
}

// 类型别名
func typeAliasSyntax()  {

	// 将NewInt定义为int类型
	type NewInt int

	// 将int取一个别名叫IntAlias
	type IntAlias = int

	// 将a声明为NewInt类型
	var a NewInt

	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)

	// 将a2声明为IntAlias类型
	var a2 IntAlias

	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)

	// 结果显示a的类型是 main.NewInt，表示 main 包下定义的 NewInt 类型。a2 类型是 int。IntAlias 类型只会在代码中存在，
	// 编译完成时，不会有 IntAlias 类型。
}

// -------------------------变量和数据类型--------------------------------//

// -----------------------------容器-----------------------------------//

// 数组
func arraySyntax()  {
	// 数组
	const size = 4
	var arr [size]int

	// 数组长度
	fmt.Println("arr length: ", len(arr))
	// 默认值
	fmt.Println("arr value：", arr)

	// 数组赋值，数组越界，运行时会报出宕机
	arr[0] = 1
	fmt.Println("arr value: ", arr)

	// 需要确保大括号后面的元素数量与数组的大小一致
	var nameList = [3]string{"dazuo", "wang", "li"}
	fmt.Println(nameList)

	// 让编译器在编译时，根据元素个数确定数组大小
	var ageList = [...]int{1, 3, 4}
	fmt.Println(ageList)

	// 遍历数组
	for k, v := range ageList {
		fmt.Println(k, v)
	}
}

// 切片
func sliceSyntax()  {
	var nameList = [...]string{"dazuo", "wang", "li"}
	// 取出元素不包含结束位置对应的索引
	var subList = nameList[1:2]
	fmt.Println(subList)
	fmt.Println("中间至尾部：", nameList[1:])
	fmt.Println("开头至中间：", nameList[:2])
	fmt.Println("全部：", nameList[:])

	fmt.Println("********************************")

	// 声明字符串切片
	var strList []string
	// 声明整形切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	fmt.Println(strList, numList, numListEmpty)
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	// 声明但未使用的切片的默认值是 nil
	// 本来会在{}中填充切片的初始化元素，这里没有填充，所以切片是空的。但此时 numListEmpty 已经被分配了内存，但没有元素。
	// 因此和 nil 比较时是 false。
	fmt.Println(numListEmpty == nil)

	fmt.Println("********************************")

	a := make([]int, 2)
	fmt.Println("a len: ", len(a))
	fmt.Println("a value: ", a)

	b := make([]int, 2, 10)
	fmt.Println("b len: ", len(b))

	fmt.Println("********************************")
	// 每个切片会指向一片内存空间，这片空间能容纳一定数量的元素。当空间不能容纳足够多的元素时，切片就会进行“扩容”
	c := append(b, 1)
	fmt.Println("c value: ", c)

	// 修改切片值，操作的是同一片内存空间
	c[0] = 3
	fmt.Println(b)

	fmt.Println("********************************")

	// 复制切片
	var d = make([]int, 1, 10)
	i := copy(d, c[:])
	fmt.Println("copy count：", i)
	fmt.Println("d value: ", d)
}

// map映射
func mapSyntax()  {
	// 使用make创建map
	s := make(map[string]string)
	s["name"] = "dazuo"
	fmt.Println(s)
	fmt.Println(s["name"])

	// 声明map时填充元素
	m := map[string]string{
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}
	fmt.Println(m)

	// 遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 删除map中的键值对
	delete(m, "W")
	fmt.Println(m)

	// 并发环境中的map
	var scene sync.Map

	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

// list列表
func listSyntax () {
	// 创建一个列表实例
	myList := list.New()
	// 将 fist 字符串插入到列表的尾部，此时列表是空的，插入后只有一个元素。
	myList.PushBack("first")

	// 将数值 68 放入列表。此时，列表中已经存在 fist 元素，67 这个元素将被放在 fist 的前面。
	myList.PushFront(68)

	// 保存元素句柄
	element := myList.PushFront(67)

	fmt.Println(element)
	fmt.Println(myList)

	// 获取链表首部元素 67，然后找到链表下一个元素 68
	fmt.Println(myList.Front().Next().Value)

	// 通过句柄移除元素
	myList.Remove(element)

	fmt.Println("********************************")

	// 遍历列表
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

// -----------------------------容器-----------------------------------//

// -----------------------------流程控制-----------------------------------//

// 流程控制
func processControlSyntax()  {
	num := 10
	if num > 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// if 的特殊写法
	// 可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断
	if err := Connect; err != nil {
		fmt.Println("ye")
	}

	fmt.Println("*********************************")

	// for（循环结构）
	step := 2
	for ; step > 0; step-- {
		fmt.Println(step)
	}

	var i int
	for ; ; i++ {
		if i > 10 {
			fmt.Println(i)
			break
		}
	}

	fmt.Println("*********************************")

	// switch结构
	var a = "hello"
	switch a {
		case "hello":
			fmt.Println(1)
		case "world":
			fmt.Println(2)
		default:
			fmt.Println(0)
	}

	// 一个分支多个值
	var b = "mum"
	switch b {
	case "mum", "daddy":
		fmt.Println("family")
	}

	// 分支表达式
	var r int = 11
	switch {
		case r > 10 && r < 20:
			fmt.Println(r)
	}

	fmt.Println("*********************************")

	// 使用goto退出多层循环
	for x := 0; x < 10; x++ {
		if x == 2 {
			goto breakHere
		}
	}

	fmt.Println("in coming")
	// 手动返回，避免执行进入标签；此处如果不手动返回，在不满足条件时，也会执行标签处代码
	return

	// 标签
	breakHere:
		fmt.Println("done")

	fmt.Println("*******************************")

	fmt.Println("跳出指定循环")
	OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}

	fmt.Println("********************************")
	// continue 将结束当前循环，开启下一次的外层循环，而不是内层的循环。
	OuterFlag:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterFlag
			}
		}
	}
}

// 类型返回值
func Connect() int {
	return 1
}

// -----------------------------流程控制-----------------------------------//

// -----------------------------函数-----------------------------------//

// 函数语法
func funcSyntax()  {
	// 函数变量：在 Go 语言中，函数也是一种类型，可以和其他类型一样被保存在变量中。
	f := fire
	fmt.Printf("f type: %T\n", f)
	// 调用函数
	_ = f("dazuo")

	fmt.Println("*******************************")

	// 链式调用
	chainInvoke()

	fmt.Println("*******************************")

	// 匿名函数
	// 1.在定义时调用匿名函数
	func(data int) {
		fmt.Println(data)
	}(100)

	// 2.将匿名函数赋值给变量
	m := func() {
		fmt.Println("this is m func")
	}
	m()

	// 3.匿名函数用作回调函数
	visit([]int{1, 2, 3}, func(i int) {
		fmt.Println(i)
	})

	// 4.函数体类型实现接口
	var invoker Invoker
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function:", v)
	})
	invoker.Call("hello")

	// 5.TODO 函数闭包：http://c.biancheng.net/view/59.html

	fmt.Println("******************************")

	// defer语句
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(2)
	// 证明：defer语句在函数即将返回时，逆序执行
	fmt.Println(3)
}

// 回调函数
func visit(myList []int, f func(int))  {
	for _, v := range myList {
		f(v)
	}
}

func fire(msg string) string {
	fmt.Println("hello fire")
	return msg
}

// 链式调用
func chainInvoke()  {
	// 待处理的字符串列表
	myList := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 处理函数链
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	stringProcess(myList, chain)

	for _, str := range myList {
		fmt.Println(str)
	}
}

// 移除前缀
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

// 字符串处理函数
func stringProcess(myList []string, chain []func(string)string)  {
	for k, v := range myList {
		result := v

		for _, proc := range chain {
			result = proc(result)
		}

		// 将结果放回切片
		myList[k] = result
	}
}

//-------------------------- 函数实现接口 ---------------------//
// 调用器接口
type Invoker interface {
	Call (interface{})
}

// 函数定义为类型
type FuncCaller func (interface{})

// 实现接口的方法
func (f FuncCaller) Call (p interface{}) {
	// 调用f函数本体
	f(p)
}




//-------------------------- 定义错误 --------------------------//

// 自定义一个错误
func tokenError(text string) error {
	// 实例化结构体，并初始化结构体
	var errorStr = &errorString{text}
	return errorStr
}

// 定义结构体（包含错误字符串）
type errorString struct {
	s string
}

// 结构体添加方法（实现错误接口，返回错误描述）
func (e *errorString) Error() string {
	return e.s
}

// 运行时错误
func errorSyntax()  {
	// 使用 errors 包进行错误的定义
	var err = errors.New("this is an error")
	// 输出错误信息
	fmt.Println(err.Error())

	// 使用自定义错误
	err = tokenError("occur error")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 服务宕机恢复
func panicSyntax()  {
	// 延迟处理函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()

		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error", err)
		default:
			// 非运行时错误
			fmt.Println("error: ", err)
		}
	}()

	fmt.Println("手动触发宕机")
	panic("occur panic")
}

// -----------------------------函数-----------------------------------//

// -----------------------------结构体-----------------------------------//

// 定义结构体
type Point struct {
	x int
	y int
}

// 结构体语法
func structSyntax()  {
	// 1.基本的实例化形式
	p := Point{2, 3}
	fmt.Println(p.x)
	p2 := Point{}
	fmt.Println(p2.x)
	fmt.Printf("p2 type %T\n", p2)

	fmt.Println("****************************************")

	// 2.创建指针类型的结构体
	p3 := new(Point)
	fmt.Println("p3.x value: ", p3.x)
	// p3 为指针类型
	fmt.Printf("p3 type %T\n", p3)

	fmt.Println("****************************************")

	// 3.取结构体的地址实例化
	p4 := &Point{}
	fmt.Println("p4.x value: ", p4.x)
	// p4 为指针类型（对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作）
	fmt.Printf("p4 type %T\n", p4)

	fmt.Println("****************************************")

	// 4.使用“键值对”初始化结构体
	p5 := Point{
		x: 1,
		y: 2,
	}
	fmt.Println(p5.x)

	p6 := Point{2 ,3}
	fmt.Println(p6.x)

	// 5.初始化匿名结构体
	ins := struct {
		x int
		y int
	}{
		x: 1,
		y: 2,
	}
	fmt.Println(ins)
}

// -----------------------------接收器-----------------------------------//

// 定义一个背包
type Bag struct {
	items []int
}

// 往背包里面放物品
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

// 指针接收器加方法（传递的是引用）
func (b *Bag) Insert2(itemid int) {
	b.items = append(b.items, itemid)
}

// 非指针接收器加方法（Go 语言会在代码运行时将接收器的值复制一份）
func (b Bag) Insert3(itemid int) Bag {
	b.items = append(b.items, itemid)
	return b
}

// 定义一个类型
type NewInt int

// 为新类型添加方法
func (m NewInt) isZero() bool {
	return m == 0
}

// 类型内嵌
type Data struct {
	int
	float32
	bool
}

// 结构体内嵌
type BasicColor struct {
	R, G, B float32
}

type Color struct {
	BasicColor
	Alpha float32
}

// 结构体添加方法
func structMethodSyntax()  {
	// 1.面向过程的实现方法
	ins := new(Bag)
	Insert(ins, 1)
	fmt.Println("ins value: ", ins)

	fmt.Println("***************************")

	// 2.Go语言的结构体方法
	ins.Insert2(2)
	fmt.Println("ins value: ", ins)

	int2 := Bag{}
	int2 = int2.Insert3(2)
	fmt.Println("int2 value: ", int2)

	fmt.Println("***************************")

	// 3.为基本类型添加方法
	var i NewInt
	// 默认值是0
	fmt.Println(i.isZero())

	fmt.Println("***************************")

	// 4.类型内嵌
	ins2 := Data {
		int: 10,
		float32: 2.14,
		bool: true,
	}
	fmt.Println(ins2)

	ins3 := Data {
		10,
		2.14,
		true,
	}
	fmt.Println(ins3)

	fmt.Println("***************************")

	// 5.结构体内嵌
	var c Color
	// 设置基本颜色分量
	c.BasicColor.B = 1
	c.BasicColor.G = 2
	// 简化
	c.R = 3
	c.Alpha = 4
	fmt.Println(c)

	fmt.Println("***************************")

	// 6.初始化内嵌结构体
	cIns := Color{
		BasicColor{
			1, 2, 3,
		},
		4,
	}
	fmt.Println("cIns", cIns)
}

// -----------------------------结构体-----------------------------------//

// -----------------------------接口-----------------------------------//

// 定义一个数据写入器（接口）
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构，用于实现DataWriter
type file struct {

}

// 接收器添加方法WriteData，实现了DataWriter接口
func (d *file) WriteData(data interface{}) error  {
	// 模拟写入数据
	fmt.Println("writeData：", data)
	return nil
}

// 定义一个飞行接口
type Flyer interface {
	Fly()
}

// 定义一个行走接口
type Worker interface {
	Work()
}

// 定义一个鸟类
type bird struct {
}

// 接收器加方法（实现飞行动物接口）
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// 接收器加方法（实现行走接口）
func (b *bird) Work() {
	fmt.Println("bird: worker")
}

// 接口语法
func interfaceSyntax()  {
	f := new(file)

	// 声明一个接口
	var writer DataWriter

	// 将 *file 类型的 f 赋值给 DataWriter 接口的 writer，虽然两个变量类型不一致。但是 writer 是一个接口，
	// 且 f 已经完全实现了 DataWriter() 的所有方法，因此赋值是成功的。
	writer = f

	// 使用DataWriter接口进行数据写入
	_ = writer.WriteData("data")

	fmt.Println("*************************************")

	b1 := new(bird)
	var flyer Flyer = b1

	// 接口和类型转换——接口转换为其他接口
	fl1, _ := flyer.(Flyer)
	fmt.Printf("fl1 type: %T  fl1 value: %v\n", fl1, fl1)
	fl2, _ := flyer.(Worker)
	fmt.Printf("fl2 type: %T  fl2 value: %v\n", fl2, fl2)

	// 接口和类型转换——将接口转换为其他类型
	f2 := writer.(*file)
	fmt.Printf("f1 type: %T  f1 value: %v\n", f2, f2)

	fmt.Println("********************************")

	// 创建动物名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
	}

	for name, obj := range animals {
		fmt.Printf("name : %T\n", name)
		fmt.Printf("type : %T\n", obj)

		// 判断是否是飞行动物
		t, isFlyer := obj.(Flyer)
		fmt.Println("name: ", name, " t: ", t, " isFlyer: ", isFlyer)

		// 判断是否会行走
		w, isWorker := obj.(Worker)
		fmt.Println("name：", name,  " w: ", w, " isWorker: ", isWorker)

		if isFlyer {
			fmt.Printf("name：%s isFlyer: %v\n", name, isFlyer)
		}

		if isWorker {
			fmt.Printf("name：%s isWorker: %v\n", name, isWorker)
		}
	}

	fmt.Println("********************************")

	// 声明空接口
	var any interface{}

	// 将值保存到空接口中
	any = 1
	fmt.Println("any: ", any)

	any = "hello"
	fmt.Println("any: ", any)

	// 从空接口中获取值
	// 声明a变量, 类型int, 初始值为1
	var a int = 1

	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	var i interface{} = a

	// 编译器告诉我们，不能将i变量视为int类型赋值给b。
	//var b int = i

	// 使用类型断言
	var b int = i.(int)
	fmt.Printf("b type: %T  b value: %v", b, b)
}

// -----------------------------接口-----------------------------------//
