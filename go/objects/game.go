package objects

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Deck    *Deck
	Players []Player
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
	// Name themselves
	for _, player := range game.Players {
		player.NameMyself()
	}

	// shuffle deck
	game.Deck.Shuffle()

	// draw card
	for i := 0; i < 13; i++ {
		for _, player := range game.Players {
			player.Draw(game.Deck)
		}
	}

	// 13 rounds
	for i := 0; i < 13; i++ {
		roundWinner := game.findOneRoundWinner()
		roundWinner.GainPoint()
	}

	// winner
	winner := 0
	for i := 1; i < 4; i++ {
		if game.Players[i].GetPoints() > game.Players[winner].GetPoints() {
			winner = i
		}
	}

	fmt.Printf("The winner is p%d %s\n", winner+1, game.Players[winner].GetName())
}

func (game *Game) findOneRoundWinner() Player {
	roundCards := []Card{}
	for i, player := range game.Players {
		fmt.Printf("P%d\n", i+1)
		if player.ExchangeOrNot() {
			exchangePlayer := inputFromTerminal("Which player do you want to exchange with")
			ei, _ := strconv.Atoi(exchangePlayer)
			player.Exchange(game.Players[ei-1])
		}
		showCard := player.Show()
		roundCards = append(roundCards, showCard)
	}
	fmt.Printf("This round's card: %v\n", roundCards)

	winner := 0
	for i := 1; i < 4; i++ {
		if roundCards[i].Win(roundCards[winner]) {
			winner = i
		}
	}
	return game.Players[winner]
}

func inputFromTerminal(question string) string {
	fmt.Printf("%s: ", question)
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return ""
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	return input
}
