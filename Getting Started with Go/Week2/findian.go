package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

 
func main() {
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("Enter First String:\n")  

        var first string    
       if scanner.Scan() {
		first =  scanner.Text()
	}  

	 var r rune 
           r = 'a'
         first = strings.ToLower(first) 
         Len:=len(first)
        
         if (first[0] == 'i' && first[Len-1] =='n' && strings.ContainsRune(first, r)) {
         	fmt.Print("Found!\n")
         } else
         {
         	fmt.Print("Not Found!\n")
         }

         
}