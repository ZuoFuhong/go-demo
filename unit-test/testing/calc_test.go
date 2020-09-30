package testing

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := 2
	b := 3
	ret := Add(a, b)
	if ret != a+b {
		t.Fatalf("Expect：%d Actual：%d", a+b, ret)
	}
}

func TestSubtests(t *testing.T) {
	// 嵌套
	t.Run("TestSubstract", func(t *testing.T) {
		a := 3
		b := 2
		ret := Substract(a, b)
		if ret != a-b {
			t.Fatalf("Expect：%d Actual：%d", a-b, ret)
		}
	})
}
