
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
  
    "strings"
)

type name struct  {
    fname string
    lname string
}


func main() {
    var filename string
    fmt.Println("Enter file name")
    fmt.Scan(&filename)
    
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()
    
   
    slice := make([]name, 0, 1)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
       words := strings.Fields(scanner.Text())
       
       slice = append(slice, name{words[0], words[1]})
       
    }
    
    for _, x := range slice {
		fmt.Printf("%s  %s\n", x.fname, x.lname)
	}

   
}