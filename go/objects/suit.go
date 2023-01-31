package objects

type SuitValue string

const (
	Club    SuitValue = "Club"
	Diamond SuitValue = "Diamond"
	Heart   SuitValue = "Heart"
	Spade   SuitValue = "Spade"
)

type Suit struct {
	Value SuitValue
}

func NewSuit(value SuitValue) Suit {
	suit := Suit{}
	suit.Value = value
	return suit
}

func (suit *Suit) Compare(anotherSuit Suit) string {
	if anotherSuit.Value == suit.Value {
		return "draw"
	}
	if suit.Value == Club || (suit.Value == Diamond && anotherSuit.Value != Club) ||
		(suit.Value == Heart && anotherSuit.Value == Spade) {
		return "win"
	}
	return "lose"
}
