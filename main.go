package main

import (
    "fmt"
    "luckysix/database"
    "luckysix/dh_lottery"
)

func main() {
    fmt.Println("> Start Process")
    db, err := database.GetDB()

    if err != nil {
        panic(err)
    }

    currentRound := db.GetLatestRound()

    fmt.Println("> Latest Saved Round:", currentRound)

    dhLotteryApi := dh_lottery.GetApi()
    for {
        currentRound = currentRound + 1

        fmt.Println("> Request Next Round:", currentRound)

        dhLotteryApiRaw, err := dhLotteryApi.Get(currentRound)

        if err != nil {
            fmt.Println("> Err, Failed Request")
            break
        }

        if dhLotteryApiRaw.ReturnValue != "success" {
            fmt.Println("> Err, Not Exists Next Round")
            break
        }
        saved := db.Insert(dhLotteryApiRaw)

        if saved {
            fmt.Println("> Success Insert. Round:", currentRound)
        } else {
            fmt.Println("> Failed Insert Database. Round:", currentRound)
        }
    }

    fmt.Println("> Done.")
}
