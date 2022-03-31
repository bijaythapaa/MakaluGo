package main
import "fmt"

func mars_age(age int) int {
    mars_day, earth_day := 687, 365
    return (age*earth_day)/mars_day
}

func main() {
    var age int
    fmt.Scanln(&age)

    mars := mars_age(age)
    fmt.Println(mars)
}