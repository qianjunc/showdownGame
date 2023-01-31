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
	suitResult := card.Suit.Compare(anotherCard.Suit)
	if suitResult == "win" {
		return true
	}
	if suitResult == "lose" {
		return false
	}

	if card.Rank.Compare(anotherCard.Rank) == "win" {
		return true
	}
	return false
}
