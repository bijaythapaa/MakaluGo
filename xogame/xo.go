package main

import "fmt"

func main() {
	// init the game board with empty strings
	xoBoard := [3][3]string{}

	// variable to carry the current player name
	player := "x"

	// in this infinite-loop if a user enters an invalid input, the loop will continue.
	// that's the logic
	for {
		fmt.Println("player", player)

		// read row value
		var row int
		fmt.Println("please enter number 0, 1 or 2")
		fmt.Scanln(&row)
		if row < 0 || row > 2 {
			fmt.Println("invalid input, please enter 0, 1 or 2")
			continue
		}

		// read column value
		var column int
		fmt.Println("please enter number 0, 1 or 2")
		fmt.Scanln(&column)
		if column < 0 || column > 2 {
			fmt.Println("invalid input, please enter 0, 1 or 2")
			continue
		}

		// set value into game board
		if xoBoard[row][column] == "" {
			xoBoard[row][column] = player
		} else {
			// index already has value
			fmt.Println("invalid entry: ", row, column, "value", xoBoard[row][column])
			continue
		}

		//display game board after each move
		fmt.Println(xoBoard[0])
		fmt.Println(xoBoard[1])
		fmt.Println(xoBoard[2])

		// did someone win
		for i := 0; i < 3; i++ {
			// check row || column
			if xoBoard[i][0] == xoBoard[i][1] && xoBoard[i][1] == xoBoard[i][2] && xoBoard[i][2] == player ||
				xoBoard[0][i] == xoBoard[1][i] && xoBoard[1][i] == xoBoard[2][i] && xoBoard[2][i] == player {
				fmt.Println("game ended, winner is player: ", player)
				return
			}
		}
		if xoBoard[0][0] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][2] && xoBoard[2][2] == player ||
			xoBoard[0][2] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][0] && xoBoard[2][0] == player {
			fmt.Println("game ended, the winner is player: ", player)
			return
		}

		// swap players
		if player == "x" {
			player = "o"
		} else {
			player = "x"
		}

	}

}
