package testing

import "testing"

// 覆盖率
// go test -v -run=Coverage eval_test.go eval.go
// go test -v -run=Coverage -coverprofile=c.out assert_test.go eval.go
// go tool cover -html=c.out
func TestCoverage(t *testing.T) {
	Eval('+')
	Eval('-')
	Eval('*')
	Eval('/')
}
