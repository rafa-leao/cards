package cards

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(fileName string) [][]string {

	csvFile, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("Something went wrong when trying to open your file", err)
	}

	result, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return result
}
