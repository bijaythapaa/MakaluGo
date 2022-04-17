package maps

import "fmt"

func main() {
	m := map[string]bool{
		"x": true,
		"y": false,
	}
	fmt.Println("map: ", m)

	player := make(map[string]int)
	fmt.Println("player: ", player)

	player["Bijay"] = 24
	player["Amar"] = 19
	player["Nanu"] = 28
	fmt.Println("player: ", player)

	fmt.Println("player", player["Amar"], "next", player["Nanu"])

	delete(player, "Bijay")

	// check if key exists and get value
	// value will be default value if key does not exists
	v, ok := player["Jyoti"]
	fmt.Println("missing key, Jyoti", v, "ok", ok)

	// map of slices
	mp := map[string][]string{
		"admins": []string{"Amar", "Bijay"},
		"users":  []string{"Jyoti", "Nanu"},
	}
	fmt.Println(mp)
}
