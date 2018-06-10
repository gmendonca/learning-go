package main

func main() {
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {

	}

}

func newCard() string {
	return "Five of Diamonds"
}
