// errors is a type to represent any issue in the execucation of program flow

// "error" type is an interface that have one function:
// type error interface {
// 	Error() string
// }

package main

import (
	"errors"
	"fmt"
)

// this function will return error if we got true parameter
func getError(b bool) (err error) {
	if b {
		// err = fmt.Errorf("This is an error :(")
		err = errors.New("this is an error :(")
	}
	return
}

type invalidInput struct {
	fieldName string
}

// implement error interface in invalidInput struct
func (i invalidInput) Error() string {
	return fmt.Sprintf("invalid input on field '%v' ", i.fieldName)
}

func validateInput(field, value string) (err error) {
	if value == "" {
		err = invalidInput{fieldName: field}
	}
	return
}

func main() {
	// {
	// 	err := getError(true)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }

	// call to validate input func with valid input
	// err := validateInput("firstName", "Amar")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// check on error value
	// the errors.Is func compares and error to value
	// {
	// err := invalidLength("")
	// err := invalidLength("Amar")
	// if errors.Is(err, errInvalidLength) {
	// 	fmt.Println("We know err value using 'error.Is'")
	// } else {
	// 	fmt.Println("We donot know about this error")
	// }
	// fmt.Println(errors.Is(err, errInvalidLength))
	// }

	// check on error type
	// the errors.Is func compares and error to type
	// {
	// 	err := validateInput("name", "")
	// 	if errors.As(err, &invalidInput{}) {
	// 		fmt.Println("We know err type using 'error.As'")
	// 	} else {
	// 		fmt.Println("We donot know about this error")
	// 	}
	// }

	{
		// err := fmt.Errorf("could not save your input: '%v'", errInvalidLength)
		err := fmt.Errorf("could not save your input: '%w'", errInvalidLength)
		if errors.Is(err, errInvalidLength) {
			fmt.Println("we know err value using 'errors.Is'", err.Error())
			fmt.Println("this is an unware message", errors.Unwrap(err))
		} else {
			fmt.Println("we donot know about this error !!")
		}
	}
}

var errInvalidLength = errors.New("Invalid length")

func invalidLength(s string) error {
	if len(s) < 1 {
		return errInvalidLength
	} else if s == "error" {
		return errors.New("Invalid input")
	}
	return nil
}
