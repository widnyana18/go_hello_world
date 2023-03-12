package main

import "fmt"

func sqlExec() {
	db, err := connectDB()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO student VALUES (?,?,?,?)")

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = stmt.Exec("BB05", "mikey aley", 22, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("insert data success!")

	_, err = db.Exec("UPDATE student SET age = ? WHERE id = ? ", 33, "BB02")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("update data success!")

	_, err = db.Exec("DELETE FROM student WHERE id = ?", "BB04")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("delete data success!")
}
