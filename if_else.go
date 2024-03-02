// // // package main
// // // import ("fmt")
// // // func main() {
// // // 	if 20 > 18 {
// // // 		fmt.Println( "20 is greater than 18" )
// // // 	}
// // // }

// // // package main
// // // import ("fmt")
// // // func main() {
// // // 	var a,b = 20,18
// // // 	if a > b {
// // // 		fmt.Println( "a is greater than b" )
// // // 	}
// // // }



// // // package main
// // // import ("fmt")
// // // func main() {
// // // 	if !(20 < 18) {
// // // 		fmt.Println( "20 is greater than 18" )
// // // 	}
// // // }

// // package main
// // import ("fmt")
// // func main() {
// // 	x:=8
// // 	if y:=10;  x<y {
// // fmt.Println( "x is less than y" )
// // 	}
// // }





// // package main
// // import (
// // 	"fmt"
// // 	"os"
// // )
// // func main() {
// // 	var name string
// // 	var age int

// // 	if _, err:= fmt.Scan(&name,&age); err!=nil{
// // 		fmt.Println(err)
// // 		os.Exit(1)
// // }
// // fmt.Printf("Name:%s\n",name)
// // fmt.Printf("Age:%d\n",age)

// // }


// // Write a program to print greatest of the 3 no.s  using if else control structure.

// package main
// import ("fmt")
// func main() {
// 	var a,b,c = 10,20,30
	
// 	if  (a>b && a>c) {
// 		fmt.Println("a is the greatest")		
// 	} else if  (b>a && b>c) {
// 		fmt.Println("b  is the greatest")
// 	}else{
// 		fmt.Println("c is the greatest")
// 	}

	
// }
	

// If - Else

// package main
// import "fmt"

// func main() {
// 	x:=8
// 	y:=10

// 	if(x < y) {
// 	fmt.Println(x,"is less than",y)
// 	}else{
// 	fmt.Println(x,"is greater than",y)
// 	}
// }


package main
import "fmt"

func main() {
	var x int 
	fmt.Scanf("%d\n",&x)
	if(x%2==0){
		fmt.Printf("%d is an even number\n",x)
	}else{
		fmt.Printf( "%d is an odd number \n",x)
	}
}