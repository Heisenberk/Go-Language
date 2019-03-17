# TD3 Exercice 2 : Solution 

Q1) On crée la fonction qui ouvre le fichier d'entrée, lit les données de ce fichier, puis écrit les données lues dans le fichier de copie : 

```go
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
```

Ensuite, on gère les commandes possibles de l'utilisateur dans la fonction main : 

```go
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
```

Par exemple, pour exécuter le programme sur le test fourni, on fait la commande suivante : `./copy examples/test1.txt examples/test2.txt`.

Q2) On va créer une fonction qui récupère le chemin choisi par l'utilisateur et qui va lister pour chaque élément de son contenu ses droits, sa taille, la date de sa dernière modification, le fait qu'il soit un dossier, et son nom : 

```go
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
```

On utilise la fonction main pour gérer les arguments de l'utilisateur : 

```go
// Fonction principale
func main() {
    arg := os.Args[1:]
    if len(arg) == 0 {
        fmt.Println("Aucun argument. ")
    }else {
        listing(arg[0])
    }
}
```
Par exemple, pour exécuter le programme sur le dossier qui contient la liste des différents TDs, on fait la commande suivante : `./listing ../../../`.


Pour avoir le code complet : [exo3.2](exo3.2.zip)