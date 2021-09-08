package main

import (
	"fmt"
	"time"
)

//WITH GO ROUTINES, ALL FUNCTION FINISHED JUST IN 10 SEC
func main() {
	for i := 0; i < 3; i++ {
		go counting(i)
	}

	var input string
	fmt.Scanln(&input)
}

func counting(n int) {
	for j := 0; j < 10; j++ {
		fmt.Printf("%d%d", n, j)
		fmt.Println()
		time.Sleep(time.Second)
	}
}

//BASIC SEQUENTIAL,
//IT TAKES 20 SEC TO FINISH ALL FUNCTION

// func main() {
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 10; j++ {
// 			fmt.Printf("%d%d", i, j)
// 			fmt.Println()
// 			time.Sleep(time.Second)
// 		}
// 	}

// 	var input string
// 	fmt.Scanln(&input)
// }
