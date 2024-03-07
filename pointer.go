package main
import "fmt"

func main() {

	var intData = 20

	var intPointer *int

	intPointer = &intData

	fmt.Println("What intData stores?:",intData)
	fmt.Println("Address of intData:",&intData)
	fmt.Println("What intPointer stores:",intPointer)

	*intPointer = 30

	fmt.Println("What intData now stores:",intData)

}