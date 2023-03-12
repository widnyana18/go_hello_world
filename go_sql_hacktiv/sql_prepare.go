package main

import "fmt"

func sqlPrepare() {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT name, age, grade FROM student where id = ?")

	if err != nil {
		fmt.Println(err)
		return
	}

	var student1 = student{}
	stmt.QueryRow("BB02").Scan(&student1.name, &student1.age, &student1.grade)
	fmt.Println("name1: ", student1.name)

	var student2 = student{}
	stmt.QueryRow("BB03").Scan(&student2.name, &student2.age, &student2.grade)
	fmt.Println("name2: ", student2.name)

	var student3 = student{}
	stmt.QueryRow("BB04").Scan(&student3.name, &student3.age, &student3.grade)
	fmt.Println("name3: ", student3.name)
}
