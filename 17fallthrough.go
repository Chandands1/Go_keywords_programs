package main
import "fmt"

func main() {
	switch 1 {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}
}