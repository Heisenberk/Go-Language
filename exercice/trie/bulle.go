package main

import "fmt"

func  echange(tab []int,x int,y int)  {
	tmp := tab[x]
	tab[x]= tab[y]
	tab[y]=tmp
}
func bulle (tab []int) []int {

	for i:=0 ;i<len(tab) ; i++ {
		for j:=i+1 ;j<len(tab) ; j++ {
			if tab[j]< tab[i]{
				echange(tab , i , j)
			}
			
		}
	}
	return tab
}

func main() {

	tab := []int{5 ,1 , 2, 6, 4, 3 }
	fmt.Print(tab,"----->")
	fmt.Println(bulle(tab))




	
	
   
}