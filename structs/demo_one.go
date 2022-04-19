// Structs are collection of fields

package main

import "fmt"

type Player struct {
	firstname string
	lastname  string
	score     int
	// any       bool
}

type Game struct {
	players map[string]*Player
}

func (g *Game) getWinner() *Player {
	var maxScore int
	var winner *Player

	for _, val := range g.players {
		if val.score > maxScore {
			winner = val
			maxScore = val.score
		}
	}
	return winner
}

func main() {

	g := &Game{
		players: make(map[string]*Player),
	}

	p1 := Player{}
	fmt.Println(p1)

	g.players["p1"] = &p1
	fmt.Println(g)

	var p2 Player
	p2 = Player{}
	p2.score = 11
	fmt.Println(p2)

	g.players["p2"] = &p2
	fmt.Println("Game: ", g)

	p3 := Player{
		lastname:  "Thapa",
		firstname: "Bijay",
		score:     45,
	}
	p3.score = 55
	fmt.Println(p3)

	g.players["p3"] = &p3
	fmt.Println("Game: ", g)

	fmt.Println("Game: ", g.players["p2"])
	fmt.Println("Game: ", *g.players["p2"])

	winner := g.getWinner()
	fmt.Println("*Winner: ", winner)
}
