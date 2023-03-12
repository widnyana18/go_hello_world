package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Printf("============ FUNCTION ============\n\n")

	luas(6)
	luas(15)

	order("Georges", "shoes", "t-shirt", "hat", "belt")

	var age = []int{4, 3, 2, 5, 1, 20, 6, 21, 17, 14, 7, 25}
	var distance = make([]int, len(age), 15)

	copy(distance, age)

	var newDis = append(distance, 20, 33, 55)

	fmt.Printf("distance: %d | len: %d | cap: %d \n", newDis, len(newDis), cap(newDis))

	meanA, medianA := operation(age...)
	fmt.Printf("mean A : %f\nmedian A : %f\n", meanA, medianA)

	meanB, medianB := operation(newDis...)
	fmt.Printf("mean B : %1f\n", meanB)
	fmt.Printf("median B : %1f\n", medianB)

	fmt.Printf("\n\n\n")
	fmt.Printf("============== CLOSURE ============\n\n")

	var stock []int
	var animals = []string{"dog", "cat", "horse", "snake", "Giraffe", "Dragon", "bird"}
	var library = []string{"Galaxy", "andromeda", "galaxy", "stone war", "one piece", "Dragon ball", "naruto", "one piece"}

	var animalName string
	var bookName string

	for i := 0; i < 10; i++ {
		stock = append(stock, rand.Intn(20))
	}

	ascend, descend := sorting(stock)

	fmt.Println("Data stock : ", stock)
	fmt.Println("Data sort ascending : ", ascend)
	fmt.Println("Data sort descending : ", descend)
	fmt.Printf("\n")

	fmt.Println("data animals : ", animals)
	fmt.Printf("Pick your favorite animals : ")
	fmt.Scan(&animalName)

	var dataLen, animalClass = clasifier(animals, animalName)
	fmt.Printf("animals classification : %v\nData length : %d\n\n", animalClass(), dataLen)

	fmt.Println("library : ", library)
	fmt.Printf("Select book : ")
	fmt.Scan(&bookName)

	var searchBook = func(bookEntry string) bool {
		var tersedia bool
		if strings.Contains(bookEntry, bookName) {
			tersedia = true
		}
		return tersedia
	}

	selectedBook := selection(library, searchBook)

	fmt.Printf("Selected Book : %v\nBook Count : %s\n\n", selectedBook["name"], selectedBook["count"])

	fmt.Printf("\n\n\n")
	fmt.Printf("============== STRUCT ============\n\n")

	var media = media{author: author{name: "NF", age: 28, from: "England"}, size: 2, release: 2020}

	var music1 = music{title: "clouds", genre: "rap", duration: "3:25", media: media}

	var music2 *music = &music1
	music2.updateMusic("Lost")

	fmt.Println("title of music : ", music2.title)
}
