package game

import (
	"math/rand"
	"strconv"
)

// Cards architecture

type Card struct {
	color string
	number int
}

func newCard(number int, color string) Card {
	card := Card{color, number}
	return card
}

func (card *Card) printCard () string {
	var value string;
	switch card.number {
		case 11:
			value = "J"
			break
		case 12:
			value = "Q"
			break
		case 13:
			value = "K"
			break
		case 14:
			value = "A"
			break
		default:
			value = strconv.Itoa(card.number)
	}
	switch card.color {
		case "pik":
			value += "\u2664"
			break
		case "kier":
			value += "\u2661"
			break
		case "karo":
			value += "\u2662"
			break
		case "trefl":
			value += "\u2663"
			break
		default:
			panic("Incorrect card color!")
	}
	//fmt.Print(value, " ")
	return value
}

// Deck architecture

type Deck struct {
	cards []Card
}

func newDeck() Deck {
	var deck Deck
	for i := 2; i <= 14; i++ {
		for _, c := range []string{"trefl", "karo", "kier", "pik"} {
			deck.cards = append(deck.cards, newCard(i, c))
		}
	}
	deck.shuffle()
	return deck
}

func (deck *Deck) draw() Card {
	drawn := deck.cards[len(deck.cards) - 1]
	deck.cards = deck.cards[:len(deck.cards) - 1]
	return drawn
}

func (deck *Deck) shuffle() {
	rand.Shuffle(len(deck.cards), func(i, j int) {deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]})
}