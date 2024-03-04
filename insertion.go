package main
import "fmt"
func main() {


	originalArray := []int{1,2,3,4,5} // Original array
	newItem := 6  //Item to insert
	index := 2  //Index at which item is inserted
	newArray := insertIntoArray(originalArray,newItem,index) //Insert the item into array

	//Print original and new array	
	fmt.Println("Original Array:",originalArray) 
	fmt.Println("Item to insert:",newItem)
	fmt.Println("Index to insert:",index)
	fmt.Println("New Array:",newArray)

}

func insertIntoArray(arr[]int, item, index int)[]int {
	newArr := make([]int, len(arr)+1)
	copy(newArr[:index], arr[:index])
	newArr[index] = item
	copy(newArr[index+1:], arr[index:])
	return newArr
}