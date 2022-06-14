package main

import (
	"fmt"
	"time"
)

func print(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%v : %v \n", s, i)
	}
}

func main() {
	// print("hello")
	// print("world")
	// go print("hello")
	// go print("world")

	// go func (s string)  {
	// 	for x := 0; x < 5; x++ {
	// 		time.Sleep(100 * time.Millisecond)
	// 		fmt.Printf("%v : %v \n", s, x)
	// 	}
	// }("Amar")

	// bad practise
	// for x := 0; x < 5; x++ {
	// 	go func ()  {
	// 		fmt.Println(x)
	// 	}()
	// }

	// better way
	for i := 0; i < 5; i++ {
		go func(x int) {
			time.Sleep(time.Second * 3)
			fmt.Println(x)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
