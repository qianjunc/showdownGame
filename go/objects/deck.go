package objects

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	deck := Deck{}
	suits := []SuitValue{Spade, Heart, Club, Diamond}
	ranks := []RankValue{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, J, Q, K, A}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.Cards = append(deck.Cards, NewCard(rank, suit))
		}
	}
	return &deck
}

func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) { deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i] })
}

func (deck *Deck) Draw() Card {
	card, cards := deck.Cards[0], deck.Cards[1:]
	deck.Cards = cards
	return card
}
