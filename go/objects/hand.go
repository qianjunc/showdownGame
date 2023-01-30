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

func (hand *Hand) Show(i int) Card {
	card := hand.Cards[i]
	hand.Cards = append(hand.Cards[:i], hand.Cards[i+1:]...)
	return card
}
