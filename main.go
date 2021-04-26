package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)
}
