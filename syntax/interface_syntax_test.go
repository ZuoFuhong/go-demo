package syntax

import (
	"fmt"
	"testing"
)

// 定义一个飞行接口
type Flyer interface {
	Fly()
}

// 定义一个行走接口
type Worker interface {
	Work()
}

// 定义结构体
type Bird struct {
}

// 接收器加方法（实现飞行动物接口）
func (b *Bird) Fly() {
	fmt.Println("bird: fly")
}

// 接收器加方法（实现行走接口）
func (b *Bird) Work() {
	fmt.Println("bird: worker")
}

// 接口语法
func TestInterfaceSyntax(t *testing.T) {
	b1 := new(Bird)
	var flyer Flyer = b1

	// 接口和类型转换——接口转换为其他接口
	fl1, _ := flyer.(Flyer)
	fmt.Printf("fl1 type: %T  fl1 value: %v\n", fl1, fl1)
	fl2, _ := flyer.(Worker)
	fmt.Printf("fl2 type: %T  fl2 value: %v\n", fl2, fl2)

	fmt.Println("********************************")

	// 创建动物名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(Bird),
	}

	for name, obj := range animals {
		fmt.Printf("name : %T\n", name)
		fmt.Printf("type : %T\n", obj)

		// 判断是否是飞行动物
		t, isFlyer := obj.(Flyer)
		fmt.Println("name: ", name, " t: ", t, " isFlyer: ", isFlyer)

		// 判断是否会行走
		w, isWorker := obj.(Worker)
		fmt.Println("name：", name, " w: ", w, " isWorker: ", isWorker)

		if isFlyer {
			fmt.Printf("name：%s isFlyer: %v\n", name, isFlyer)
		}

		if isWorker {
			fmt.Printf("name：%s isWorker: %v\n", name, isWorker)
		}
	}

	fmt.Println("********************************")

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
