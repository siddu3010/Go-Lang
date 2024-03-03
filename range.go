// package main
// import "fmt"
// func main() {
// 	colors :=[]string{"Red","Blue","Green","Yellow"}
// 	for index,val:= range colors {
// 		fmt.Println(index,"=>",val)
// 	}
// }

// package main
// import "fmt"
// func main() {
// 	colors :=[]string{"Red","Blue","Green","Yellow"}
// 	for _,val:= range colors {
// 		fmt.Println("=>",val)
// 	}
// }

package main
import "fmt"
func main() {
	colors :=[]string{"Red","Blue","Green","Yellow"}  // Slice type not array
	for index,_:= range colors {
		fmt.Println("=>",index)
	}
}