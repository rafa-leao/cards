package main

import (
	"fmt"
	"os"
)

func Learn(csvFile [][]string) {

	var (
		cardsToLearn    []Card = CreateSliceOfCards(csvFile)
		cardsLearned    []Card
		cardsNotLearned []Card
		card            Card
	)

	if len(cardsToLearn) < 5 {
		fmt.Println("Please, your CSV file must have at least 5 cards")
		os.Exit(1)
	}

	for len(cardsToLearn) > len(cardsLearned) {

		for i := range cardsToLearn {

			if card = cardsToLearn[i]; !IsCardInSliceOfCards(card, cardsLearned) {

				fmt.Printf("%s = ", card.frontFace)

				var answer string
				fmt.Scan(&answer)

				if answer == card.backFace {
					cardsLearned = append(cardsLearned, card)
				} else {
					cardsNotLearned = append(cardsNotLearned, card)
				}
			}
		}

		if cardsNotLearned != nil {

			fmt.Println("You still don't know some cards")
			fmt.Println("Let's try those cards again!")

			for i := range cardsNotLearned {

				card = cardsNotLearned[i]

				for IsCardInSliceOfCards(card, cardsNotLearned) && !IsCardInSliceOfCards(card, cardsLearned) {

					fmt.Printf("%s = ", card.frontFace)

					var answer string
					fmt.Scan(&answer)

					if answer == card.backFace {
						cardsLearned = append(cardsLearned, card)
					} else {
						fmt.Println("It's wrong, try again")
					}
				}
			}
		}
	}

	fmt.Println("You learned every cards in your deck!")
	os.Exit(1)
}
