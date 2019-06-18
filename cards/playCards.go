package cards

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type deckFace struct {
	front string
	back  string
}

func createCards(cardsToPractice [][]string) (cards []deckFace) {

	cards = make([]deckFace, len(cardsToPractice))

	for i, card := range cardsToPractice {

		cards[i] = deckFace{
			front: card[0],
			back:  card[1],
		}
	}

	return
}

func randomNumber(number int) int {

	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(number)

}

func cardInCardsAnswered(card deckFace, cardsAnswered []deckFace) bool {

	for _, cardAnswered := range cardsAnswered {

		if card == cardAnswered {
			return true
		}
	}

	return false

}

func Practice(cardsToPractice [][]string) {

	cards := createCards(cardsToPractice)

	var (
		card            deckFace
		cardsAnswered []deckFace
	)

	for totalOfCards := len(cards); totalOfCards > len(cardsAnswered); {

		card = cards[randomNumber(len(cardsToPractice))]

		if !cardInCardsAnswered(card, cardsAnswered) {

			fmt.Printf("%s = ", card.front)

			var answer string
			fmt.Scan(&answer)

			if answer != card.back {
				fmt.Println("Wrong answer. Do not give up! Try again!")
				os.Exit(1)
			} else {
				cardsAnswered = append(cardsAnswered, card)
			}
		}
	}

	fmt.Println("You got every thing!")
	os.Exit(1)
}
