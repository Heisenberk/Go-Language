package main

import "fmt"

type Element struct {
	val int
	suiv  *Element 
}


func list_init(n int) *Element {
	v := Element{n,nil}
	list := &v
return list
}

func add_element_debut(i int,list *Element) *Element {
	v := Element{i,list}

	return &v
}
func add_element_fin(n int,list *Element) *Element {
	elem := Element{n,nil}
	list_save :=list
		    for  list.suiv!=nil {
       		 list=list.suiv
    }

    list.suiv=&elem
    return list_save
}	

func affichage(list *Element) {
		    for e := list; e != nil; e=e.suiv {
       		 fmt.Println(e.val)
    }
}
func supprimer_un_element(n int, list *Element) *Element{
	debut := list
		compteur := 0

			    for  list.suiv.suiv != nil{
			    	if compteur == n-1 {
       		 			tmp := list.suiv
       		 			fmt.Println("suppresion de ",tmp.val)
       		 			list.suiv= list.suiv.suiv
       		 			tmp = nil
       		 		}
       		 		compteur++
       		 		list=list.suiv
    }
	return debut
}

func main() {
	//v := Element{5,nil}
	//t := Element{3,nil}
	list := list_init(1)

	list=add_element_debut(2,list)
	list=add_element_debut(3,list)
	//list=add_element_debut(4,list)
	//list=add_element_debut(5,list)
	//list=add_element_debut(6,list)
	//list = add_element_fin(0,list)
	list= supprimer_un_element(1,list)
		affichage(list)
		//fmt.Println(list)
	//affichage(list)
	//fmt.Println(list)
	//fmt.Println(list.suiv)
}