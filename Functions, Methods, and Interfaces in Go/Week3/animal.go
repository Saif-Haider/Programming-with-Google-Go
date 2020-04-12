
package main
import "fmt"

type Animal struct{
    food string
    locomotion string 
    noise string
}

func (A Animal) Eat(){
     fmt.Println(A.food) 
}


func (A Animal) Move(){
     fmt.Println(A.locomotion) 
}


func (A Animal) Speak(){
     fmt.Println(A.noise) 
}
func main(){
    
    var A [3]Animal
    A[0].food = "grass"
    A[0].locomotion = "walk"
    A[0].noise = "moo"
    
    A[1].food = "worms"
    A[1].locomotion = "fly"
    A[1].noise = "peep"
    
    A[2].food = "mice"
    A[2].locomotion = "slither"
    A[2].noise = "hsss"
    
    m := make(map[string]int8)
    
    m["cow"] = 0
    m["bird"] = 1 
    m["snake"] = 2 
    
    for {
        
     fmt.Println(">")
     var name , information  string
     fmt.Scan(&name)
     fmt.Scan(&information)
     
     i := m[name]
     
     if information == "eat"{
         A[i].Eat()
     }else if information == "move"{
         A[i].Move()
     }else {
         A[i].Speak()
     }
     
        
    } 
     

}