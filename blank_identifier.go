package main
import "fmt"

func main()  {
	places := map[string]string{"Nepal": "Kathmandu", "US": "Washington DC", "Norway": "Oslo"}

for _, capital := range places {
	fmt.Println(capital)
}


}