package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the LRC Game !!")
	g := new()
	fmt.Println("How many players will play the Game?")
	fmt.Println("Note: at least 3 players")

	// need players count and will take it as input
	var playersCount int
	for {
		fmt.Scanln(&playersCount)
		if playersCount < 3 {
			fmt.Println("Note: Enter number more than 2")
			continue
		}
		break
	}

	for i := 0; i < playersCount; i++ {
		playerName := fmt.Sprintf("p%v", i)
		P := g.join(playerName)
		fmt.Println(fmt.Sprintf("player:%v joined", P.name))
	}

	// turn
	turn := g.players[0]

}

type Game struct {
	players []*Player
}

// join add new player to the game
func (g *Game) join(playerName string) *Player {
	// create new player
	p := Player{
		name:   playerName,
		tokens: 3,
	}
	playerCount := len(g.players)
	if playerCount > 0 {
		lastPlayer := g.players[playerCount-1]
		p.left = lastPlayer
		p.right = g.players[0]
		lastPlayer.right = &p
		g.players[0].left = &p
	}
	g.players = append(g.players, &p)
	return
}

// finished check do we have only one player with Tokens
// this should be done after every turn
func (g *Game) finished() (p *Player) {
	playersWithTokens := 0
	for _, v := range g.players {
		if v.tokens > 0 {
			playersWithTokens++
			if playersWithTokens > 1 {
				return nil
			}
			p = v
		}
	}
	return
}

// new int game
func new() *Game {
	return &Game{}
}

// player
type Player struct {
	name   string
	tokens int
	right  *Player
	left   *Player
}

// roll uses rand to return dicefaces
func (p *Player) rollDice() (result []string) {
	// find out how many dices we should roll
	// if user has more than 3 tokens he can roll 3 dices
	// or he roll exact number of tokens as dices
	dices := p.tokens
	if p.tokens > 2 {
		dices = 3
	}

	//roll the dices and update the value on the player
	for index := 0; index < dices; index++ {
		d := Dice{}
		diceResults := d.roll()
		result = append(result, diceResults)
		switch diceResults {
		case "right":
			p.tokens--
			p.right.tokens++
		case "left":
			p.tokens--
			p.left.tokens++
		case "center":
			p.tokens--
		}
	}
}

// dice
type Dice struct {
}

func (d *Dice) roll() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(6)
	switch r {
	case 0:
		return "right"
	case 1:
		return "left"
	case 2:
		return "center"
	default:
		return "DoNothing"
	}
}
