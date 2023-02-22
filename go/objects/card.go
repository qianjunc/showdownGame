package objects

type Card struct {
	Rank Rank
	Suit Suit
}

func NewCard(rank RankValue, suit SuitValue) Card {
	card := Card{}
	card.Rank = NewRank(rank)
	card.Suit = NewSuit(suit)
	return card
}

func (card *Card) Win(anotherCard Card) bool {
	rankResult := card.Rank.Compare(anotherCard.Rank)
	if rankResult == "win" {
		return true
	}
	if rankResult == "lose" {
		return false
	}

	if card.Suit.Compare(anotherCard.Suit) == "win" {
		return true
	}
	return false
}
