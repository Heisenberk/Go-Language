// Package main contient la fonction principale pour exécuter les fonctions 
// sur les polynomes. 
package main

import "fmt"
import "polynome/polynomeTab"

// main est la fonction Principale du programme.
func main() {
    p1 := polynomeTab.PolynomeTab{}
    polynomeTab.InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(2,0)
    _ = p1.InsertPolynomeTab(5,1)
    _ = p1.InsertPolynomeTab(2,3)

    p2 := polynomeTab.PolynomeTab{}
    polynomeTab.InitPolynomeTab(&p2)
    _ = p2.InsertPolynomeTab(3,0)
    _ = p2.InsertPolynomeTab(2,2)
    _ = p2.InsertPolynomeTab(2,4)

    p3, _ := polynomeTab.AddPolynomeTab(p1,p2)
    p4, _ := polynomeTab.MultPolynomeTab(p1,p2)

    fmt.Println(p1.PrintPolynomeTab())
    fmt.Println(p2.PrintPolynomeTab())
    fmt.Println(p3.PrintPolynomeTab())
    fmt.Println(p4.PrintPolynomeTab())
}