package main

import "fmt"


func main() {
n := 97
	for i := 2 ; i < n ; i++ {
		if n % i == 0{
			fmt.Println("n n'est pas premier")
			break
		}
	}
fmt.Println("n est premier")
}