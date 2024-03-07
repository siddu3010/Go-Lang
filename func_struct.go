// Function inside a struct in Go

package main
import "fmt"

type rectangle func(int,int)int

type rectanglepara struct {
	l int
	b int
	color string


	rect rectangle
} 

func main() {
	result := rectanglepara{
		l: 10,
		b: 20,
		color: "Red",
		rect: func (l int, b int) int  {
			return l*b
		},
	}

	fmt.Println("Color of the rectangle is : ", result.color)
	fmt.Println("Area of the rectangle is : ", result.rect(result.l, result.b))
}