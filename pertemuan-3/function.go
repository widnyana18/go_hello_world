package main

import (
	"fmt"
	"math"
)

func luas(r int) {
	var l = math.Pi * math.Pow(float64(r), 2)

	if r%2 == 0 {
		fmt.Printf("Luas lingkaran: %2f\n", l)
		return
	}

	fmt.Printf("Beda luas: %2f\n", l)
}

func order(name string, items ...string) {
	fmt.Printf("name: %s\n", name)

	fmt.Println("Products cart:")
	for _, item := range items {
		fmt.Println(item)
	}
}

func operation(values ...int) (mean float64, median float64) {
	var total int
	for _, value := range values {
		total += value
	}

	mean = float64(total) / float64(len(values))

	if len(values)%2 == 0 {
		median = float64(len(values)/2) + 0.5
	} else {
		median = float64(int(len(values)/2) + 1)
	}

	return mean, median
}
