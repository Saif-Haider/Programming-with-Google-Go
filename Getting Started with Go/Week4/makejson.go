
package main
import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
)


func main() {
    var name , address string
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter name")
    
    if scanner.Scan() {
		name = scanner.Text()
	}
	
	fmt.Println("Enter address")
	
	if scanner.Scan() {
		address = scanner.Text()
	}
    
    m := make(map[string]string)
    
    m["name"] = name
    m["address"] = address
    
    jsonObject , err := json.Marshal(m)
    if err!=nil {
         fmt.Println("Some error occured\n")
    
    }
    fmt.Println(string(jsonObject))
}