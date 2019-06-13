package cards

import "fmt"

func Practice(cardsToPractice [][]string) {

	type deckFace struct {
		front string
		back  string
	}

	for _, card:= range cardsToPractice {

		card := deckFace {
			front: card[0],
			back:  card[1],
		}

		fmt.Printf("%s = ", card.front)

		var answer string
		fmt.Scanf("%s", &answer)

		if answer != card.back{
			fmt.Println("Wrong answer. Do not give up! Try again!")
			break
		}
	}
}
