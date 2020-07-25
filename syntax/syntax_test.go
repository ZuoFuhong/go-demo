package syntax

import (
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"
)

/*
 	变量
	  1.变量的声明
      2.变量的初始化（默认值）
	  3.短变量
      4.多重赋值
      5.匿名变量
*/
func TestVariableSyntax(t *testing.T) {
	// 变量声明（自动对内存区域进行初始化操作，初始化默认值）
	var name string
	// 批量声明
	var (
		age  int
		flag bool
	)
	// 声明变量时，赋予默认值
	var contury string = "china"

	// 编译器推导类型的格式
	var city = "wuhan"

	// 短变量（只能方法中使用）
	gender := "wucang"

	fmt.Println(name, age, flag, contury, city, gender)

	// 多重赋值（变量的左值和右值按从左到右的顺序赋值）
	city, gender = gender, city
	fmt.Println("city: "+city, "gender: "+gender)

	// 匿名变量
	_ = time.Now().Day()
}

/*
	数据类型
      1.整型
      2.浮点型
      3.布尔型
      4.字符型
*/
func TestDataTypeSyntax(t *testing.T) {
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
	// 和C语言的字符串不同，Go语言中的字符串内容是不可变更的。在字符串作为参数传递给fmt.Println函数时，
	// 字符串的内容并没有被复制——传递的仅仅是字符串的地址和长度（字符串的结构在reflect.StringHeader中定义）
	str := "hello world"
	fmt.Println("str: " + str)
	fmt.Printf("str: %s\n", str)

	// 定义多行字符串，使用反引号
	str = `第一行
	第二行
第三行
`
	fmt.Println("mutiple line：", str)

	/*
			Go语言的字符有以下两种：
		      1）一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
		      2）另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。
		         rune 类型等价于 int32 类型。

			  byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题，例如 var ch byte = 'A'，
		      字符使用单引号括起来。
	*/

	var a byte = 'a'
	fmt.Printf("a: %d, 变量的类型：%T\n", a, a)
	var b rune = '你'
	fmt.Printf("b：%d, 变量的类型：%T\n", b, b)

	// 字符串可以转换为字节数组
	bytes := []byte(str)
	fmt.Println(len(bytes))
	// 也可以转换为 Unicode 的字数组
	runes := []rune(str)
	fmt.Println(runes)
}

type point struct {
	x int
	y int
}

func Test_printf(t *testing.T) {
	var a int = 112
	var b float32 = 10.232
	var c byte = 48
	var s string = "100"
	var stru point = point{x: 1, y: 2}
	var sp = &stru

	fmt.Printf("|%d| \t |%5d| \t |%-5d| \t |%+d|\n", a, a, a, a)                             // 打印整型
	fmt.Printf("八进制 = %#o, 16进制 = %#x, 二进制 = %b\n", a, a, a)                                 // 进制转换显示
	fmt.Printf("%f \t %.2f \n", b, b)                                                        // 打印浮点型
	fmt.Printf("%g\n", b)                                                                    // 用最少的数字来表示
	fmt.Printf("%q\n", c)                                                                    // 打印单引号
	fmt.Printf("%t\n", true)                                                                 // 打印true或false
	fmt.Printf("|%s| \t |%5s| \t 字符串的16进制表示 = 0x%x\n", s, s, s)                              // 打印字符串字符串
	fmt.Printf("%v \t %+v \t %#v \n", stru, stru, stru)                                      // 以默认的方式打印变量的值
	fmt.Printf("%p \t %#p \n", sp, sp)                                                       // 打印指针地址
	fmt.Printf("a = %T, b = %T, c = %T, s = %T, stru = %T, sp = %T\n", a, b, c, s, stru, sp) // 打印变量类型
}

/*
	数据类型转换
*/
func TestTypeConv(t *testing.T) {
	// 1.浮点数在转换为整型时，会将小数部分去掉，只保留整数部分。
	var c = 3.22
	fmt.Println("c：", int8(c))

	// 2.整型转字符串类型，使用strconv.Itoa(int) 函数
	var d = 23
	sd := strconv.Itoa(d)
	fmt.Printf("sd变量类型  %T\n", sd)

	// 3.字符转整型
	num, err := strconv.Atoi(sd)
	if err == nil {
		fmt.Printf("type:%T value:%#v\n", num, num)
	} else {
		fmt.Printf("%v 转换失败！", num)
	}

	// 4.Parse 系列函数用于将字符串转换为指定类型的值，其中包括 ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
	b, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Printf("转换成功 b = %v\n", b)
	} else {
		fmt.Printf("转换失败")
	}

	// 5.Format 系列函数用于将给定类型数据格式化为字符串类型的功能，
	// 其中包括 FormatBool()、FormatInt()、FormatUint()、FormatFloat()。
	sBool := strconv.FormatBool(true)
	fmt.Println("sBool: ", sBool)

	// 6.Append 系列函数用于将指定类型转换成字符串后追加到一个切片中，
	// 其中包含 AppendBool()、AppendFloat()、AppendInt()、AppendUint()。

	// 7.interface类型转换
	var e interface{}
	e = "32"

	// 如下转换，会跑异常
	//_ = e.(int)

	f, ok := e.(int)
	// ok 是bool类型，判断转换状态
	fmt.Printf("f = %v, ok = %v\n", f, ok)

	g, ok := e.(string)
	fmt.Printf("g = %v, ok = %v\n", g, ok)
}

// 常量和const关键字
func TestConstSyntax(t *testing.T) {
	const size = 4
	const (
		pi = 3.141592
		e  = 2.718281
	)
	fmt.Println("size:", size)
	fmt.Println("pi: ", pi, " e: ", e)
}

func Test_iota(t *testing.T) {
	// Go语言中现阶段没有枚举，可以使用 const 常量配合 iota 模拟枚举

	// 标识后，const 下方的常量可以是默认类型的，默认时，默认使用前面指定的类型作为常量类型。该行使用 iota 进行常量值自动生成。
	// iota 起始值为 0，一般情况下也是建议枚举从 0 开始，让每个枚举类型都有一个空值，方便业务和逻辑的灵活使用。

	// 一个 const 声明内的每一行常量声明，将会自动套用前面的 iota 格式，并自动增加。这种模式有点类似于电子表格中的单元格自动填充。
	// 只需要建立好单元格之间的变化关系，拖动右下方的小点就可以自动生成单元格的值。
	const (
		Arrow int = iota // 开始生成枚举值, 默认为0
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	fmt.Println("Arrow: ", Arrow)
	fmt.Println("Blower: ", Blower)
}

/*
	使用type关键字 定义 类型别名
*/
func Test_TypeAlias(t *testing.T) {
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

/*
	流程控制
      1.if else（分支结构）
      2.for（循环结构）
      3.switch语句
      4.goto语句
*/
func TestProcessControlSyntax(t *testing.T) {
	num := 10
	if num > 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// if 的特殊写法
	// 可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断
	if num++; num > 0 {
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
			fmt.Println("for i = ", i)
			break
		}
	}

	// 无限循环
	var w = 0
	for {
		if w > 10 {
			break
		}
		w++
	}

	// 自由一个条件的循环
	var n = 0
	for n < 10 {
		if n > 11 {
			break
		}
		n++
	}

	fmt.Println("*********************************")

	// switch结构
	// 不需要用break明确退出一个case
	// 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case；
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
		fallthrough
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

// Deprecated: Use strings.HasPrefix instead.
func deprecatedMethod() {
}
