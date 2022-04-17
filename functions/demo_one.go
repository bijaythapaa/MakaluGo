package main

import (
	"errors"
	"fmt"
)

// multiple arguments acceptance
func callMe(params ...interface{}) {
	fmt.Println("Length: ", len(params))
}

func sum(x, y int) int {
	return x + y
}

// return multiple types
func add(x, y string) (string, int) {
	r := x + y
	return r, len(r)
}

// named return type
func concat(x, y string) (res string, length int) {
	res = x + y
	length = len(res)
	// will return the named vars
	return
}

// function that returns string or nil or default string value and error
func returnError(b bool) (string, error) {
	if b {
		return "", errors.New("this is a Go error")
	}
	return "okay", nil
}

func main() {
	callMe("Hey, ", "Greyshadow !!", 24, 22.22)
	sum(2, 5)
	r, l := add("Bijay", "Thapa")
	fmt.Println(r, l)
	res, err := returnError(true)
	if err != nil {
		fmt.Println("We have error: ", err.Error())
	} else {
		fmt.Println("We got the result: ", res)
	}

	// type of function
	var x func(i, j int) int
	x = sum
	fmt.Println("x: ", x(11, 22))
}
