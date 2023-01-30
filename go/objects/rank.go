package objects

type RankValue string

const (
	Two   RankValue = "2"
	Three RankValue = "3"
	Four  RankValue = "4"
	Five  RankValue = "5"
	Six   RankValue = "6"
	Seven RankValue = "7"
	Eight RankValue = "8"
	Nine  RankValue = "9"
	Ten   RankValue = "10"
	J     RankValue = "J"
	Q     RankValue = "Q"
	K     RankValue = "K"
	A     RankValue = "A"
)

type Rank struct {
	Value RankValue
}

func NewRank(value RankValue) Rank {
	rank := Rank{}
	rank.Value = value
	return rank
}

func (rank *Rank) Compare(anotherRank *Rank) string {
	if rank.Value == anotherRank.Value {
		return "draw"
	}

	if rank.Value == A || (rank.Value == K && anotherRank.Value != A) ||
		(rank.Value == Q && anotherRank.Value != A && anotherRank.Value != K) ||
		(rank.Value == J && anotherRank.Value != A && anotherRank.Value != K && anotherRank.Value != Q) ||
		(rank.Value == Ten && anotherRank.Value != J && anotherRank.Value != A && anotherRank.Value != K && anotherRank.Value != Q) ||
		(rank.Value > anotherRank.Value) {
		return "win"
	}
	return "lose"
}
