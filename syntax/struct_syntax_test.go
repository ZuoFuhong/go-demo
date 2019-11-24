package syntax

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 使用type关键字 定义结构体
type Point struct {
	x int
	y int
}

/*
	结构体
      1.定义结构体
      2.结构体的实例化
      3.为结构体添加方法
*/
func TestStructSyntax(t *testing.T) {
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

	p6 := Point{2, 3}
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

// 定义一个结构体
type Bag struct {
	num int
}

func (b *Bag) Insert(num int) {
	b.num += num
}

// 为结构体添加方法
func TestStructMethodSyntax(t *testing.T) {
	b := new(Bag)
	b.Insert(1)
	b.Insert(2)
	fmt.Println("b num: ", b.num)

	fmt.Println("***************************")
}

// 定义一个类型
type NewInt int

// 为新类型添加方法
func (m NewInt) isZero() bool {
	return m == 0
}

// 为基本类型添加方法
func TestBasicTypeMethod(t *testing.T) {
	var i NewInt
	// 默认值是0
	fmt.Println(i.isZero())
}

// 指定字段的tag，实现json字符串的首字母小写
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"-"`
	Wechat string `json:"wechat,omitempty"`
}

// 将结构体序列化为 JSON
func TestStructMarshal(t *testing.T) {
	/*
			json.Marshal() JSON输出的时候必须注意：
		      1）首字母为小写时，为private字段，不会转换
		      2）tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中
		      3）字段的tag是"-"，那么这个字段不会输出到JSON
		      4）tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中，
		         比如 false、0、nil、长度为 0 的 array，map，slice，string
		      5）如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，
		         那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
	*/
	person := Person{"dazuo", 1, ""}
	data, _ := json.Marshal(person)
	fmt.Println(string(data))
}

// JSON 数据转换成 Go 类型的值（Decode）
func TestStructUnmarshal(t *testing.T) {
	data := []byte(`{"name":"dazuo","age":22}`)
	person := new(Person)
	err := json.Unmarshal(data, &person)
	if err != nil {
		_ = fmt.Errorf("Can not decode data: %v\n", err)
	}
	fmt.Printf("%v\n, 类型 = %T\n", *person, person)
}
