package main

import (
	"flag"
	"fmt"
	"goCards/cards"
)

func main() {

	playCSV := flag.String("play", "", "Path for CSV file to start practicing ")
	lookCSV := flag.String("look", "", "Path for CSV file and see how the cards are organized")

	flag.Parse()

	if *playCSV != "" {

		openedCSV := cards.ReadCSV(*playCSV)
		cards.Practice(openedCSV)

	} else if *lookCSV != "" {

		for i := range cards.ReadCSV(*lookCSV) {
			fmt.Println(cards.ReadCSV(*lookCSV)[i])
		}


	} else {

		fmt.Println("Pass a valid path for a CSV file!")

	}

}
