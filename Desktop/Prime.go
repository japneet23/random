package main

import (
	"fmt"
)

func main() {
	n := 19
	i := 0
	for i = 2; i <= n/2; i++ {
		if n%i == 0 {
			fmt.Println("Not a Prime number")
			break
		}
	
	else {
		fmt.Println("Its a Prime number")
	}
}
	
}
