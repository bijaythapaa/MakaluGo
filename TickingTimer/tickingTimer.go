package main
import "fmt"

type Timer struct {
    id int
    value int
}

func (time *Timer) tick() {
    time.id += 1
    // fmt.Println(time.id)
}


func main() {
    var x int
    fmt.Scanln(&x)

    t := Timer{0, x}
    
    for i:=0;i<x;i++ {
        t.tick()
        fmt.Println(t.id)
    }
}