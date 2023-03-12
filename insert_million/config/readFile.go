package config

import (
	"encoding/csv"
	"log"
	"os"
)

func OpenCsvFile() (*csv.Reader, *os.File, error) {
	log.Println("Start reading CSV")

	file, err := os.Open("majestic_million.csv")

	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(file)
	return reader, file, nil
}
