package main
import "fmt"

func main() {
    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
    var lastMatch string
    fmt.Scanln(&lastMatch)
    results = append(results, lastMatch)
    point := 0

    for _, v := range results {
        switch v {
            case "w":
                point += 3
            case "d":
                point += 1
        }
    }
    fmt.Println(point)
}