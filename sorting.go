// package main
// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	nums := []float64{4.2,2.7,7.1,1.5,9.0,3.3}
// 	sort.Float64s(nums)
// 	fmt.Println("\nSorted Order:",nums)
// }

package main
import (
	"fmt"
	"sort"
)

func main() {
	words:=[]string{"banana","apple","orange","grape","pineapple"}
	sort.Strings(words)
	fmt.Println("Sorted Order:",words)
}


