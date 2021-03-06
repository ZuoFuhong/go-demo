package goconvey

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 运行测试
// 1.终端窗口使用：go test -v
// 2.Web浏览器使用：在相应的目录执行 goconvey，然后访问http://localhost:8080

// 参考文章：https://www.cnblogs.com/WayneZeng/p/9290711.html
func Test_Add(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func Test_Substract(t *testing.T) {
	// 忽略断言
	SkipConvey("将两数相减", t, func() {
		So(Subtract(2, 1), ShouldEqual, 1)
	})
}

func Test_Multiply(t *testing.T) {
	Convey("将两数相乘", t, func() {
		So(Multiply(3, 2), ShouldEqual, 6)
	})
}

func Test_Division(t *testing.T) {
	Convey("将两数相除", t, func() {
		// 嵌套
		Convey("除以非 0 数", func() {
			num, err := Division(10, 2)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 5)
			// 忽略断言
			SkipSo(num, ShouldNotBeNil)
		})

		Convey("除以 0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})
	})
}
