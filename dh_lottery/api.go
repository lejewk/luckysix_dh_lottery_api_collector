package dh_lottery

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "luckysix/entity"
    "net/http"
    "time"
)

var client = &http.Client{Timeout: 10 * time.Second}

type DhLottery struct {}

func GetApi() *DhLottery {
    return &DhLottery{}
}

func (s *DhLottery) Get(drwNo int) (*entity.DhLotteryRaw, error) {
    url := fmt.Sprintf("https://www.dhlottery.co.kr/common.do?method=getLottoNumber&drwNo=%d", drwNo)
    res, err := client.Get(url)

    if err != nil {
        return nil, err
    }

    defer res.Body.Close()
    resBody, err := ioutil.ReadAll(res.Body)

    if err != nil {
        return nil, err
    }

    dhLotteryRaw := &entity.DhLotteryRaw{}

    err = json.Unmarshal(resBody, dhLotteryRaw)
    if err != nil {
        return nil, err
    }

    return dhLotteryRaw, nil
}
