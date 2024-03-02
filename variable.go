// // package main
// // import "fmt"
// // func main() {
// // var student1 string = "Siddhesh"
// // var student2 = "Bhai"
// // x:= "Don"

// // 	// fmt.Println(student1)
// // 	// fmt.Println(student2)
// // 	// fmt.Println(x)

// // 	fmt.Printf("%s\n",student1)
// // 	fmt.Printf("%s\n",student2)
// // 	fmt.Printf("%s\n",x)


// // }




// // package main
// // import "fmt"
// // func main() {
// // var x float64 = 20.0
// // y := 42

// // 	fmt.Printf("%f\n",x)
// // 	fmt.Printf("%d\n",y)

// // 	fmt.Println(x)
// // 	fmt.Println(y)

// // }

// package main
// import "fmt"
// func main() {
// var x = 20
// y := 40
// var z bool

// 	fmt.Println(x)
// 	fmt.Println(y)
// 	bool (y>x) {
// 		fmt.Println(z)	
// 	}
	

// }


// package main 
// import "fmt"
// func main() {
// 	var a,b,c,d int = 1,3,5,7
// 	fmt.Println(a,b,c,d)
// }

// package main 
// import "fmt"
// func main() {
// 	var a, b, c, d = 100, "God of War","Kratos", 99.99
// 	fmt.Println(a ,b ,c ,d)
// }

// package main 
// import "fmt"
// func main() {
// 	var(
// 		a int
// 		b int = 20
// 		c string = "Kratos"
// 	) 
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)

// 	fmt.Printf("\n %T\n %T\n %T\n ", a, b, c)
// 	fmt.Printf("\n %d\n %d\n %s\n ", a, b, c)
// }


// Var can be used inside and outsdie the functions that is it is global variable

// := can be used only within a function that is as a local variable

package main 
import "fmt"
var a = 1
var b = 2
func main() {
	fmt.Println(a)
	fmt.Println(b)
}