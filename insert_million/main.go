package main

import (
	"fmt"
	"insert_million/config"
	ctrl "insert_million/controller"
	"log"
	"math"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	start := time.Now()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	data, file, err := config.OpenCsvFile()
	defer file.Close()

	var jobs = make(chan []interface{})
	var wg = new(sync.WaitGroup)

	go ctrl.DispacthWorker(db, jobs, wg)
	ctrl.TransportDataToWorker(data, jobs, wg)

	wg.Wait()

	duration := time.Since(start)

	fmt.Println("Done in ", math.Ceil(duration.Seconds()), " Seconds")
}
