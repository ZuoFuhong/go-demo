package gomonkey

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
)

type Calc struct {
	val int
}

func (c *Calc) Incr() {
	c.val += 1
}

func (c *Calc) IncrVal(v int) {
	c.val += v
}

// 成员方法打桩
func TestMockMethod(t *testing.T) {
	p := gomonkey.NewPatches()
	p.Reset()

	var c *Calc
	p.ApplyMethod(reflect.TypeOf(c), "Incr", func(c *Calc) {
		fmt.Println("Incr ...")
		c.val += 1
	})

	p.ApplyMethod(reflect.TypeOf(c), "IncrVal", func(c *Calc, v int) {
		fmt.Println("IncrVal ...")
		c.val += v
	})

	calc := new(Calc)
	calc.Incr()
	calc.IncrVal(10)
	fmt.Printf("val: %d\n", calc.val)
}
