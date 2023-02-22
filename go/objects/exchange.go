package objects

type Exchange struct {
	ExchangeRound int
	Player1       Player
	Player2       Player
}

func NewExchange(p1, p2 Player) *Exchange {
	exchange := Exchange{}
	exchange.ExchangeRound = 1
	exchange.Player1 = p1
	exchange.Player2 = p2
	return &exchange
}

func (ex *Exchange) IncreaseRound() {
	ex.ExchangeRound += 1
	// if ex.ExchangeRound == 3 {
	// 	ex.Player1.Exchange(ex.Player2)
	// }
}
