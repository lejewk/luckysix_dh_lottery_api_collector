package dh_lottery_test

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "luckysix/dh_lottery"
    "testing"
)

func TestApi(t *testing.T) {
    res, _ := dh_lottery.GetApi().Get(1)
    assert.Equal(t, res.FirstWinamnt , 0)
    fmt.Println(res)
}
