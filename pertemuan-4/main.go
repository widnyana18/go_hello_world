package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var dimen2 dimensi2 = &kerucut{10, 20, 30}

	fmt.Println("kerucut =========")
	fmt.Println("keliling : ", dimen2.keliling())

	var kerucut = dimen2.(*kerucut)
	kerucut.setUkuranKerucut(25, 10, 12)
	fmt.Printf("selimut : %f  |  jari: %f  |  tinggi : %f\n", kerucut.selimut, kerucut.jari, kerucut.tinggi)
	fmt.Println("new keliling : ", dimen2.keliling())

	var dimen3 hitung = &bola{20, 3}
	var bola = dimen3.(*bola)
	bola.setUkuranBola(8, 17)
	fmt.Println("bola =========")
	fmt.Printf("jari: %f  |  tinggi : %f\n", bola.jari, bola.tinggi)
	fmt.Println("luas : ", dimen3.luas())
	fmt.Println("volume : ", dimen3.volume())

	var name = "kelly"
	var meta = map[string]interface{}{
		"name":  "Sonroryu",
		"skill": 3,
		"rage":  true,
		"rate":  5.78,
	}

	checkDataType(name)
	inspectData(meta)

	// method := reflect.ValueOf(bola)
	// calling := method.MethodByName("calling")
	// calling.Call(
	// 	[]reflect.Value{reflect.ValueOf(6)},
	// )

	// println("new jari: " )
	var core = runtime.NumCPU()
	runtime.GOMAXPROCS(core - 1)

	var as = sync.WaitGroup{}
	as.Add(2)

	go calculate(8)
	as.Wait()
	calculate(5)

	as.Done()
}

// func (b bola) calling(jari float64) {
// 	b.jari = jari
// }

func calculate(angka int) {
	var data = 10 % angka
	println("data: ", data)
}
