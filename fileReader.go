package main

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(fileName string) [][]string {

	csvFile, err := os.Open(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	result, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	return result
}
