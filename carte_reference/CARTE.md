# Carte de référence du langage Go

## Bases du langage 



[My page](coucou.md)

### Variable

Une variable permet de déclarer une variable modifiable à tout moment. 
    
```go
var a int = 1
/* ou directement */
a := 1 
```

### Constante 

Une constante permet de déclarer une constante (non modifiable).
    
```go
const pi float32 = 3.14
```

### Type

Il existe plusieurs types en langage Go : 
- le type booléen `bool` prend 2 valeurs possibles : `true` ou `false`. 
- le type `string` représentant une chaîne de caractères. 
- le type entier non signé : `uint`, `uint8` (ou `byte` pour représenter l'octet), `uint16`, `uint32` et `uint64`.
- le type entier signé : `int`, `int8`, `int16`, `int32` (ou `rune` pour représenter un caractère unicode) et `int64`. 
- le type réel : `float32` et `float64`. 
- le type complexe : `complex64` et `complex128`

Les manipulations arithmétiques (`+`, `-`, `*`, `/`, `%`, `++` et `--`), relationnelles (`==`, `!=`, `>`, `<`, `>=` et `<=`) et binaires (`<<`, `>>`, `&`, `|`, `^`) fonctionnent en Go. 



### import

Le mot clé `import` permet d'importer des fichiers tels que les bibliotheques standards (ex : `fmt`).
    
```go
import "fmt"
```
    

### Condition 

Une condition permet de faire des instructions si la condition est vraie (`if`). On peut effectuer des instructions si la condition est fausse (`else`). Il est possible de faire une succession de tests avec `if/else if/.../else`.
    
```go
var a int = 3
if a < 0{
    fmt.Println(a, "negatif")
} else if a > 0 {
    fmt.Println(a, "positif")
} else {
    fmt.Println(a, "nul")
}
```

### Boucle 

Une boucle se fait avec le mot clé `for`. 
   
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

/* ou bien */

i := 0 
for i < 10 {
    fmt.Println(i)
    i = i + 1
}
```

### Break

`break` permet de sortir de la boucle dans laquelle se trouve l'instruction.
    
```go
for i:=0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
```

### Switch

`switch` permet de faire une succession de cas possibles au lieu de plusieurs tests successifs.
    
```go
a := 1
switch a%2 {
    case 0:
        fmt.Println("pair")
    case 1:
        fmt.Println("impair")
}
```

### Tableau  

Un tableau permet de déclarer un ensemble fini (à une ou plusieurs dimensions) et ordonné de variables ayant le même type.
    
```go
var b [2] int
b[0]=0
b[1]=1

/* ou directement */ 

b := [2]int{0, 1}
```

### Slice 

`make` permet de déclarer une slice (tableau dynamique) à une ou plusieurs dimensions.
    
```go
s := make([]string, 2)
s[0] = "a"
s[1] = "b"

/* ou directement */ 

s := []string{"a", "b"}
```

### len()

`len()` permet de récupérer la taille d'un tableau ou d'une slice.
    
```go
var a [2] int
fmt.Println(len(a))
```

### append()

`append()` permet d'ajouter un élément à la slice.
    
```go
s := make([]string, 2)
s[0] = "a"
s[1] = "b"
s = append(s, "c") 
```

### copy()

`copy()` permet de copier une slice dans une autre.
    
```go
s := make([]string, 2)
s[0] = "a"
s[1] = "b"
c := make([]string, len(s))
copy(c, s) 
```

### Map 

Une map permet de stocker des valeurs en leur associant une clé.
    
```go
m := make(map[string]int)
m["mehdi"] = 1
m["clement"] = 2
```

### delete() 

`delete()` permet de supprimer une valeur dans la map.
    
```go
m := make(map[string]int)
m["mehdi"] = 1
m["clement"] = 2
delete(m, "clement")
```

### identificateur _

Cet identificateur permet de déclarer une variable permettant de tester la présence d'une clé (true ou false).
    
```go
m := make(map[string]int)
m["mehdi"] = 1
m["clement"] = 2
_, test1 := m["clement"] // test1 vaut true
_, test2 := m["inconnu"] // test2 vaut false
```

### range

`range` permet d'itérer facilement dans une structure de tableau, de slice ou de map.
    
```go
tab := []int{1, 2, 3}
somme := 0
for _, val := range tab {
    somme += val
}
fmt.Println("Somme:", somme)
```

### Fonction 

Déclarer une fonction se fait avec le mot clé `func`. Une fonction peut avoir plusieurs retours.
    
```go
func calcul1 (val1 int, val2 int) int {
    return val1 + val2
}

func calcul2 (val1 int, val2 int) (int, int){
    return val1 + val2, val1 - val2
}

func main() {
    var test1, test2, test3 int = 0, 0, 0
    test1, test2 = calcul2(2,3)
    test3 = calcul1(1,2)
}
```

### Pointeur 

Un pointeur permet de faire des références sur des variables. Les pointeurs Go s'utilisent comme en langage C.
    
```go
func add (val *int){
    *val = *val + 1
}

func main() {
    var cal int = 2
    add(&cal)
    fmt.Println(cal)
}
```

### Structure 

Déclarer une structure se fait avec le mot clé `struct`. Une structure est un nouveau type de variables possédant plusieurs champs.
    
```go
type etudiant struct {
    nom string
    num  int
}

func main() {
    e := etudiant{nom: "Clement", num: 21501810}
    fmt.Println(e)
}
```

### Méthode 

Une méthode correspond à une fonction déclarée pour une structure (hors de la déclaration de la structure).
    
```go
func (e etudiant) change_num (nouveau int) etudiant{
    e.num=nouveau
    return e
}

func main() {
    e := etudiant{nom: "Clement", num: 21501810}
    e = e.change_num(21500000)
    fmt.Println(e)
}
```

### Interface 

Une interface permet de regrouper des signatures de méthodes communes à plusieurs structures.
    
```go
type politesse interface {
    saluer() string
}

type francais struct {
    nom string
}

func (f francais) saluer() string {
    return "bonjour"
}

type anglais struct {
    name string
}

func (a anglais) saluer() string{
    return "hello"
}

func parler(p politesse) {
    fmt.Println(p.saluer())
}

func main() {
    angl := anglais{name: "James"}
    fran := francais{nom: "Jean"}
    parler(angl) // affiche "hello"
    parler(fran) // affiche "bonjour"
}
```

### Gestion d'erreurs 

Les erreurs (`errors`) permettent de gérer les exceptions en langage Go.
    
```go
func quotient (num float32, denom float32) (float32, error){
    if denom == 0 {
        return -1, errors.New("Erreur mathematique")
    }else {
        return num / denom, nil
    }
}

func main() {
    var c1 float32
    var err error
    var a1, b1 float32 = 1,2
    c1,err = quotient(a1, b1)
    fmt.Println(c1)
    fmt.Println(err)
}
```

## Bibliothèques du langage Go

### Principales bibliothèques standards

#### Bibliothèque `fmt`

`fmt` est la bibliothèque permettant de gérer les entrées-sorties de l'utilisateur. En langage C, on peut donc la comparer à la bibliothèque `stdio.h`. 
On peut utiliser la fonction `Printf` pour faire un affichage "formaté" et `Scanf` pour récupérer l'entrée de l'utilisateur (le clavier par défaut).  

```go
package main

import "fmt"

func main() {
    test := 2
    fmt.Printf("test = %d \n", test)
    var entree int
    fmt.Scanf("%d", &entree)
    fmt.Printf("entree : %d\n", entree)
    /* 
        affiche "test = 2"
        puis attend l'entrée de l'utilisateur
        Si l'utilisateur écrit "3" alors l'affichage sera "entree : 3"
    */
}
```
Pour plus d'informations, notamment la gestion des formats ainsi que les retours d'erreurs : [Golang-fmt](https://golang.org/pkg/fmt/). 

#### Bibliothèque `error`

#### Bibliothèque `io`

`io` est la bibliothèque permettant notamment d'intéragir avec les fichiers. Il est donc possible de lire et d'écrire dans des fichiers. 

Dans cet exemple, on va lire en entier le fichier `test1.txt` et recopier son contenu dans le fichier `test1_copie.txt`. A noter que si le fichier `test1_copie.txt` n'existe pas, il le crée ; sinon, il l'écrase. 

```go
package main

import "io/ioutil"

func main(){

    donnees_copiees, err1 := ioutil.ReadFile("test1.txt")
    if err1 != nil {
        panic(err1)
    }

    donnees_a_recopier := []byte(string(donnees_copiees))
    err2 := ioutil.WriteFile("test1_copie.txt", donnees_a_recopier, 0644) // 0644 correspond aux permissions.
    if err2 != nil {
        panic(err2)
    }

}
```
Pour plus d'informations, notamment la gestion des formats ainsi que les retours d'erreurs : [Golang-io](https://golang.org/pkg/io/) et [Golang-io/ioutil](https://golang.org/pkg/io/ioutil/). 


#### Bibliothèque `os`

#### Bibliothèque `strings`

#### Bibliothèque `time`

### Principales bibliothèques tierces

## Outils de développement

## Ressources d'apprentissage

- [Préambule pour découvrir le langage Go](https://fr.wikipedia.org/wiki/Go_(langage))
- [Site officiel de Go](https://golang.org/)
- [Bases du langage Go](https://www.tutorialspoint.com/go/index.htm)
- [Illustrations de l'utilisation des bases du langage Go](https://gobyexample.com/)
- [Introduction de la concurrence en langage Go](https://blog.fedora-fr.org/metal3d/post/Go-et-les-goroutines-introduction-au-langage)
- [Go vs C++](https://www.scriptol.fr/programmation/go.php)
- [Application smartphone pour apprendre le langage Go](https://play.google.com/store/apps/details?id=in.intelitech.golang&hl=en_US)

