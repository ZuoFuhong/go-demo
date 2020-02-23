package httptest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	City    string `json:"city"`
	Date    string `json:"date"`
	TemP    string `json:"temP"`
	Weather string `json:"weather"`
}

func GetWeatherInfo(api string) ([]Weather, error) {
	url := fmt.Sprintf("%s/weather?city=%s", api, "wuhan")
	resp, err := http.Get(url)

	if err != nil {
		return []Weather{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []Weather{}, fmt.Errorf("resp is didn't 200 OK:%s", resp.Status)
	}
	bodybytes, _ := ioutil.ReadAll(resp.Body)
	personList := make([]Weather, 0)

	err = json.Unmarshal(bodybytes, &personList)

	if err != nil {
		_ = fmt.Errorf("decode data fail")
		return []Weather{}, fmt.Errorf("decode data fail")
	}
	return personList, nil
}
