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