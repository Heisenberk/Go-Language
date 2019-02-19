package main

import "fmt"

// fonction qui échange les valeurs du tableau qui sont aux positions x et y
func  echange(tab []int, x int, y int)  {
	tmp := tab[x]
	tab[x]= tab[y]
	tab[y]=tmp
}

// pour chaque couple de valeurs, si elles sont dans le mauvais ordre, on les échange
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
	tab := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Tableau avant le tri : ", tab)
	bulle(tab)
	fmt.Println("Tableau après le tri : ", tab)
}