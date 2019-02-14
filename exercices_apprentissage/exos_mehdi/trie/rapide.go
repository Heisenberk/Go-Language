package main



import "fmt"



func rapide(left []int, right []int) []int {

  taille :=len(left)+len(right)

  var tab = make([]int, taille)

   i := 0

   j := 0

  for i < len(left) && j < len(right) {

  

    if left[i] <= right[j] {

      tab[i+j] = left[i]

      i++

    } else {

      tab[i+j] = right[j]

      j++

    }

    

  }
  
  for i < len(left) { tab[i+j] = left[i]; i++ }

  for j < len(right) { tab[i+j] = right[j]; j++ }

  return tab
}



func Tri_rapide(tab []int) []int {



  if len(tab) < 2 {

    return tab

  }

  var millieu = len(tab) / 2

  left := Tri_rapide(tab[:millieu])

  right := Tri_rapide(tab[millieu:])

  return rapide(left, right)

}

func main () {



  fmt.Print(Tri_rapide([]int{ 3, 9, 6, 4, 5, 8, 7, 10, 2, 1 }), "\n")

  

}