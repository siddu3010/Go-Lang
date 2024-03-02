package main
import (

	"fmt"
    "strings"
)

func main() {
	var str string = "Hello World"
	var subStr string = "War"

if strings.Contains(str, subStr) == true {
	fmt.Printf("String %s contains substring %s", str,subStr)
}else{
	fmt.Printf("String %s does not contains substring %s", str,subStr)
}

}