package gomonkey

import (
	"fmt"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey/test/fake"
)

// 函数变量打桩
func Test_MockFuncVar(t *testing.T) {
	p := gomonkey.NewPatches()
	p.Reset()

	p.ApplyFuncVar(&fake.Marshal, func(v interface{}) ([]byte, error) {
		fmt.Println("ssss")
		return nil, nil
	})
	_, _ = fake.Marshal(nil)
}
