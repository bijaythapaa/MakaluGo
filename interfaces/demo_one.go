package main

import (
	"errors"
	"fmt"
)

// Interfaces are types, it defines a set of method signatures that should be implemented by struct.
// allows us to make functions that can be used with different types and create highly-decoupled code
// whilst still maintainging type-safety.

// type, something that walks
type Walker interface {
	walk(p Point) error
	getPosition() Point
}

// type, something that talks
type Talker interface {
	talk(s string)
}

type Point struct {
	x int
	y int
}

type Human struct {
	position Point
}

// implement interface Walker on human
// interface implementation comes implicitly when struct implements
// all the functions in interface type
func (h *Human) walk(p Point) error {
	if p.x < 0 || p.y < 0 {
		return errors.New("Invalid Point ")
	}
	h.position = p
	fmt.Println("human walked to", h.position)
	return nil
}

func (h *Human) getPosition() Point {
	return h.position
}

// implement interface talker on human
func (h *Human) talk(s string) {
	fmt.Println("This human talking: ", s)
}

// animal
type Animal struct {
	position Point
}

// implement interface Walter on struct Animal
func (a *Animal) walk(p Point) error {
	if p.x < 0 || p.y < 0 {
		return errors.New("invalid point")
	}
	a.position = p
	fmt.Println("animal walked to:  ", a.position)
	return nil
}

func (a *Animal) getPosition() Point {
	return a.position
}

func main() {
	amar := &Human{}
	dog := &Animal{}
	steps := []Point{
		{1, 1},
		{2, 2},
		{3, 3},
	}
	// fmt.Println(steps)
	err := move(amar, steps)
	if err != nil {
		fmt.Println(err)
	}

	err = move(dog, steps)
	if err != nil {
		fmt.Println(err)
	}
	// err = move(dog, []Point{Point{-1, 2}})
	err = move(dog, []Point{{-1, 2}})
	if err != nil {
		fmt.Println(nil)
	}

	moveHuman(amar, steps)
}

// function takes interface as parameter type
// can take any struct type implements the interface
func move(w Walker, points []Point) error {
	for _, p := range points {
		err := w.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}

// function takes interface as parameter type
// can take any struct type implements the interface
func moveHuman(w *Human, points []Point) error {
	for _, p := range points {
		err := w.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}
