package slices

import (
	"fmt"
)

func main() {
	src := []string{"Ram", "x", "y", "z"}
	fmt.Println(src)
	src = append(src, "d")

	dest := make([]string, 2)
	copy(dest, src)

	src[0] = "Shyam"
	fmt.Println(src)
	fmt.Println(dest)

	// since, Slice is mainly a pointer to underlying Array,
	// slicing a slice does not create a copy of values.
	// this will create a slice that point to same values in the original Array.
	s := src[1:2]
	fmt.Println(s)

	// updating the source array, the sliced array also gets updated,
	// since the sliced array points to the same original array address.
	src[2] = "updated"
	fmt.Println(src)
	fmt.Println(s)

	gameMap := [][]string{
		make([]string, 3),
		make([]string, 5),
		[]string{"-", "-"},
		//{"-", "-"} also possible
	}
	fmt.Println(gameMap)
	gameMap[0] = make([]string, 6)
	gameMap[0] = make([]string, 10)
	fmt.Println(gameMap)
	gameMap[1][2] = "-"
	fmt.Println(gameMap)
}
