// Using make() funcation map

package main
import "fmt"

func main() {
student := make(map[int]string)

student[1] = "Harold"
student[2] = "Martin"
student[3] = "Luther"

fmt.Println(student)

}