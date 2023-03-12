package main

import "fmt"

func sqlQueryRow() {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	var grade = 2
	var student = student{}

	row := db.QueryRow("SELECT name, age, grade FROM student where grade = ?", grade)
	err = row.Scan(&student.name, &student.age, &student.grade)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("name: ", student.name)
}
