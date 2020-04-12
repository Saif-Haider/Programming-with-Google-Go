// Rance condition is a condition when the output of the program
// depends on the timing of one or more processes threads to function correctly.
// In this code race condition occur as both routine access x and modify it. 
// Both of them print a value which will differ in case the are run concurrently.


package main

import "fmt"

func first(x *int){
    (*x)++
    fmt.Println("Hello ",*x)
}


func second(x *int){
    (*x)--
    fmt.Println("Hello ",*x)
}
func main() {
   x:=5
   go first(&x)
   go second(&x)  
     
    fmt.Println(x)   
}