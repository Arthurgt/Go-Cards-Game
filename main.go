package main

import "fmt"

func main() {
	cards := []string {"Six of diamonds", newCard()}
	cards = append(cards, "Three of glances")
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of diamonds"
}
