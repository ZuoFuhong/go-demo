package http_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	使用标准的httptest库完成http请求的mock测试
*/

var weatherResp = []Weather{
	{
		City:    "wuhan",
		Date:    "10-22",
		TemP:    "15℃~21℃",
		Weather: "rain",
	},
	{
		City:    "guangzhou",
		Date:    "10-22",
		TemP:    "15℃~21℃",
		Weather: "sunny",
	},
	{
		City:    "beijing",
		Date:    "10-22",
		TemP:    "1℃~11℃",
		Weather: "snow",
	},
}
var weatherRespBytes, _ = json.Marshal(weatherResp)

func Test_GetInfoOK(t *testing.T) {
	// 1.搭建http server，检查请求，模拟返回结果
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(weatherRespBytes)

		if r.Method != "GET" {
			t.Errorf("Except 'Get' got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/weather" {
			t.Errorf("Except to path '/person',got '%s'", r.URL.EscapedPath())
		}
		_ = r.ParseForm()
		topic := r.Form.Get("city")
		if topic != "wuhan" {
			t.Errorf("Except rquest to have 'city=wuhan',got '%s'", topic)
		}
	}))
	defer ts.Close()

	api := ts.URL
	fmt.Printf("Url:%s\n", api)

	// 2.调用方法
	resp, err := GetWeatherInfo(api)
	if err != nil {
		fmt.Println("ERR:", err)
	} else {
		fmt.Println("resp:", resp)
	}
}

func TestGetInfoUnauthorized(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write(weatherRespBytes)
		if r.Method != "GET" {
			t.Errorf("Except 'Get' got '%s'", r.Method)
		}

		if r.URL.EscapedPath() != "/weather" {
			t.Errorf("Except to path '/person',got '%s'", r.URL.EscapedPath())
		}
		_ = r.ParseForm()
		topic := r.Form.Get("city")
		if topic != "shenzhen" {
			t.Errorf("Except rquest to have 'city=shenzhen',got '%s'", topic)
		}
	}))
	defer ts.Close()
	api := ts.URL
	fmt.Printf("Url:%s\n", api)

	resp, err := GetWeatherInfo(api)
	if err != nil {
		t.Errorf("err: %s", err)
	} else {
		fmt.Println("resp:", resp)
	}
}
