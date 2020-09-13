package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

/**
 * 简单的Demo
 * 1.新建db.go文件
 * 2.使用mockgen：mockgen -source=db.go -destination=db_mock.go -package=main
 * 3.写测试用例 TestGetFromDB
 * 4.执行测试：go test . -cover -v
 */
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	// 打桩
	m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
