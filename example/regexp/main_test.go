package regexp

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_string_Match(t *testing.T) {
	str := "/cms/api"
	result := strings.HasPrefix(str, "/cms")
	fmt.Println(result)
}

func Test_rege_find(t *testing.T) {
	text := `Hello 世界！123 Go.`
	reg := regexp.MustCompile(`[a-z]+`)
	// 查找连续的非小写字母
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
}

func Test_rege_match(t *testing.T) {
	text := `Hello 世界！123 Go.`
	// 查找行首以 H 开头，以空格结尾的字符串
	reg := regexp.MustCompile(`^H.[a-z]+\s`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
}
