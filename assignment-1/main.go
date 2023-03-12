package main

import (
	"fmt"
	st "menampilkan_data/student"
	"os"
	"strconv"
)

var students st.ListStudent

func init() {

	students = st.ListStudent{
		st.Student{
			Person: st.Person{
				Name:    "I Komang Widnyana",
				Address: "Badung, Bali",
			},
			Comment: "Agar menambahkan wawasan  dibidang fullstack dan golang",
		},
		st.Student{
			Person: st.Person{
				Name:    "Andri Nur Hidayatulloh",
				Address: "Yogyakarta",
			},
			Comment: "saya ikut course ini berharap ada product yang bisa saya buat dan membantu masyarakat",
		},
		st.Student{
			Person: st.Person{
				Name:    "Andri Kuwito",
				Address: "Bogor",
			},
			Comment: "saya ikut course ini ingin mempelajari hal baru",
		},
		st.Student{
			Person: st.Person{
				Name:    "Erico",
				Address: "Medan",
			},
			Comment: "saya ikut course ini ingin mempelajari hal baru",
		},
		st.Student{
			Person: st.Person{
				Name:    "Firman Aulia Insani",
				Address: "Cilacap",
			},
			Job:     "Mahasiswa",
			Comment: "dengan course ini saya berharap bisa mempelajari bidang baru, web khususnya backend karena selama ini hanya berkutat di android.",
		},
		st.Student{
			Person: st.Person{
				Name:    "Melvita Sari",
				Address: "Lhokseumawe",
			},
			Job:     "Mahasiswa",
			Comment: "Melalui course ini saya berharap bisa menambah ilmu dan memperbanyak teman. Berkeinginan menjadi seorang yang memiliki daya ingat yang kuat",
		},
		st.Student{
			Person: st.Person{
				Name:    "Jose Yolanda Purba",
				Address: "Medan",
			},
			Comment: "saya ikut course ini berharap ada product yang bisa saya buat dan membantu masyarakat",
		},
		st.Student{
			Person: st.Person{
				Name:    "Rafli Andreansyah",
				Address: "Blitar",
			},
			Comment: "saya ikut course ini berharap ada product yang bisa saya buat dan membantu masyarakat",
		},
		st.Student{
			Person: st.Person{
				Name:    "Taufiq Hidayah",
				Address: "Yogyakarta",
			},
			Comment: "saya ikut course ini berharap ada product yang bisa saya buat dan membantu masyarakat",
		},
		st.Student{
			Person: st.Person{
				Name:    "Sandy Budiman",
				Address: "Yogyakarta",
			},
			Job:     "Mahasiswa",
			Comment: "saya gabung course ini ingin belajar programming, khususnya golang karena, sedang hype dan high demand. Semoga bisa membuat product yg dapat bermanfaat bagi semua.",
		},
	}
}

func main() {
	var input = os.Args
	var (
		numInput int
		_        error
	)

	for _, num := range input {
		numInput, _ = strconv.Atoi(num)
	}
	if numInput > len(students) {
		fmt.Println("The participant number you are looking for, doesn't exist")
	} else if numInput == 0 {
		fmt.Println("Please enter the participant number you are looking for first!!!")
	} else {
		st.SelectMember(students, numInput)
	}
}
