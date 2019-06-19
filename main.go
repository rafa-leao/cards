package main

import (
	"flag"
	"fmt"
	"goCards/cards"
	"os"
)

var (
	learnCSV = flag.String("learn", "", "Path for CSV file to learn your cards ")
	playCSV  = flag.String("play", "", "Path for CSV file to start practicing ")
	lookCSV  = flag.String("look", "", "Path for CSV file and see how the cards are organized ")
)

func main() {

	flag.Parse()

	if *learnCSV != ""{

		openedCSV := cards.ReadCSV(*learnCSV)
		cards.Learn(openedCSV)

	} else if *playCSV != "" {

		openedCSV := cards.ReadCSV(*playCSV)
		cards.Practice(openedCSV)

	} else if *lookCSV != "" {

		for i := range cards.ReadCSV(*lookCSV) {

			fmt.Println(cards.ReadCSV(*lookCSV)[i])

		}

	} else {
		flag.Usage()
		os.Exit(1)
	}
}
