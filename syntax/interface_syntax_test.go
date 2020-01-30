package syntax

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"testing"
)

// 声明接口
type Flyer interface {
	Fly()
}

// 声明接口
type Worker interface {
	Work()
}

// 定义结构体
type Bird struct {
}

// 实现接口方法
func (b Bird) Fly() {
	fmt.Println("bird: fly")
}

// 实现接口方法
func (b *Bird) Work() {
	fmt.Println("bird: worker")
}

// 调用接口的方法
func TestInterfaceSyntax(t *testing.T) {
	bird := Bird{}
	bird.Fly()
	bird.Work()
	(&bird).Work()

	fmt.Println("*************************")

	// 接口和类型转换
	var flyer Flyer = bird
	fl, _ := flyer.(Flyer)
	fmt.Printf("fl1 type: %T  fl1 value: %v\n", fl, fl)
	fl.Fly()

	fmt.Println("*************************")

	var worker Worker = &bird
	tmpWorker, _ := worker.(Worker)
	fmt.Printf("fl2 type: %T  fl2 value: %v\n", tmpWorker, tmpWorker)
	tmpWorker.Work()
}

// 空接口 + 类型断言
// 语法：value, ok := x.(T)
// 其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。
// 该断言表达式会返回 x 的值（也就是 value）和一个布尔值（也就是 ok），可根据该布尔值判断 x 是否为 T 类型：
// - 如果 T 是具体某个类型，类型断言会检查 x 的动态类型是否等于具体类型 T。如果检查成功，类型断言返回的结果是 x 的动态值，其类型是 T。
// - 如果 T 是接口类型，类型断言会检查 x 的动态类型是否满足 T。如果检查成功，x 的动态值不会被提取，返回值是一个类型为 T 的接口值。
// - 无论 T 是什么类型，如果 x 是 nil 接口值，类型断言都会失败。
func Test_emptyInterface(t *testing.T) {
	// 声明空接口（空接口可以保存任何类型）
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

// error 接口，返回错误信息
func Test_error(t *testing.T) {
	s, e := getName()
	if e != nil {
		panic(e)
	}
	fmt.Println(s)
}

func getName() (string, error) {
	return "dazuo", errors.New("this is error")
}

// 自定义错误类型
func Test_custom_error(t *testing.T) {
	i, e := getAge()
	if e != nil {
		baseError, _ := e.(BaseError)
		fmt.Println("自定义异常：", baseError.Error())

		panic(e)
	}
	fmt.Println(i)
}

type BaseError struct {
	Code int
}

func (e BaseError) Error() string {
	return "code = " + strconv.Itoa(e.Code)
}

func getAge() (int, error) {
	return 22, BaseError{1}
}
