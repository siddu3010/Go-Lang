package main
import(
	"fmt"
)
func main() {
	var x interface{}  // var x interface{} = "RKNEC" (Try)
	switch i := x.(type) {
		case nil:
		fmt.Printf("Type of X: %T",i)
		case int:
		fmt.Println("X is int")
		case float64:
		fmt.Println(" X is float 64")
		case func(int):
		fmt.Println(" X is func{int}")
		case bool,string:
		fmt.Println("X is bool or string")
		default:
		fmt.Println("Dont  know the type of X")
	}
}
