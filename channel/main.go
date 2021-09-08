package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculate Random Values: {}", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Printf("Only Executes after another goroutine performs a receive on the channel %d", value)
	fmt.Println()
}

func main() {
	fmt.Println("Go Channel Tutorial")

	valueChannel := make(chan int, 2)
	defer close(valueChannel)

	go CalculateValue(valueChannel)
	go CalculateValue(valueChannel)

	values := <-valueChannel

	fmt.Println(values)
	time.Sleep(1000 * time.Millisecond)
}
