package main

import "fmt"

// file scope, accessible from everywhere
// var name = "Bijay"

func main() {
	// fmt.Println(name)

	// inside any function we can group some codes using {}
	a := 23
	{
		fmt.Println("a In", a)
		{
			fmt.Println("a Inner", a)
		}
	}
	fmt.Println("a out", a)

	x, y := 1, 2
	fmt.Println("variables : ", x, y)

	g := Game{
		started: false,
		winner: Player{
			name:  "Amar",
			score: 77,
		},
	}
	fmt.Println("G: ", g)

	// t := PrivateType{} 

}

type Game struct {
	started bool
	winner  Player
}

type Player struct {
	name  string
	score int
}

func t() {
	type PrivateType struct {
	}
	t := PrivateType{}
	_ = t
}
