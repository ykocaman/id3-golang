package main

import (
	"encoding/csv"
	"log"
	"os"
)

func Load(path string) (dataset *datasetBase) {
	dataset = new(datasetBase)

	file, err := os.Open(path)

	// add after-work to close file session
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// create stream reader
	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	dataset.head = rows[0]
	dataset.rows = rows[1:]
	return
}
