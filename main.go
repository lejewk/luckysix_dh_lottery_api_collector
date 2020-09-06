package main

import (
    "fmt"
    "luckysix/database"
    "luckysix/dh_lottery"
)

func main() {
    db, err := database.GetDB()

    if err != nil {
        panic(err)
    }

    latestDbRow := db.GetLatest()

    currentRound := 0
    if latestDbRow != nil {
        currentRound = latestDbRow.DrwNo
    }

    dhLotteryApi := dh_lottery.GetApi()
    for {
        currentRound = currentRound + 1
        dhLotteryApiRaw, err := dhLotteryApi.Get(currentRound)

        if err != nil {
            break
        }

        if dhLotteryApiRaw.ReturnValue != "success" {
            break
        }
        db.Insert(dhLotteryApiRaw)
        fmt.Println("Round: ", currentRound)
    }
}
