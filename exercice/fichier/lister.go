package main

import "io/ioutil"
import "fmt"

func main() {
entries, err := ioutil.ReadDir("/home/user/Bureau/")

if err != nil {
    fmt.Println(err)
}

for i, entry := range entries {
    fmt.Println(entry.Name())    // Nom du fichier ("myphoto.jpg")
    fmt.Println(entry.Size())    // Taille en octet (/1024 = Ko)
    fmt.Println(entry.Mode())    // Droits d'écritures "-rw-rw-rw-"
    fmt.Println(entry.ModTime()) // Date de dernière modification
    fmt.Println(entry.IsDir())   // "false" par défaut (car on ne liste pas des "directories" / répertoires)
    fmt.Println(i)
}
}