package main

//import "errors"

/*func main() {
	contenu, erreur := ioutil.ReadFile("copy.go")
	if erreur != nil {
		fmt.Println("Erreur EOF")
	}

	fmt.Printf("%s", contenu)

}*/

import "fmt"
//import "io/ioutil"
import "os"

func check(e error){
    if e != nil {
        panic(e) // fonction pour arrêter le programme
    }
}

func main(){
	/*donnees, erreur := ioutil.ReadFile("test1.txt") // on lit le fichier jusqu'à la fin
	check(erreur)
	fmt.Print(string(donnees)) // */


	//LIRE UN FICHIER EN ENTIER 

	// OU LIRE PETIT A PETIT

	 f, err := os.Open("test2.txt")
    check(err)

    b1 := make([]byte, 124)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))
    for i := 0; i < n1; i++ {
		fmt.Printf("[%c,%d]", b1[i], b1[i])
	}
	if b1[n1] == -1 {
		fmt.Println("EOF")
	}
	//fmt.Println(b1[n1])


	/*f, err := os.Open("test.txt")
    check(err)*/

    /*b1 := make([]byte, 1)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))*/

    /*var car byte
	for car != EOF {
		n1, err := f.Read(car)
    	check(err)
    	fmt.Printf("%d bytes: %s\n", n1, car)
	}*/

    /*o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))*/
}