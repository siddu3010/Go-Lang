package main
import "fmt"

type Vertex struct {
	Lat, Long float64  //Latitude and Longitude 
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	}, 
	"Google Headquarters": Vertex{
		37.44202, -122.08408,
	},
}

func main()  {
	fmt.Println(m)

m["Nagpur"] = Vertex{21.146633, 79.088860}

fmt.Println(m)

}