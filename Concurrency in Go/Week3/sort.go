
package main

import (
    "fmt"
    "sort"
    "sync"
)
func sortit(wg *sync.WaitGroup,arr []int){
  sort.Ints(arr)
  fmt.Println(arr)
  wg.Done()
}

func merge(wg *sync.WaitGroup,left []int, right []int,result []int) {
   
      
    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }
      
    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }
      
   wg.Done()
}

func main() {
   var arr []int
   var x ,n  int
   fmt.Println("N = ")
   fmt.Scan(&n)

   fmt.Println("Enter ",n," number")
   for i:=0 ;i<n; i++ {
       fmt.Scan(&x)
       arr = append(arr,x)
   }
   
   blk := n/4
   var wg sync.WaitGroup
   wg.Add(1)
   go sortit(&wg,arr[0:blk])
   wg.Add(1)
   go sortit(&wg,arr[blk:2*blk])
   wg.Add(1)
   go sortit(&wg,arr[2*blk:3*blk])
   wg.Add(1)
   go sortit(&wg,arr[3*blk:])
   
   wg.Wait()
   
   m1 := make([]int ,len(arr[0:blk]) + len(arr[blk:2*blk]))
   m2 := make([]int ,len(arr[2*blk:3*blk]) + len(arr[3*blk:]))
   
   wg.Add(1)
   go merge(&wg,arr[0:blk],arr[blk:2*blk],m1)
   wg.Add(1)
   go merge(&wg,arr[2*blk:3*blk],arr[3*blk:],m2)
   wg.Wait()
   
   m3 := make([]int,len(m1) + len(m2))
   wg.Add(1)
   go merge(&wg,m1,m2,m3)
   wg.Wait()
   
   fmt.Println(m3)
   
   
}