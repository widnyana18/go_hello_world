package main

import "fmt"

func array() {
	var name, email string
	name = "kkak"
	email = "en"

	const (
		order   = iota
		address = "ponorogo"
	)

	fmt.Println(name, email, order, address)

	var os = [...]string{"mac", "windows", "linux", "unix"}

	fmt.Println("operating system ", os[0])

	var car = []string{"mitsubitsi", "Lamborgini", "Mercendess", "Niisan"}

	fmt.Println(car)

	var price = make([]int, 10)
	price[2] = 3000
	price[5] = 6000
	price[7] = 1000

	fmt.Println(price)

	date := price[:4]
	fmt.Println("date: ", date, "cap: ", cap(date))

	var newDate = append(date, 20)
	fmt.Printf("price: %d %d\n", price, newDate)

	var country = []string{"USA", "Indonesia", "Japan", "Jerman", "Australia"}

	var asia = make([]string, 8)
	asia[6] = "Maroco"
	asia[7] = "China"

	copy(asia, country)

	asia[0] = "North Korea"

	fmt.Println("europe: ", country, "asia: ", asia)
	fmt.Printf("Select the country you wont! ")
	fmt.Scanf("%s")
}
