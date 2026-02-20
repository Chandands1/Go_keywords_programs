package main
import "fmt"

func main() {
	ch := make(chan string)
	go func() { ch <- "Hello" }()
	select {
	case msg := <-ch:
		fmt.Println(msg)
	}
}