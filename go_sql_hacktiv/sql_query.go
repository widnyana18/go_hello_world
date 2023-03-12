package main

import "fmt"

func sqlQuery() {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	var age = 22
	rows, err := db.Query("SELECT name, age, grade FROM student where age = ?", age)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		err := rows.Scan(&each.name, &each.age, &each.grade)

		if err != nil {
			fmt.Println(err)
			return
		}

		result = append(result, each)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}
