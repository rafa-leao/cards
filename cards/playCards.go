package cards

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func randomNumber(number int) int {

	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(number)

}

func Practice(csvFile [][]string) {

	var (
		cardsToPractice []Card = CreateSliceOfCards(csvFile)
		cardsAnswered   []Card
		card              Card
	)

	for totalOfCards := len(cardsToPractice); totalOfCards > len(cardsAnswered); {

		card = cardsToPractice[randomNumber(len(cardsToPractice))]

		if !IsCardInSliceOfCards(card, cardsAnswered) {

			fmt.Printf("%s = ", card.frontFace)

			var answer string
			fmt.Scan(&answer)

			if answer != card.backFace {
				fmt.Printf("Wrong answer. Do not give up! You made %d of %d.\n", len(cardsAnswered), len(cardsToPractice))
				fmt.Println("Try again or learn your deck of cards with the '--learn' command!")
				os.Exit(1)
			} else {
				cardsAnswered = append(cardsAnswered, card)
			}
		}
	}

	fmt.Println("You got every thing!")
	os.Exit(1)
}
