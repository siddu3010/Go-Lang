// // package main
// // import "fmt"
// // const Golden  = 1.618033988749895;
// // func main() {
// // 	fmt.Println(Golden)
// // }

// package main
// import "fmt"
// const A int = 10   // Typed constant
// func main()  {
// 	fmt.Println(A)
// }

// package main
// import "fmt"
// const A = 10   // Un-typed constant
// func main()  {
// 	fmt.Println(A)
// }

// package main
// import "fmt"
// const(
// 	A int = 100
// 	B = 1.618033
// 	C = "Siddhesh :)"
// )
// func main()  {
// 	fmt.Println("\n",A,"\n",B,"\n",C,"\n")
// }


package main
import "fmt"
func main()  {
	const(
		l int = 100
		b int = 80
	)
	var area int
	area = l*b
	fmt.Printf("Area of Rectangle: %d", area)
}