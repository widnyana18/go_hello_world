package controller

import (
	"database/sql"
	repo "insert_million/repository"
	"sync"
)

const maxWorker = 100

func DispacthWorker(db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for worker := 0; worker < maxWorker; worker++ {
		go func(workerIdx int, db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			workerIdx++
			var counter = 0

			for job := range jobs {
				repo.InsertDataToDB(workerIdx, counter, db, job)
				wg.Done()
				counter++
			}
		}(worker, db, jobs, wg)
	}
}
