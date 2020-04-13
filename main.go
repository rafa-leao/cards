package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	learnCSV = flag.String("learn", "", "Path for CSV file to learn your cards ")
	playCSV  = flag.String("play", "", "Path for CSV file to start practicing ")
	lookCSV  = flag.String("look", "", "Path for CSV file and see how the cards are organized ")
)

func main() {

	flag.Parse()

	switch {

	case *learnCSV != "":

		openedCSV := ReadCSV(*learnCSV)
		Learn(openedCSV)

	case *playCSV != "":

		openedCSV := ReadCSV(*playCSV)
		Practice(openedCSV)

	case *lookCSV != "":

		for i := range ReadCSV(*lookCSV) {

			fmt.Println(ReadCSV(*lookCSV)[i])
		}

	default:
		flag.Usage()
		os.Exit(1)
	}

}
