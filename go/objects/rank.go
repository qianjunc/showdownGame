package objects

type RankValue string

const (
	Diamond RankValue = "Diamond"
	Club    RankValue = "Club"
	Spade   RankValue = "Spade"
	Heart   RankValue = "Heart"
)

type Rank struct {
	Value RankValue
}
