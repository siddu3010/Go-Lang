// HTTP GET resquest API

package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main()  {

	apiUrl := "https://www.api.example.com/data"

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error making GET request: %s\n",err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}
fmt.Println(string(body))

}