// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package testify

import (
	"testing"
	"time"
	"unicode"
)

// 示例：go test -v assert_test.go
func TestBuiltin(t *testing.T) {
	t.Logf("%s", t.Name())
	time.Sleep(time.Second)
	t.Skip()
	//t.Fatal()
}

func TestPalindrome(t *testing.T) {
	isPalindrome := IsPalindrome("boob")
	t.Log(isPalindrome)
}

// IsPalindrome 判断字符串是否是回文字符串
// 忽略字母大小写，以及非字母字符
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// 覆盖率
// go test -v -run=Coverage assert_test.go eval.go
// go test -v -run=Coverage -coverprofile=c.out assert_test.go eval.go
// go tool cover -html=c.out
func TestCoverage(t *testing.T) {
	Eval('+')
	Eval('-')
	Eval('*')
	Eval('/')
}

// 基准测试
// go test -bench=.
// 标记 -bench 的参数指定了要运行的基准测试。它是一个匹配 Benchmark函数名称的正则表达式，它的默认值不匹配任何函数。
// 模式"." 使它匹配 word 中所有的基准测试函数，因为这里只有一个基准测试函数，所以和指定 -bench=IsPalindrome效果一样。
//
// 基准测试运行期开始的时候并不清楚这个操作的耗时长短，所以开始的时候它使用了比较小的N值来做检测，然后为了检测稳定的运行时间，
// 推断出足够大的N值。
// 使用基准测试函数来实现循环而不是在测试驱动程序中调用代码的原因是，在基准测试函数中的循环外面可以执行一些必要的初始化代码
// 并且这段时间不加到每次迭代的时间中。
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("abcba")
	}
}
