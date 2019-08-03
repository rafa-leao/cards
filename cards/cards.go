package cards

type Card struct {
	frontFace string
	backFace  string
}

func CreateSliceOfCards(csvFile [][]string) (cardsToPractice []Card) {

	cardsToPractice = make([]Card, len(csvFile))

	for i, card := range csvFile {

		cardsToPractice[i] = Card {
			frontFace: card[0],
			backFace:  card[1],
		}
	}

	return
}

func IsCardInSliceOfCards(card Card, cardsAnswered []Card) bool {

	for _, cardAnswered := range cardsAnswered {

		if card == cardAnswered {
			return true
		}
	}

	return false

}
