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
    
### Types 

mettre un tableau des differents types importants A REMPLIR

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

