package controller

import (
	"encoding/csv"
	"io"
	"sync"
)

var DataHeaders = make([]string, 0)

func TransportDataToWorker(reader *csv.Reader, jobs chan<- []interface{}, wg *sync.WaitGroup) {
	for {
		row, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if len(DataHeaders) == 0 {
			DataHeaders = row
			continue
		}

		dataRows := make([]interface{}, 0)
		for _, each := range row {
			dataRows = append(dataRows, each)
		}

		wg.Add(1)
		jobs <- dataRows
	}

	close(jobs)
}
