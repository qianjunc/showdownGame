package objects

type SuitValue string

const (
	Two   SuitValue = "2"
	Three SuitValue = "3"
	Four  SuitValue = "4"
	Five  SuitValue = "5"
	Six   SuitValue = "6"
	Seven SuitValue = "7"
	Eight SuitValue = "8"
	Nine  SuitValue = "9"
	Ten   SuitValue = "10"
	J     SuitValue = "J"
	Q     SuitValue = "Q"
	K     SuitValue = "K"
	A     SuitValue = "A"
)

type Suit struct {
	Value SuitValue
}
