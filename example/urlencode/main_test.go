package urlencode

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_URLEncode(t *testing.T) {
	v := url.Values{}
	v.Add("abc", "1")
	v.Add("country", "中国")
	encodeStr := v.Encode()
	fmt.Println(encodeStr) // abc=1&country=%E4%B8%AD%E5%9B%BD

	tmpStr := url.QueryEscape("中国")
	fmt.Println(tmpStr) // %E4%B8%AD%E5%9B%BD

	fmt.Println([]byte("中国")) // [228 184 173 229 155 189]

	/*
			URLEncode是对字符编码

		    如下字符不会被编码，是安全的：
		    - 1.大写字母A-Z
		    - 2.小写字母a-z
		    - 3.数字 0-9
		    - 4.标点符 '.' '-' '*' and '_'

		    所有其他的字符都被认为是不安全的，首先都根据指定的编码scheme被转换为1个或者多个字节。然后每个字节都被表示成'%xy'格式的由3个字符
		    组成的字符串，xy是字节的2位16进制的表达（xy is the two-digit hexadecimal representation of the byte），推荐的编码
		    scheme为UTF-8,然而，出于兼容性的考虑，如果没有制定编码的scheme，那么将使用当前操作系统的编码的scheme。

		    示例：如果编码scheme是UTF-8
		      中国         %E4%B8%AD%E5%9B%BD
		      提取字节数组：[228 184 173 229 155 189]
		      对应的16进制：[E4  B8  AD  E5  9B  BD]
	*/
}

func Test_URLDecode(t *testing.T) {
	encodeStr := "%E4%B8%AD%E5%9B%BD"
	decodeStr, err := url.QueryUnescape(encodeStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeStr)
}
