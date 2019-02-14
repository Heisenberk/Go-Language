package main

import "fmt"
import "container/list"

type test struct {
	val int

}

func list_init(n int) list.List {

var a list.List
fmt.Println(a)
a.Init()
var e list.Element
t := test{5}

e.Value=t

a.PushBack(e)
fmt.Println(a)
fmt.Println(e)
return a
}
func main() {
    fmt.Println()
      //var x list.List
        var a list.List
    //x.PushBack(1)
    //x.PushBack(2)
    //x.PushBack(3)
    a=list_init(4)
    //x.PushBackList(&a)
    fmt.Println(a)
    //for e := x.Front(); e != nil; e=e.Next() {
        //fmt.Println(e.Value.(int))
    //}
}