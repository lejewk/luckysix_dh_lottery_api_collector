package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type LotteryApiRaw struct {
	TotSellamnt    int    `json:"totSellamnt"`
	DrwNoDate      string `json:"drwNoDate"`
	FirstWinamnt   int    `json:"firstWinamnt"`
	FirstPrzwnerCo int    `json:"firstPrzwnerCo"`
	DrwtNo6        int    `json:"drwtNo6"`
	DrwtNo5        int    `json:"drwtNo5"`
	DrwtNo4        int    `json:"drwtNo4"`
	DrwtNo3        int    `json:"drwtNo3"`
	DrwtNo2        int    `json:"drwtNo2"`
	DrwtNo1        int    `json:"drwtNo1"`
	BnusNo         int    `json:"bnusNo"`
	FirstAccumamnt int    `json:"firstAccumamnt"`
	DrwNo          int    `json:"drwNo"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string) (string, error) {
	r, err := myClient.Get(url)

	if err != nil {
		return "", err
	}

	defer r.Body.Close()
	responseData, err := ioutil.ReadAll(r.Body)
	return string(responseData), nil
}

func main() {
	response, err := getJson("https://www.dhlottery.co.kr/common.do?method=getLottoNumber&drwNo=927")

	if err != nil {
		panic(err)
	}

	//println(response)

	apiRaw := LotteryApiRaw{}
	err = json.Unmarshal([]byte(response), &apiRaw)

	if err != nil {
		panic(err)
	}

	fmt.Println(apiRaw)
}
