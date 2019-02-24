package main

import "fmt"

type Element struct {
    suiv *Element
    val int 
}
 
type Liste struct {
    deb *Element
}

func (l *Liste) compte_elem() int {
    if l.deb == nil {
        return 0
    }else{
        comptage := 1
        cpy := l.deb
        for cpy.suiv != nil {
            cpy = cpy.suiv
            comptage ++
        }
        return comptage
    }
}

func (l *Liste) insere_deb(val int) {
    nouveau := &Element{
        suiv: l.deb,
        val:  val,
    }
    l.deb = nouveau
}

func (l *Liste) insere_fin(val int) {
    ancienne := l.deb
    nouveau := &Element{
        suiv: nil,
        val:  val,
    }
    for ancienne.suiv != nil {
        ancienne = ancienne.suiv
    }
    ancienne.suiv=nouveau
}

func (l *Liste) supprime_deb(){
    comptage := l.compte_elem()
    if comptage == 0 {
    }else if comptage == 1 {
        l.deb = nil
    }else {
        ancienne := l.deb
        l.deb = ancienne.suiv
        ancienne.suiv = nil
    }
}

func (l *Liste) supprime_fin() {
    comptage := l.compte_elem()
    if comptage == 0 {
    }else if comptage == 1 {
        l.deb = nil
    }else {
        ancienne := l.deb
        for ancienne.suiv.suiv != nil {
            ancienne = ancienne.suiv
        }
        ancienne.suiv = nil
    }
}

func (l *Liste) affiche() {
    if l.deb == nil {
        fmt.Println("[Liste Vide]")
    }else{
        list := l.deb
        fmt.Print("[")
        for list.suiv != nil {
            fmt.Printf("%d -> ", list.val)
            list = list.suiv
        }
        fmt.Printf("%d]\n", list.val)
    }
}

func main() {
    link := Liste{}
    link.affiche() // [Liste Vide]
    link.insere_deb(2)
    link.affiche() 
    link.insere_deb(1)
    link.affiche() // [1 -> 2]
    link.insere_fin(3)
    link.affiche() // [1 -> 2 -> 3]
    link.supprime_deb()
    link.affiche() // [2 -> 3]
    link.supprime_fin()
    link.affiche() // [2]
    link.supprime_deb()
    link.affiche() // [Liste Vide]
}


