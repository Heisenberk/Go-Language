package main

import "io/ioutil"
import "fmt"
import "os"

// Fonction qui copie path1 dans path2
func copy (path1 string, path2 string) int {
    // Ouverture du fichier path1 en lecture
    file1, err1 := ioutil.ReadFile(path1)
    if err1 != nil {
        return 1
    }

    // Enregistrement des données de path1
    donnees := []byte(string(file1))

    // Ecriture des données enregistrées dans path2
    err2 := ioutil.WriteFile(path2, donnees, 0644)
    if err2 != nil {
        return 2
    }
    return 0
}

// Fonction principale
func main() {
    arg := os.Args[1:]
    if len(arg) == 0 {
        fmt.Println("Manque 2 arguments. ")
    }else if len(arg) == 1{
        fmt.Println("Manque 1 argument. ")
    }else {
        err := copy(arg[0], arg[1])
        if err == 1 {
            fmt.Println("Problème de lecture de " + arg[0] + ". ")
        }else if err == 2 {
            fmt.Println("Problème d'écriture dans " + arg[1] + ". ")
        }else {
            fmt.Println("Copie de \"" + arg[0] + "\" dans \"" + arg[1] + "\" terminée. ")
        }
    }
}