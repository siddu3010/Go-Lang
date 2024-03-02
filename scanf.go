// // // package main
// // // import "fmt"

// // // func main() {
// // // 	var name,last string

// // // 	fmt.Print("Enter your name:")
// // // 	fmt.Scanf("%s %s", &name,&last)
// // // 	fmt.Println("Welcome!!!",  name, last)
// // // }

// // package main
// // import "fmt"

// // func main() {

// // 	fmt.Print("Enter your name:")
// // 	var name string
// // 	fmt.Scanln(&name)
	
// // 	fmt.Print("Enter your last name:")
// // 	var last_name string
// // 	fmt.Scanln(&last_name)
// // 	fmt.Println("Welcome!!!", name  +" "+ last_name)
// // }

// // package main
// // import "fmt"

// // func main() {
// // var a,b int = 10,20
// // fmt.Println(a,b)
// // fmt.Println(a+b)
// // }

// // package main
// // import "fmt"

// // func main() {
// // var name, age, branch, college = "Siddhesh Singabhatti", 22, "AIML", "RCOEM" 
// // fmt.Println("Name:",name,"\t","Age:",age)
// // fmt.Println("Branch:",branch,"\t","College:",college)
// // }

// package main
// import "fmt"
// func main() {
// 	var a string = "Siddhesh"
// 	fmt.Printf("Value of a: '%s' ",a)
// 	fmt.Printf("Value of a: \"%s\"",a)
// }



// // Print: No new line is generated

// // Printf

// // Println


// 


// package main
// import("fmt")
// func main() {
// 	var str = "Siddhesh"
// 	fmt.Printf("Name: %s\n",str)
// 	var n1 int = 21
// 	fmt.Printf("Decimal: %d\n",n1)
// 	var n2 float32 = 7.784
// 	fmt.Printf("Float: %g\n",n2)
// 	var n3 int = 14
// 	fmt.Printf("Binary: %b\n",n3)
// 	var n4  = 4 + 1i
// 	fmt.Printf("Complex: %e\n",n4)

// 	fmt.Printf("%v %v %v %v %v",str,n1,n2,n3,n4)
// }

// Literals are always r valued

// package main
// func main() {
// 	println(15 == 017)
// 	println(15 == 0xF)
// }

// package main
// import "fmt"
// func main() {
// var a =15
// fmt.Printf("Octal: %o\n",a)
// fmt.Printf("Hexdecimal: %x\n",a)
// fmt.Printf("Octal: %b\n",a)

// fmt.Printf("Expression 15 == 0xF: %t\n",15 == 0xF)
// fmt.Printf("Expression 15 == 0xF: %T\n",15 == 0xF)
// }



// package main
// import "fmt"
// func main() {

// var i,j string = "Hello", "Siddhesh"
// var a,b int = 10, 20
// var x,y = 100, "Siddhesh"
// fmt.Printf("%s\n%s\n",i,j)
// fmt.Printf("%s %s\n",i,j)
// fmt.Print(a,b)
// fmt.Print("\n")
// fmt.Print(x,y)


// }



// package main
// import "fmt"
// func main() {
// 	var i = 15.5
// 	var txt = "Hello World!"
// 	fmt.Printf("%v\n",i)
// 	fmt.Printf("%#v\n",i)
// 	fmt.Printf("%v %%\n",i)
// 	fmt.Printf("%T\n",i)
// 	fmt.Printf("%v\n",txt)
// 	fmt.Printf("%#v\n",txt)
// 	fmt.Printf("%T\n",txt)
// 	fmt.Printf("%s",txt)
// }



// package main
// import "fmt"
// func main() {
// 	var i = 15

// 	fmt.Printf("%b\n",i)
// 	fmt.Printf("%d\n",i)
// 	fmt.Printf("%+d\n",i)  // + sign used here is for sign of the no.
// 	fmt.Printf("%o\n",i)
// 	fmt.Printf("%O\n",i)
// 	fmt.Printf("%x\n",i)
// 	fmt.Printf("%X\n",i)
// 	fmt.Printf("%#x\n",i)
// 	fmt.Printf("%4d\n",i)
// 	fmt.Printf("%-4d\n",i)
// 	fmt.Printf("%04d\n",i)
// }



package main
import "fmt"
func main() {
	var txt = "Hello"

	fmt.Printf("%s\n",txt)
	fmt.Printf("%q\n",txt)
	fmt.Printf("%8s\n",txt)
	fmt.Printf("%-8s\n",txt)
	fmt.Printf("%x\n",txt)
	fmt.Printf("% x\n",txt)
	
}