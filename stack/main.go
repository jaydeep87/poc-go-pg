package main

import (
	"fmt"
	// "sync"
)

func main(){
	var s []int

	var stackSize int

	fmt.Println("please enter size of stack")

	fmt.Scanln(&stackSize)


	for i:=1; i<=stackSize; i++ {

		s = append(s, i)
	}
	fmt.Println("stack :: ", s);

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }

    fmt.Println(s)
}