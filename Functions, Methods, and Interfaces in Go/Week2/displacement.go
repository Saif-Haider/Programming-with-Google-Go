package main
import "fmt"

func GenDisplaceFn(a float64,v float64 ,s float64) func(float64)float64{
    fn:= func(t float64)float64{
        s1 := 0.5*a*t*t  + v*t  + s
        return s1
    }
    
    return fn
    
}
func main(){
     var a ,v ,s ,t float64 
     fmt.Println("acceleration = ");
     fmt.Scan(&a)
     fmt.Println("initial velocity = ");
     fmt.Scan(&v)
     fmt.Println("initial displacement = ");
     fmt.Scan(&s)
     
     fn := GenDisplaceFn(a,v,s)
     fmt.Println("time = ");
     fmt.Scan(&t)
     
     fmt.Println("displacement after ",t," = ",fn(t))
     
     

}