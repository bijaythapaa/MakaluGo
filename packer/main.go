// Any go program is made of packages.
// Program starts running in package main.
// main package name is defined in go.mod file
// Every sub-directory is a sub-package.
package main

import (
	"fmt"

	"github.com/bijaythapaa/MakaluGo/packer/numbers"
	"github.com/bijaythapaa/MakaluGo/packer/strings"
	"github.com/bijaythapaa/MakaluGo/packer/strings/greeting"
)

func main() {
	fmt.Println(numbers.IsPrime(19))
	fmt.Println(strings.Reverse("Amar"))
	fmt.Println(greeting.Welcome)
	fmt.Println(greeting.MorningText)
	// fmt.Println(greeting.eveningText)
	fmt.Println(greeting.EveningText)
}
