package main
import(
	"fmt"
)

func main() {
   thisMonth := 6

   switch thisMonth {
case 1,3,5,7,8,10,12:
	fmt.Println("31 Days")
case 4,6,9,11:
	fmt.Println("31 Days")
case 2:
	fmt.Println("28 or 29 Days")
}

}