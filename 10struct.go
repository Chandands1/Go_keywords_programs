package main
import "fmt"

type Person struct {
	Name string
}

func main() {
	p := Person{Name: "Go"}
	fmt.Println(p.Name)
}