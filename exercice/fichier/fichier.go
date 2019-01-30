package main
import "io/ioutil"
import "fmt"
//import "os"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    dat, err := ioutil.ReadFile("/home/user/Bureau/testbis")
    check(err)
    d1 := []byte(string(dat))
    err1 := ioutil.WriteFile("/home/user/Bureau/cqdf", d1, 0644)
    check(err1)
    fmt.Print(string(dat))
}