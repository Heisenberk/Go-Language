package main

import (
    "fmt"
    "math/rand"
)
func alea(len int) []int {
    a := make([]int, len)
    for i := 0; i <= len-1; i++ {
        a[i] = rand.Intn(len)
    }
    return a
}
func meme_val() int{
	var tab []int =alea(99)
	var tmp int = 0
		for i := 0; i<99 ;i++ {
			for j := 0 ;j< 99 ;j++{
				if tab[i] == tab[j] {
				tmp++;
				i++	
				}
				
			}
		}

return tmp
}
func different_val() int {

return 99-meme_val()

}
func div_10() int {
var tab []int =alea(99)
	for i := 0 ; i< len(tab) ; i++{
		if tab[i] == len(tab) % 10 {
			return i
}
}
return -1
}
func main() {

	fmt.Println(div_10())	


}