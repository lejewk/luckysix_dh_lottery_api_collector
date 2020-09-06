package database

import (
    "database/sql"
    "fmt"
    "luckysix/entity"
)
import _ "github.com/go-sql-driver/mysql"

type CRUD struct {
    db *sql.DB
}

func GetDB() (*CRUD, error) {
    db, err := sql.Open("mysql", "luckysix:1111@/luckysix")
    if err != nil {
        panic(err)
    }

    return &CRUD{db: db}, nil
}

// 최근 회차 조회
func (s *CRUD) GetLatest() *entity.DhLotteryRaw {
    dhLotteryRaw := entity.DhLotteryRaw{}
    err := s.db.QueryRow("SELECT * FROM dh_lottery_raw ORDER BY round DESC LIMIT 1").Scan(&dhLotteryRaw)
    if err != nil {
        return nil
    }
    return &dhLotteryRaw
}

func (s *CRUD) FindByRound(round int) *entity.DhLotteryRaw {
    dhLotteryRaw := entity.DhLotteryRaw{}
    err := s.db.QueryRow("SELECT * FROM dh_lottery_raw WHERE round = ?", round).Scan(&dhLotteryRaw)
    if err != nil {
        return nil
    }
    return &dhLotteryRaw
}

func (s *CRUD) Insert(raw *entity.DhLotteryRaw) bool {
    res, err := s.db.Exec(
        "INSERT INTO dh_lottery_raw(round, draw_date, no_1, no_2, no_3, no_4, no_5, no_6, bonus_no, first_winner_amount, first_winner_count, total_sell_amount) " +
        "VALUE(?,?,?,?,?,?,?,?,?,?,?,?)",
        raw.DrwNo,
        raw.DrwNoDate,
        raw.DrwtNo1,
        raw.DrwtNo2,
        raw.DrwtNo3,
        raw.DrwtNo4,
        raw.DrwtNo5,
        raw.DrwtNo6,
        raw.BnusNo,
        raw.FirstWinamnt,
        raw.FirstPrzwnerCo,
        raw.TotSellamnt)

    if err != nil {
        fmt.Println(err)
        return false
    }

    affected, _ := res.RowsAffected()
    return affected > 0
}
