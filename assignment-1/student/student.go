package student

type Person struct {
	Name    string
	Address string
}

type Student struct {
	Person
	Job     string
	Comment string
}
