package objects

type Hand struct {
	Cards []Card
}

func NewHand() Hand {
	hand := Hand{}
	hand.Cards = []Card{}
	return hand
}

func (hand *Hand) Draw(deck Deck) {
	card := deck.Draw()
	hand.Cards = append(hand.Cards, card)
}
