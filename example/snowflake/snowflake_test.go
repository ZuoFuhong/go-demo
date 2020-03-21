package snowflake

import (
	"fmt"
	"github.com/golang/glog"
	"testing"
)

/*
	https://gitee.com/GuaikOrg/go-snowflake?_from=gitee_search
*/
func Test_generate(t *testing.T) {
	// NewSnowflake(datacenterid, workerid int64) (*Snowflake, error)
	// 参数1 (int64): 数据中心ID (可用范围:0-31)
	// 参数2 (int64): 机器ID    (可用范围:0-31)
	// 返回1 (*Snowflake): Snowflake对象 | nil
	// 返回2 (error): 错误码
	s, err := NewSnowflake(int64(0), int64(0))
	if err != nil {
		glog.Error(err)
		return
	}
	// 生成唯一ID
	id := s.NextVal()
	fmt.Println(id)

	// 通过ID获取生成ID时的时间戳
	timestamp := GetGenTimestamp(id)
	fmt.Println(timestamp)
}
