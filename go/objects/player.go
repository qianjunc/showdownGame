package objects

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Name   string
	Hand   *Hand
	Points int
}

func NewPlayer() Player {
	player := Player{}
	player.Points = 0
	hand := NewHand()
	player.Hand = &hand
	return player
}

func (player *Player) NameMyself() {
	name := player.inputFromTerminal("Name yourself")
	player.Name = name
}

func (player *Player) inputFromTerminal(question string) string {
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

func (player *Player) GainPoint() {
	player.Points += 1
}

func (player *Player) Decide() {

}

func (player *Player) Show() {

}

func (player *Player) Exchange() {

}
