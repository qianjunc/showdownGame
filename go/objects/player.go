package objects

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// type Player struct {
// 	Name           string
// 	Hand           *Hand
// 	Points         int
// 	ExchangeRounds *Exchange
// 	Type           string
// }

// func NewPlayer() *Player {
// 	playerType := inputFromTerminal("What type of player is this? AI or Human")
// 	player := Player{}
// 	player.Points = 0
// 	hand := NewHand()
// 	player.Hand = &hand
// 	player.Type = playerType
// 	return &player
// }

// func (player *Player) GetHand() *Hand {
// 	return player.Hand
// }

// func (player *Player) SetHand(hand *Hand) {
// 	player.Hand = hand
// }

// func (player *Player) GainPoint() {
// 	player.Points += 1
// }

// func (player *Player) NameMyself() {
// 	name := ""
// 	if player.Type == "AI" {
// 		rand.Seed(time.Now().UnixNano())

// 		nameSlice := make([]byte, 5)
// 		for i := range name {
// 			nameSlice[i] = letterBytes[rand.Intn(len(letterBytes))]
// 		}
// 		name = string(nameSlice)
// 	} else {
// 		name = inputFromTerminal("Name yourself")
// 	}
// 	player.Name = name
// }

// // TODO: fix this bug
// func (player *Player) ExchangeOrNot() bool {
// 	if player.ExchangeRounds != nil {
// 		if player.ExchangeRounds.ExchangeRound >= 3 {
// 			player.exchangeHand(player.ExchangeRounds.Player2)
// 		}

// 		if player.ExchangeRounds.ExchangeRound < 3 {
// 			player.ExchangeRounds.IncreaseRound()
// 		}
// 		return false
// 	}

// 	decision := ""
// 	if player.Type == "AI" {
// 		choices := []string{"show", "exchange"}
// 		rand.Seed(time.Now().UnixNano())
// 		decision = choices[rand.Intn(len(choices))]
// 	} else {
// 		decision = inputFromTerminal("Do you exchange this round")
// 	}

// 	if decision == "yes" {
// 		return true
// 	}
// 	return false
// }

// func (player *Player) Show() Card {
// 	if player.Type == "AI" {
// 		i := rand.Intn(len(player.Hand.Cards))
// 		return player.Hand.Show(i)
// 	}
// 	decision := inputFromTerminal(fmt.Sprintf("choose index from: %v", player.Hand.Cards))
// 	i, err := strconv.Atoi(decision)
// 	if err != nil {
// 		return player.Hand.Show(0)
// 	}
// 	return player.Hand.Show(i)
// }

// func (player *Player) Exchange(anotherPlayer *Player) *Exchange {
// 	player.exchangeHand(anotherPlayer)
// 	newExchange := NewExchange(player, anotherPlayer)
// 	player.ExchangeRounds = newExchange
// 	return newExchange
// }

// func (player *Player) exchangeHand(anotherPlayer *Player) {
// 	currentHand := player.Hand
// 	player.Hand = anotherPlayer.GetHand()
// 	anotherPlayer.SetHand(currentHand)
// }

// // helper function
// func inputFromTerminal(question string) string {
// 	fmt.Printf("%s: ", question)
// 	reader := bufio.NewReader(os.Stdin)
// 	// ReadString will block until the delimiter is entered
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("An error occured while reading input. Please try again", err)
// 		return ""
// 	}

// 	// remove the delimeter from the string
// 	input = strings.TrimSuffix(input, "\n")
// 	return input
// }

// There are two ways of implementing inheritance in go: interface or embeded struct
// I choose interface because in this case, more functions are different
type Player interface {
	GetHand() *Hand
	SetHand(*Hand)
	Draw(*Deck)
	GetPoints() int
	GetName() string
	GainPoint()
	NameMyself()
	Decide() string
	Show() Card
	ExchangeOrNot() bool
	Exchange(anotherPlayer Player) *Exchange
}

func NewPlayer() Player {
	fmt.Printf("What type of player is this? AI or human: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	if input == "human" {
		NewHumanPlayer()
	}

	return NewAIPlayer()
}

