// package main
// import(
// 	"fmt"
// 	// "time"
// )

// func  calculate(x int , y int)(int,int)  {
// 	return x+y , x-y 	
// }

// func main () {
// 	sum, difference := calculate(10,20)
// 	fmt.Printf("Additionis %d\n Subtraction is %d\n",sum,difference)
// }

// package main
// import(
// 	"fmt"
// 	"time"
// )

// func main() {
// 	yyyy,mm,dd := time.Now().Date()
// 	fmt.Println("Date:",dd)
// 	fmt.Println("Month:",mm)
// 	fmt.Println("Year:",yyyy)
// }

// package main
// import(
// 	"fmt"
// 	"time"
// )

// func main() {

// 	currentDateTime := time.Now()

// 	day := currentDateTime.Day()
// 	month := currentDateTime.Month()
// 	year := currentDateTime.Year()

// 	hour := currentDateTime.Hour()
// 	min := currentDateTime.Minute()
// 	sec := currentDateTime.Second()

// 	fmt.Println("\nDate:",day)
// 	fmt.Println("Month:",month)
// 	fmt.Println("Year:",year)
// 	fmt.Println("Hour:",hour)
// 	fmt.Println("Minutes:",min)
// 	fmt.Println("Seconds:",sec)
// }

package main
import(
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello World")
	time.Sleep(1000*time.Millisecond)
	fmt.Println("Namaste Bharat")
	time.Sleep(500*time.Millisecond)
	fmt.Println("Bharat Mata ki Jai. Jai Shri Ram.")
}


