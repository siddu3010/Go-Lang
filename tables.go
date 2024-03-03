package main
import "fmt"
func main() {

	multiplier := 1
	var n int
	fmt.Printf("Enter the Number: ")
	fmt.Scanf("%d\n",&n)
	for multiplier <= 10 {
		product :=  n*multiplier
		fmt.Printf("5 x %d = %d\n",multiplier,product)
		multiplier++
	}
}