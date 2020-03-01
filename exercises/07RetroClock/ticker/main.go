package main

import (
	"fmt"
	"time"

	"github.com/chandrasekarar/learngo2/exercises/07RetroClock/digits"
	"github.com/inancgumus/screen"
)

func main() {
	digs := digits.Digits

	for shift := 0; ; shift++ {
		screen.Clear()
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

		for line := range clock[0] {
			l := len(clock)
			// this sets the beginning and the ending placeholder positions (indexes).
			// shift%l prevents the indexing error.
			s, e := shift%l, l

			if shift%(l*2) >= l {
				s, e = 0, s
			}

			for j := 0; j < l-e; j++ {
				fmt.Print("     ")
			}

			// draw the digits starting from 's' to 'e'
			for i := s; i < e; i++ {
				next := clock[i][line]
				if clock[i] == digits.Colon && sec%2 == 0 {
					next = "   "
				}

				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}
}
