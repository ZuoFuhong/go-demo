// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package algorithm

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

// BruteForce 使用简单粗暴的方式，对主串和模式串进行逐个字符的比较。
func Test_BF(t *testing.T) {
	idx := bruteForce("aabcdee", "ee")
	fmt.Println(idx)
}

func bruteForce(str, pattern string) int {
	m := len(str)
	n := len(pattern)
	if m >= n {
		for i := 0; i < m-n+1; i++ {
			substr := str[i : i+n]
			for j := 0; j < n; j++ {
				if substr[j] != pattern[j] {
					break
				}
				if j == n-1 {
					return i
				}
			}
		}
	}
	return -1
}

// RabinKarp 用模式串的hash值和主串的局部hash值比较。
func Test_RK(t *testing.T) {
	str := "hello world"
	pattern := "wo"
	fmt.Println(rabinKarp(str, pattern))
}

func rabinKarp(str, pattern string) int {
	m := len(str)
	n := len(pattern)

	for i := 0; i < m-n+1; i++ {
		// 用模式串的hash值和主串的局部hash值比较。
		// 如果匹配，则进行精确比较；如果不匹配，计算主串中相邻子串的hash值。
		substr := str[i : i+n]
		if hash(pattern) == hash(substr) && compareString(substr, pattern) {
			return i
		}
	}
	return -1
}

func hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 可能发生hash冲突，因此需要进行精准比较
func compareString(substr, pattern string) bool {
	for n := len(substr) - 1; n >= 0; n-- {
		if substr[n] != pattern[n] {
			return false
		}
	}
	return true
}

// Knuth-Morris-Pratt 字符串查找算法
func Test_KMP(t *testing.T) {
}
