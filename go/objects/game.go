package objects

type Game struct {
	Deck    *Deck
	Players []*Player
}

func NewGame() *Game {
	game := Game{}
	game.Deck = NewDeck()
	for i := 0; i < 4; i++ {
		player := NewPlayer()
		game.Players = append(game.Players, player)
	}

	return &game
}

func (game *Game) Start() {
	// shuffle deck

	// draw card

	// 13 rounds

	// winner
}
