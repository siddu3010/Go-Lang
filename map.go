// package main
// import "fmt"

// func main() {
// 	subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}

// 	fmt.Println(subjectMarks)
// }


//  Access values of map

// package main
// import "fmt"

// func main() {

// 	flowerColor := map[string]string{"Sunflower": "Yellow", "Jasmine": "White", "Hibiscus":"Red"}

// 	fmt.Println(flowerColor["Sunflower"])
// 	fmt.Println(flowerColor["Hibiscus"])
// }

// Change value in a map

// package main
// import "fmt"

// func main() {
// 	capital := map[string]string{"Nepal": "Kathmandu", "US": "New York"}
// 	fmt.Println("Initial Map: ",capital)

// 	capital["US"] = "Washington DC"

// 	fmt.Println("Updated Map: ",capital)
// }

// Add elements in map 

// package main
// import "fmt"

// func main() {
// 	students := map[int]string{1: "John", 2: "Lily"}
// 	fmt.Println("Initial Map: ",students)
	
// 	students[3] = "Robin"

// 	students[5] = "Julie"

// 	fmt.Println("Updated Map: ",students)
// }



// Delete Element from Map

package main
import "fmt"

func main() {
	personAge := map[string]int{"John": 10, "Lily": 12, "Harold": 16}
	fmt.Println("Initial Map: ",personAge)
	
	delete(personAge,"Lily")

	fmt.Println("Updated Map: ",personAge)
}


package main
import "fmt"

func main() {
	personAge := map[string]int{"John": 10, "Lily": 12, "Harold": 16}
	fmt.Println("Initial Map: ",personAge)
	
	delete(personAge,"Lily")

	fmt.Println("Updated Map: ",personAge)
}


