package main

import (
	"fmt"
	"sort"
	"strconv"
	
)

func main() {
	slc := make([]int, 0, 3)

	for {
		var num1 string

		fmt.Println(" number = ")
		fmt.Scan(&num1)

		
		if num1 == "X" {
			break
		}

		num, err := strconv.Atoi(num1)

		if err != nil {
			fmt.Println("You have entered wrong input")
			break
		}

		slc = append(slc, num)

		sort.Slice(slc, func(i, j int) bool {
			return slc[i] < slc[j]
		})

		fmt.Println(slc)
	}

}
