package ranges

import "fmt"

func main() {
	// Range on array
	arr := [4]string{"AMar", "Jyoti", "Rekha"}

	for k, v := range arr {
		fmt.Println("range on array key: ", k, " value: ", v)
	}

	// Range on slice
	slc := []string{"x", "y", "z"}
	for k, v := range slc {
		fmt.Println("key: ", k, " value: ", v)
	}

	// Range on map
	mp := map[string]string{
		"admin": "amar",
		"user":  "bijay",
	}
	for key, value := range mp {
		fmt.Println(key, value)
	}

	// Range on map of slices
	mapOfSlices := map[string][]string{
		"Admins": []string{"bijay", "amar", "jyoti"},
		"Users":  []string{"pratiksha", "rekha", "pramila"},
	}
	for key, value := range mapOfSlices {
		fmt.Println(key, value)

		for k, v := range value {
			fmt.Println(k, v)
		}
	}
}
