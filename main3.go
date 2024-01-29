package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string)

	finalResult := ""

	go plus(2, 2, ch)
	finalResult += <-ch + ";"

	go plus(3, 6, ch)
	finalResult += " " + <-ch + ";"

	go multiple(7, 7, ch)
	finalResult += " " + <-ch + ";"

	go division(9, 3, ch)
	finalResult += " " + <-ch + ";"

	fmt.Println(finalResult)
}

func plus(x, y int, ch chan string) {
	ch <- fmt.Sprintf("%d + %d = "+strconv.Itoa(x+y), x, y)
}

func division(x, y int, ch chan string) {
	ch <- fmt.Sprintf("%d - %d = "+strconv.Itoa(x-y), x, y)
}

func multiple(x, y int, ch chan string) {
	ch <- fmt.Sprintf("%d * %d = "+strconv.Itoa(x*y), x, y)
}
