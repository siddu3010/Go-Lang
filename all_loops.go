package main
import "fmt"

func main() {

	number := 1

	for {
		fmt.Printf("%d\n",number)
		number ++
		if number > 5 {
			break
		}
	}
}