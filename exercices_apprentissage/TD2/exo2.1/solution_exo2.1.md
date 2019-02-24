# TD2 Exercice 1 : Solution 

1. Le tri à bulles consiste à parcourir le tableau et échanger les valeurs consécutives mal triées du tableau jusqu'à ce que toutes les valeurs soient triées. Pour cela, il faut créer une fonction qui échange les valeurs du tableau, en plus de la fonction de tri à bulles : 

```go
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
```

2. Le tri fusion consiste à appeler de façon récursive au tri fusion `triFusion` pour la première moitié du tableau ainsi que la deuxième. Ensuite, il faut fusionner `fusionner` ces deux parties triées : 

```go
package main

import "fmt"

func fusionner(gauche []int, droite []int) []int {

  taille := len(gauche) + len(droite)
  var tab = make([]int, taille)
   i := 0
   j := 0

  for i < len(gauche) && j < len(droite) {
    if gauche[i] <= droite[j] {
      tab[i+j] = gauche[i]
      i++
    } else {
      tab[i+j] = droite[j]
      j++
    }
  }
  
  for i < len(gauche) { 
    tab[i+j] = gauche[i] 
    i++ 
  }

  for j < len(droite) {
   tab[i+j] = droite[j]
   j++ 
 }

  return tab
}



func triFusion(tab []int) []int {

  if len(tab) < 2 {
    return tab
  }

  var milieu = len(tab) / 2
  gauche := triFusion(tab[:milieu])
  droite := triFusion(tab[milieu:])

  return fusionner(gauche, droite)
}

func main () {

  tab := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
  fmt.Println("Tableau avant le tri : ", tab)
  result := triFusion(tab)
  fmt.Println("Tableau après le tri : ", result)

}
```

Il est important de noter ici que l'on a utilisé une slice pour le tableau fusionné dans la fonction `fusionner`. En effet, il n'est pas possible de déclarer un tableau avec une taille variable (en fonction de l'entrée de la fonction), d'où l'utilisation d'une slice `var tab = make([]int, taille)`. 

Pour avoir le code complet : [exo2.1](exo2.1.zip)