package main

import "io/ioutil"
import "fmt"
import "os"

// Fonction qui liste les éléments d'un chemin en paramètre
func listing(path string){
    elements, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Println(err)
    }

    // chaque élément du chemin est analysé
    for i, input := range elements {
        fmt.Print(i)
        fmt.Print(". ")

        fmt.Print(input.Mode())    // Droits de lecture et d'écriture
        fmt.Print(" ")
    
        fmt.Print(input.Size())    // Taille en octets (/1024 = Ko)
        fmt.Print(" ")
    
        fmt.Print(input.ModTime()) // Date de dernière modification
        test_dir := input.IsDir()   // "true" si c'est un dossier et "false" sinon.
        if test_dir == true {
            fmt.Print(" d ")
        }else {
            fmt.Print(" f ")
        }
        fmt.Println(input.Name())    // Nom du fichier en question
    }
}

// Fonction principale
func main() {
    arg := os.Args[1:]
    if len(arg) == 0 {
        fmt.Println("Aucun argument. ")
    }else {
        listing(arg[0])
    }
}