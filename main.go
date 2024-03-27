package main

import (
	"fmt"
)

func main() {

	texts := []string{"SRSSRRR", "RSSRR", "SSSRRRRS", "SRRSSR", "SSRSRRR"}
	for _, txtcase := range texts {
		fmt.Println(txtcase, checkSR(txtcase))

	}

}

func checkSR(seq string) string {
	shots := 0
	rev := 0
	lastShot := 'R'

	for _, act := range seq {
		switch act {
		case 'S':
			shots++
			lastShot = 'S'

		case 'R':
			if shots == 0 {
				fmt.Println("R fist :", shots)
				return "Bad boy"
			}
			rev++
			lastShot = 'R'

		}
	}
	fmt.Println("S Count :", shots)
	fmt.Println("R Count :", rev)

	if shots >= rev || lastShot != 'R' {
		return "Bad boy"
	}

	return "Good boy"
}
