package main
import "fmt"

func main() {
	nums := []int{1,2,3}
	for i, v := range nums {
		fmt.Println(i, v)
	}
}