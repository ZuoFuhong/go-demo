package syntax

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	Go语言提供了反射功能，支持程序动态地访问变量的类型和值
*/
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Game interface {
	Play()
}

func (u User) Play() {
	println("this is play")
}

// 1.反射-实例的值信息
// 反射包中通过 reflect.飞1alue0f()函数获取实例的值信息
func Test_reflect_value(t *testing.T) {
	uValue := reflect.ValueOf(User{10, "dazuo"})
	uType := uValue.Type()
	println("value: ", uValue.String())
	println("获取类型：", uType.String())

	value := uValue.Field(1).Interface()
	println("字段值：", value.(string))
}

// 2.反射-通用方法
// 通过实例获取反射对象的 Type，直接使用 reflect.TypeOf()函数
func Test_reflect_common(t *testing.T) {
	// 返回一个Type类型的接口，使用者通过接口来获取对象的类型信息（只读）
	u := User{}
	stype := reflect.TypeOf(u)
	println("返回包含包名的类型名字，对于未命名类型返回的是空：", stype.Name())
	// Kind 返回该类型的底层基础类型
	println(stype.Kind().String())

	var game Game = u
	// 判断当前类型的实例是否能赋位给 type 为 u 的类型交量
	fmt.Println(stype.AssignableTo(reflect.TypeOf(game)))

	/// todo:
	//println("", stype.Implements(gType))

	println("返回一个类型的方法的个数: ", stype.NumMethod())
	// 通过方法名获取 Method
	method, b := stype.MethodByName("Play")
	println(b, method.Name)

	println("返回类型的包路径，如果类型是预声明类型或未命名类型，则返回空字符串：", stype.PkgPath())
	println("返回存放该类型的实例需要多大的字节空间: ", stype.Size())
}

// 2.1.反射-struct
func Test_reflect_struct(t *testing.T) {
	uType := reflect.TypeOf(User{10, "dazuo"})
	println("字段数目：", uType.NumField())
	println("通过整数索引获取字段: ", uType.Field(0).Type.String())
	println("取tag数据：", uType.Field(1).Tag)
	println("取tag数据：", uType.Field(1).Tag.Get("json"))
}

// 2.2.反射-func
func Test_reflect_func(t *testing.T) {
	calc := func(a, b int) int {
		return a + b
	}

	fType := reflect.TypeOf(calc)
	println("底层基础类型：", fType.Kind().String())
	println("输入的参数个数：", fType.NumIn())
	println("返回值个数：", fType.NumOut())
}

// 2.3.反射-map
func Test_reflect_map(t *testing.T) {
	data := make(map[string]interface{})
	mType := reflect.TypeOf(data)
	println("返回map key的type ", mType.Key().String())
}

// 3.从 Type 到 Value
// Type里面只有类型信息，所以直接从一个Type接口变量里面是无法获得实例的value的，但可以通过该Type构建一个新实例的Value。
func Test_reflectTypeToValue(t *testing.T) {
	var u = User{2, "zzz"}
	rType := reflect.TypeOf(u)

	// New返回的是一个Value，该Value的type为PtrTo(typ), 即Value的Type是指定typ的指针类型
	uValue := reflect.New(rType)
	// 如果你的变量是一个指针、map、slice、channel、Array。那么你可以使用Elem()来确定包含的类型。
	fmt.Println(uValue.Elem().String())

	uValue.Elem().Field(0).SetInt(10)
	uValue.Elem().Field(1).SetString("dazuo")
	newUser := uValue.Elem().Interface().(User)
	fmt.Println(newUser)

	// Zero返回的是一个typ类型的零值，注意返回的Value不能寻址，值不可改变
	zValue := reflect.Zero(rType)
	zUser := zValue.Interface().(User)
	zUser.Id = 10
	zUser.Name = "jing"
	fmt.Println("zero: ", zUser)
}

// 4.从 Value 到 Type
// 从反射对象 Value 到 Type 可以直接调用 Value 的方法，因为 Value 内部存放着到 Type 类型的指针。
func Test_reflectValueToType(t *testing.T) {
	uValue := reflect.ValueOf(User{10, "dazuo"})
	uType := uValue.Type()
	fmt.Println(uType.String())
}

// 5.从Value到实例
// Value 本身就包含类型和值信息，reflect 提供了丰富的方法来实现从value 到实例的转换。
func Test_reflectValueToInstance(t *testing.T) {
	uValue := reflect.ValueOf(User{10, "dazuo"})

	newUser := uValue.Interface().(User)
	fmt.Println(newUser)
}

// 6.从 Value 的指针到值
// 从一个指针类型的 Value 获得值类型 Value。
func Test_reflectPtrValueToValue(t *testing.T) {
	var u = User{2, "zzz"}
	rType := reflect.TypeOf(u)
	uValue := reflect.New(rType)

	// 如果 v 类型是接口，则 Elem()返回接口绑定的实例的 Value，如果 v 类型是指针，则返回指针值的 Value ，否则引起 panic
	fmt.Println(uValue.Elem())

	// 如果 v 是指针，则返回指针值的 Value，否则返回 v 自身，该函数不会引起 panic
	nValue := reflect.Indirect(uValue)
	fmt.Println(nValue)
}

// 7.Type 指针和值的相互转换
func Test_reflectPtrToValue(t *testing.T) {
	// 1）指针类型 Type 到值类型 Type
	pType := reflect.TypeOf(&User{})
	fmt.Println(pType)
	vType := pType.Elem()
	fmt.Println(vType)

	// 2）值类型 Type 到指针类型 Type
	npType := reflect.PtrTo(vType)
	println(npType.String())
}

// 8.Value值的可修改性
// 如果传进去的是一个指针，虽然接口内部转换的也是指针的副本，但通过指针还是可以访问到最原始的对象，所以此种情况获得的Value是可以修改的。
func Test_reflectValueModify(t *testing.T) {
	var u = User{2, "zzz"}
	rType := reflect.TypeOf(u)
	uValue := reflect.New(rType)
	// 通过 CanSet 判断是否能修改
	fmt.Println(uValue.Elem().CanSet())

	// 通过 Set 进行修改
	//uValue.Elem().Set()
}
