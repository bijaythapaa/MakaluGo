package main

import "fmt"

func main() {
	fmt.Println("Welcome to the LRC Game !!")
	g := new()
	fmt.Println("How many players will play the Game?")
	fmt.Println("Note: at least 3 players")

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
