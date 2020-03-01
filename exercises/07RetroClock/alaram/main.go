package main

import (
	"fmt"
	"time"

	"github.com/chandrasekarar/learngo2/exercises/07RetroClock/digits"
	"github.com/inancgumus/screen"
)

func main() {
	digs := digits.Digits
	screen.Clear()
	for {
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]digits.Digit{
			digs[hour/10], digs[hour%10],
			digits.Colon,
			digs[min/10], digs[min%10],
			digits.Colon,
			digs[sec/10], digs[sec%10],
		}

		alarmed := sec%10 == 0
		if alarmed {
			screen.Clear()
			fmt.Println("ALARAM")
		} else {
			for line := range clock[0] {
				for i, dig := range clock {
					next := clock[i][line]
					if dig == digits.Colon && sec%2 == 0 {
						next = "   "
					}
					fmt.Print(next, "  ")
				}
				fmt.Println()
			}
		}
		time.Sleep(time.Second)
	}

}
