package student

import "fmt"

type ListStudent []Student

func SelectMember(list ListStudent, dataIdx int) {
	var student Student

	for i, data := range list {
		if i == (dataIdx - 1) {
			student = data
		}
	}

	student.Job = "Mahasiswa"
	student.showData()
}

func (student *Student) showData() {
	println()
	fmt.Printf("==================================== Student Data ==================================\n\n")
	fmt.Println("Name     : ", student.Name)
	fmt.Println("Address  : ", student.Address)
	fmt.Println("Job 	 : ", student.Job)
	fmt.Println("Comment  : ", student.Comment)
	fmt.Println("===================================================================================")
	println()
}
