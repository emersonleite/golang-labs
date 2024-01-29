package main

import (
	"fmt"
	"time"

	"github.com/emersonleite/golang_labs/fibonacci"
)

func CalculateManyWithChannel(indexFibonacci int) {
	startCalculateManyWithChannel := time.Now()
	counter, fibonacci := fibonacci.CalculateManyWithChannel(indexFibonacci)
	durationCalculateManyWithChannel := time.Since(startCalculateManyWithChannel)
	fmt.Printf("Time elapsed for CalculateManyWithChannel %s with %d Interactions \n", durationCalculateManyWithChannel, counter)
	fmt.Println(fibonacci[indexFibonacci])
}

func CalculateMany(indexFibonacci int) {
	startCalculateMany := time.Now()
	counter, fibonacci := fibonacci.CalculateMany(indexFibonacci)
	durationCalculateMany := time.Since(startCalculateMany)
	fmt.Printf("Time elapsed for CalculateMany %s with %d Interactions \n", durationCalculateMany, counter)
	fmt.Println(fibonacci[indexFibonacci])
}

func calculateOne(indexFibonacci int) {
	startCalculateOne := time.Now()
	_, fibonacci := fibonacci.CalculateOne(indexFibonacci)
	durationCalculateOne := time.Since(startCalculateOne)
	fmt.Printf("Time elapsed for calculateOne %s\n", durationCalculateOne)
	fmt.Println(fibonacci)
}

func main() {

	const indexFibonacci = 10_000

	CalculateManyWithChannel(indexFibonacci)
	CalculateMany(indexFibonacci)
	calculateOne(indexFibonacci)

}
