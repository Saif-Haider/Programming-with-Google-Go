
package main
import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct {}

func (C Cow) Eat(){
     fmt.Println("grass") 
}


func (C Cow) Move(){
     fmt.Println("walk") 
}


func (C Cow) Speak(){
     fmt.Println("moo") 
}


func (B Bird) Eat(){
     fmt.Println("worms") 
}


func (B Bird) Move(){
     fmt.Println("fly") 
}


func (B Bird) Speak(){
     fmt.Println("peep") 
}

func (S Snake) Eat(){
     fmt.Println("mice") 
}


func (S Snake) Move(){
     fmt.Println("slither") 
}


func (S Snake) Speak(){
     fmt.Println("hsss") 
}
func main(){
    
    
    
    m := make(map[string]Animal)
    
     var cow Cow
     var bird Bird
     var snake Snake
    
    for {
        
     fmt.Println(">")
     var a1 , a2 ,a3 string
     fmt.Scan(&a1)
     fmt.Scan(&a2)
     fmt.Scan(&a3)
     
      if a1 =="newanimal"{
      	if a3 == "cow"{
      		m[a2] = cow
      	}else if a3 == "bird"{
            m[a2] = bird
      	}else{
      		m[a2] = snake
      	}
      	fmt.Println("Created it!")
      }else{
      	    tmp := m[a2]

      	    switch a3{
      	    	case "eat" : tmp.Eat()
      	    	case "move" : tmp.Move()
      	    	case "speak" : tmp.Speak()
      	    }  
      }

    
        
    } 
     

}