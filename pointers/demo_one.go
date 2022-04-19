package main

import "fmt"

func main() {
	name := "Bijay"
	fmt.Println(name)

	fmt.Println("&name: ", &name)
	fmt.Println("*name: ", *&name)
	fmt.Println("is name and *name are equal ?", name == *&name)

	//sending copy of the variable value
	updateCopy(name)
	fmt.Println(name)

	// sending pointer to the variable by adding '&' before it
	updateReference(&name)
	fmt.Println(name)

	// if parameter is a pointer, we can send nil as argument
	defaultValuePointer(nil)

	// var kValue string
	// fmt.Println("kValue: ", kValue)

	var pValue *string
	// fmt.Println("pValue: ", pValue)
	pValue = &name
	fmt.Println("pValue: ", pValue, *pValue)
}

func updateCopy(s string) {
	s = "Jyoti"
}

func defaultValuePointer(s *string) {
	fmt.Println("nil val: ", s)
}

func updateReference(s *string) {
	// fmt.Println("just s: ", s, "&s: ", &s, "*s: ", *s)
	*s = "Amar"
	fmt.Println("after change: just s: ", s, "&s: ", &s, "*s: ", *s)

	n := "Updated"
	s = &n
	fmt.Println("assigned to new memory address")
	fmt.Println("just s: ", s, "&s: ", &s, "*s: ", *s)
	fmt.Println("just n: ", n, "&n: ", &n)
}
