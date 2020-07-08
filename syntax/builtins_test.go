// Copyright (c) 2020 ZuoFuhong. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package syntax

import (
	"container/list"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"testing"
	"time"
)

// 内建函数的使用

func Test_sort(t *testing.T) {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

func Test_builtin_Strings(t *testing.T) {
	s := "hello world"
	fmt.Println(strings.Contains(s, "hello"))
	fmt.Println(strings.HasPrefix(s, "hello"))
}

func Test_string_format(t *testing.T) {
	str := fmt.Sprintf("%s-%d", "name", 22)
	fmt.Println(str)
}

func Test_regular(t *testing.T) {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))
}

func Test_time(t *testing.T) {
	now := time.Now()
	fmt.Println(now)

	then := time.Date(
		2020, 06, 20, 14, 30, 30, 0, time.UTC)
	fmt.Println(then)

	fmt.Println(then.Year())
	fmt.Println(then.Before(now))
	fmt.Println(then.Equal(now))

	diff := now.Sub(then)
	fmt.Println(then.Add(diff))

	fmt.Println("#################################")

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())

	fmt.Println("#################################")

	YYYYMMDDHHMMSS := "2006-01-02 15:04:05"
	fmt.Println(time.Parse(YYYYMMDDHHMMSS, "2020-06-20 16:35:35"))
	fmt.Println(now.Format(YYYYMMDDHHMMSS))
}

func Test_random(t *testing.T) {
	// 默认的数字生成器是确定性的，所以默认情况下它每次都会生成相同的数字序列。
	// 要产生不同的序列，给它一个变化的种子。
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())

	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))
}

func Test_url_parse(t *testing.T) {
	str := "https://time.zxcs.org/column/article/8701"
	u, err := url.Parse(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
}

func Test_md5(t *testing.T) {
	m := md5.New()
	m.Write([]byte("123"))
	fmt.Printf(hex.EncodeToString(m.Sum(nil)))
}

func Test_sha1(t *testing.T) {
	s := "12345"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Printf("%x\n", bs)
}

func Test_base64(t *testing.T) {
	sEnc := base64.StdEncoding.EncodeToString([]byte("123"))
	fmt.Println(sEnc)

	uDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(uDec))

	// URLEncoding
	uEnc := base64.URLEncoding.EncodeToString([]byte("https://www.xx.com/s?wd=中国"))
	fmt.Println(uEnc)
	uDec, _ = base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

// 指定字段的tag，实现json字符串的首字母小写
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"-"`
	Wechat string `json:"wechat,omitempty"`
}

// 将结构体序列化为 JSON
func Test_struct_marshal(t *testing.T) {
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
func Test_struct_unmarshal(t *testing.T) {
	data := []byte(`{"name":"dazuo","age":22}`)
	person := new(Person)
	err := json.Unmarshal(data, &person)
	if err != nil {
		_ = fmt.Errorf("Can not decode data: %v\n", err)
	}
	fmt.Printf("%v\n, 类型 = %T\n", *person, person)
}

func Test_map_marshal(t *testing.T) {
	data := make(map[string]string, 0)
	data["name"] = "dazuo"
	data["age"] = "24"
	bytes, _ := json.Marshal(data)
	fmt.Println(string(bytes))

	data2 := make(map[string]string, 0)
	_ = json.Unmarshal(bytes, &data2)
	fmt.Println(data2)
}

func Test_List_marshal(t *testing.T) {
	tmpList := list.New()
	tmpList.PushFront("dazuo")
	tmpList.PushFront("age")

	bytes, _ := json.Marshal(tmpList)
	fmt.Println(string(bytes)) // list无法序列化
}

func Test_array_marshal(t *testing.T) {
	tmpArr := []string{"dazuo", "age"}
	bytes, _ := json.Marshal(tmpArr)
	fmt.Println(string(bytes))

	tmpArr2 := make([]string, 0, 0)
	_ = json.Unmarshal(bytes, &tmpArr2)
	fmt.Println(tmpArr2)
}
