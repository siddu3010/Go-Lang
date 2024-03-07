package main
import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int 
}
func main()  {
var b1,b2 Books 

b1 = Books{"Go Programming\n","Mahesh Kumar\n","Go Tutorial\n",6495407}
b2 = Books{"Telecom Billing\n","Zara Ali\n","Telecom Tutorial\n",6495700}

// fmt.Println("Book 1:\n", b1)
// fmt.Println("Book 2:\n", b2)

// fmt.Printf("Book 1 Title: %s",b1.title)
// fmt.Printf("Book 1 Author: %s",b1.author)
// fmt.Printf("Book 1 Subject: %s",b1.subject)
// fmt.Printf("Book 1 ID: %d",b1.book_id)
// fmt.Println("\n")
// fmt.Printf("Book 2 Title: %s",b2.title)
// fmt.Printf("Book 2 Author: %s",b2.author)
// fmt.Printf("Book 2 Subject: %s",b2.subject)
// fmt.Printf("Book 2 ID: %d",b2.book_id)

printBook(b1,1)
printBook(b2,2)
}

func printBook(book Books ,number int)  {
fmt.Printf("Book %d Title: %s",number,book.title)
fmt.Printf("Book %d Author: %s",number,book.author)
fmt.Printf("Book %d Subject: %s",number,book.subject)
fmt.Printf("Book %d ID: %d\n\n",number,book.book_id)
	
}





