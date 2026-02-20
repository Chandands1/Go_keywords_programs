package main
import (
	"fmt"
	"time"
)

func say() {
	fmt.Println("Running")
}

func main() {
	go say()
	time.Sleep(time.Second)
}