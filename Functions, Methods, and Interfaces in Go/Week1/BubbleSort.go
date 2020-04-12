
package main
import 
(
	"fmt"
)

func Swap(slice []int ,i int){
	if slice[i] > slice[i+1] {
		temp:= slice[i]
		slice[i] = slice[i+1]
		slice[i+1] = temp
	}
}


func BubbleSort(slice []int){
   n := len(slice)


   for i:=0 ; i<n-1 ;i=i+1{
   	for j:=0;j<n-i-1;j=j+1{
        Swap(slice ,j)
        
   	}
   }
}
func main() {
    slice := make([]int, 10)
    fmt.Println("Enter 10 integer")

    for i:=0 ; i<10 ; i = i+1 {
       fmt.Scan(&slice[i])   
    }

    BubbleSort(slice)

    for i:=0 ; i<10 ; i = i+1 {
       fmt.Println(slice[i])   
    }

}