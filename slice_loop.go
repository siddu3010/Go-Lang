// package main
// import "fmt"

// func main() {
// 	numbers := []int{2,4,6,8,10}

// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}
// }

// Iterating a slice using range in for loop

package main
import "fmt"

func main()  {
	arr := [10]int{10,20,30,40,50,60,70,80,90,100}

	intSlice := arr[2:5]
	// Slice = []int{30,40,50}  aisa banega slice

	fmt.Println("Slice Elements:")
	for index, ele := range intSlice {
		fmt.Printf("Index = %d, Element = %d\n",index,ele)
	}
}