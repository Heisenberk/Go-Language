package main

import (

    "fmt"
    "io/ioutil"
    "os"
    //"syscall"
)
import "strings"

func cut(path string) string {

   


    path=path[:strings.LastIndexAny(path, "/")+1]


    return path
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func chemin(path string,name string,isdir bool){
        if(isdir){
        path =path+name+"/"

        lister(path)
        //fmt.Println(path)
        path=strings.TrimRight(path,"/")
        path=strings.TrimRight(path,name)
        //fmt.Println(path)
    }
}
func ecrire(path string,name string,isdir bool,Mode os.FileMode){
    fmt.Println(int(Mode),name)
      if(!isdir && (Mode==493 || Mode==420) ){
    d1 := []byte("tests")

    err := ioutil.WriteFile(path+name, d1, 0644)
    check(err)
    fmt.Println(path+name)
    }
    }
func lister(path string){
        entries, err := ioutil.ReadDir(path)
         //path=strings.TrimRight(path,"/")

    if err != nil {

        fmt.Println(err)
    }
    for i, entry := range entries {
        /*
    fmt.Println(entry.Name())    // Nom du fichier ("myphoto.jpg")
    fmt.Println(entry.Size())    // Taille en octet (/1024 = Ko)
    fmt.Println(entry.Mode())    // Droits d'écritures "-rw-rw-rw-"
    fmt.Println(entry.ModTime()) // Date de dernière modification
    fmt.Println(entry.IsDir())   // "false" par défaut (car on ne liste pas des "directories" / répertoires)
    */
    fmt.Println(i)
    //fmt.Println(path+entry.Name())
/*
    if(!entry.IsDir() ){
    d1 := []byte("hi")

    err := ioutil.WriteFile(path+entry.Name(), d1, 0644)
    check(err)
    fmt.Println(path+entry.Name())
    }
*/
    ecrire(path,entry.Name(),entry.IsDir(),entry.Mode())
    chemin(path,entry.Name(),entry.IsDir())

    /*
    if(entry.IsDir()){
        path =path+entry.Name()+"/"

        lister(path)
        //fmt.Println(path)
        path=strings.TrimRight(path,"/")
        path=strings.TrimRight(path,entry.Name())
        //fmt.Println(path)
    }
    */

}
}
/*
func lister_arbo(path string){

    //path := "/home/user/Bureau/app/"
    path=strings.TrimRight(path,"/")
    fmt.Println(cut(path))
    fmt.Println(strings.HasSuffix(path, "user"))
    for(strings.HasSuffix(path, "user")!=true){
    lister(path)

    path=cut(path)
    }
    fmt.Println(path)
}
*/
func main() {

    /*
    path := "/home/user/Bureau/app/"
     path=strings.TrimRight(path,"/")
    fmt.Println(cut(path))
    //fmt.Println(strings.LastIndexAny(cut(path), "/"))
    //path=path[:strings.LastIndexAny(cut(path), "/")]
    //path=path[:strings.LastIndexAny(cut(path), "/")]
    //path=path[:strings.LastIndexAny(cut(path), "/")]
    fmt.Println(strings.HasSuffix(path, "user"))
    for(strings.HasSuffix(path, "user")!=true){
    lister(path)
    path=cut(path)
    }
    fmt.Println(path)
    */
    path := "/home/user/Bureau/R/src/main/java/fr/uvsq/inf103/rogue_like/"
    /*
    path2 :="/home/user/Bureau/R/src/main/java/fr/uvsq/inf103/rogue_like/creature/Creature.java"
       d1 := []byte("hello")
    err := ioutil.WriteFile(path2, d1, 0644)
    check(err)
    */
    lister(path)
    //lister_arbo(path)


}