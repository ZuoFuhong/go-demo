package testify

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type Storage interface {
	Store(key, value string) (int, error)
	Load(key string) (string, error)
}

type RedisStorage struct {
}

func (s *RedisStorage) Store(key, value string) (int, error) {
	panic("There is not wifi!")
}

func (s *RedisStorage) Load(key string) (string, error) {
	panic("There is not wifi!")
}

// 测试用例，当真实对象不可用时，使用mock对象代替
type mockStorage struct {
	mock.Mock
}

func (ms *mockStorage) Store(key, value string) (int, error) {
	args := ms.Called(key, value)
	return args.Int(0), args.Error(1)
}

func (ms *mockStorage) Load(key string) (string, error) {
	args := ms.Called(key)
	return args.String(0), args.Error(1)
}

func Test_mock(t *testing.T) {
	mockS := &mockStorage{}
	mockS.On("Store", "name", "dazuo").Return(20, nil).Once()

	var storage Storage = mockS
	i, e := storage.Store("name", "dazuo")
	if e != nil {
		panic(e)
	}
	fmt.Println(i)
}
