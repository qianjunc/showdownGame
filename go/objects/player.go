package objects

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// There are two ways of implementing inheritance in go: interface or embeded struct
// I choose interface because in this case, more functions are different
type Player interface {
	GetHand() *Hand
	SetHand(*Hand)
	GainPoint()
	NameMyself()
	Decide()
	Show()
	Exchange(anotherPlayer Player) Exchange
}

type HumanPlayer struct {
	Player
	Name   string
	Hand   *Hand
	Points int
}

func NewHumanPlayer() HumanPlayer {
	player := HumanPlayer{}
	player.Points = 0
	hand := NewHand()
	player.Hand = &hand
	return player
}

func (player *HumanPlayer) GetHand() *Hand {
	return player.Hand
}

func (player *HumanPlayer) SetHand(hand *Hand) {
	player.Hand = hand
}

func (player *HumanPlayer) GainPoint() {
	player.Points += 1
}

func (player *HumanPlayer) NameMyself() {
	name := player.inputFromTerminal("Name yourself")
	player.Name = name
}

func (player *HumanPlayer) inputFromTerminal(question string) string {
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

func (player *HumanPlayer) Decide() {

}

func (player *HumanPlayer) Show() {

}

func (player *HumanPlayer) Exchange(anotherPlayer Player) Exchange {
	currentHand := player.Hand
	player.Hand = anotherPlayer.GetHand()
	anotherPlayer.SetHand(currentHand)
	return NewExchange(player, anotherPlayer)
}

type AIPlayer struct {
	Player
	Name   string
	Hand   *Hand
	Points int
}

func NewAIPlayer() AIPlayer {
	player := AIPlayer{}
	player.Points = 0
	hand := NewHand()
	player.Hand = &hand
	return player
}

func (player *AIPlayer) GainPoint() {
	player.Points += 1
}

func (player *AIPlayer) NameMyself() {

}

func (player *AIPlayer) Decide() {

}

func (player *AIPlayer) Show() {

}

func (player *AIPlayer) Exchange(anotherPlayer Player) Exchange {
	currentHand := player.Hand
	player.Hand = anotherPlayer.GetHand()
	anotherPlayer.SetHand(currentHand)
	return NewExchange(player, anotherPlayer)
}