type HumanPlayer struct {
	Player
	Name           string
	Hand           *Hand
	Points         int
	ExchangeRounds *Exchange
}

func NewHumanPlayer() *HumanPlayer {
	player := HumanPlayer{}
	player.Points = 0
	hand := NewHand()
	player.Hand = &hand
	return &player
}

func (player *HumanPlayer) GetHand() *Hand {
	return player.Hand
}

func (player *HumanPlayer) SetHand(hand *Hand) {
	player.Hand = hand
}

func (player *HumanPlayer) GetPoints() int {
	return player.Points
}

func (player *HumanPlayer) GetName() string {
	return player.Name
}

func (player *HumanPlayer) Draw(deck *Deck) {
	player.Hand.Draw(deck)
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

func (player *HumanPlayer) ExchangeOrNot() bool {
	if player.ExchangeRounds != nil {
		if player.ExchangeRounds.ExchangeRound >= 3 {
			player.Exchange(player.ExchangeRounds.Player2)
		}

		if player.ExchangeRounds.ExchangeRound < 3 {
			player.ExchangeRounds.IncreaseRound()
		}
		return false
	}
	decision := player.inputFromTerminal("Do you exchange this round")

	if decision == "yes" {
		return true
	}
	return false
}

func (player *HumanPlayer) Show() Card {
	decision := player.inputFromTerminal(fmt.Sprintf("choose index from: %v", player.Hand.Cards))
	i, err := strconv.Atoi(decision)
	if err != nil {
		return player.Hand.Show(0)
	}
	return player.Hand.Show(i)
}

// TODO
func (player *HumanPlayer) Exchange(anotherPlayer Player) *Exchange {
	currentHand := player.Hand
	player.Hand = anotherPlayer.GetHand()
	anotherPlayer.SetHand(currentHand)
	newExchange := NewExchange(player, anotherPlayer)
	player.ExchangeRounds = newExchange
	return newExchange
}

type AIPlayer struct {
	Player
	Name           string
	Hand           *Hand
	Points         int
	ExchangeRounds *Exchange
}

func NewAIPlayer() *AIPlayer {
	player := AIPlayer{}
	player.Points = 0
	hand := NewHand()
	player.Hand = &hand
	return &player
}

func (player *AIPlayer) GetHand() *Hand {
	return player.Hand
}

func (player *AIPlayer) SetHand(hand *Hand) {
	player.Hand = hand
}

func (player *AIPlayer) GetPoints() int {
	return player.Points
}

func (player *AIPlayer) GetName() string {
	return player.Name
}

func (player *AIPlayer) Draw(deck *Deck) {
	player.Hand.Draw(deck)
}

func (player *AIPlayer) GainPoint() {
	player.Points += 1
}

func (player *AIPlayer) NameMyself() {
	rand.Seed(time.Now().UnixNano())

	name := make([]byte, 5)
	for i := range name {
		name[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	player.Name = string(name)
}

func (player *AIPlayer) ExchangeOrNot() bool {
	if player.ExchangeRounds != nil {
		if player.ExchangeRounds.ExchangeRound >= 3 {
			player.Exchange(player.ExchangeRounds.Player2)
		}

		if player.ExchangeRounds.ExchangeRound < 3 {
			player.ExchangeRounds.IncreaseRound()
		}
		return false
	}

	choices := []string{"show", "exchange"}
	rand.Seed(time.Now().UnixNano())
	decision := choices[rand.Intn(len(choices))]

	if decision == "yes" {
		return true
	}
	return false
}

func (player *AIPlayer) Show() Card {
	i := rand.Intn(len(player.Hand.Cards))
	return player.Hand.Show(i)
}

func (player *AIPlayer) Exchange(anotherPlayer Player) *Exchange {
	currentHand := player.Hand
	player.Hand = anotherPlayer.GetHand()
	anotherPlayer.SetHand(currentHand)
	newExchange := NewExchange(player, anotherPlayer)
	player.ExchangeRounds = newExchange
	return newExchange
}
