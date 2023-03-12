package main

import (
	"fmt"
	"time"
)

func main() {
	// Функция подсчета времени запуска
	calcRunTime := func() (dur *time.Duration, err error) {
		var mskLoc *time.Location
		if mskLoc, err = time.LoadLocation("Europe/Moscow"); err != nil {
			return nil, err
		}

		mskNow := time.Now().In(mskLoc)
		year, month, day := mskNow.Date()
		planned := time.Date(year, month, day, 10, 56, 0, 0, mskLoc)

		fmt.Println(planned)
		d := planned.Sub(mskNow)

		fmt.Println(d)
		return &d, nil
	}

	var delay *time.Duration
	t := time.NewTimer(time.Second)
	var mtxErr error
	for {
		if delay, mtxErr = calcRunTime(); mtxErr != nil {
			return
		}

		t.Reset(*delay)

		select {
		case <-t.C:
			fmt.Println("Yeee")

		}
	}
}
