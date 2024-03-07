// package main
// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	slice := []int{70,20,30,60,50,40,10,90,80,100}


// 	sort.Ints(slice)

// 	fmt.Println("Sorted Slice:")
// 	for _,ele := range slice {
// 		fmt.Printf("%d ",ele)
// 	}

	
// }


package main
import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"honesty","is","the","best","policy"}


	sort.Strings(slice)

	fmt.Println("Sorted Slice:")
	for _,item := range slice {
		fmt.Printf("%s ",item)
	}

}