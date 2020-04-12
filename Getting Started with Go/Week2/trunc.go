package main

import "fmt"

func main() {
	fmt.Printf("Enter first floating point number\n")
        var x float64
        
        fmt.Scan(&x)
          
        y := int(x)
        fmt.Printf("%d\n",y)

        fmt.Printf("Enter second floating point number\n")
        
        fmt.Scan(&x)
          
        y = int(x)
        fmt.Printf("%d\n",y)

       
        
          
}