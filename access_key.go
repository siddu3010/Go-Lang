// Golang map access keys

package main
import "fmt"

func main()  {
	
places := map[string]string{"Nepal": "kathmandu", "US": "Washington DC", "Norway": "Oslo"}

for country := range places {
	fmt.Println(country)
	}	

}