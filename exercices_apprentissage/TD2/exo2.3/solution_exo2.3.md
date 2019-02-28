# TD2 Exercice 3 : Solution 

Q1) Pour définir une liste, il va falloir définir deux structures : l'élément `Element` qui sera inséré dans la liste et la liste elle-même `Liste`. 

```go
type Element struct {
    suiv *Element
    val int 
}
 
type Liste struct {
    deb *Element
}
```

Un `Element` représente un entier. il aura donc deux champs : un champ `val` représentant cet entier et un champ `suiv` qui sera un pointeur vers l'élément suivant. 
Sachant que chaque élément pointe vers le suivant, il suffit de stocker dans `Liste` le pointeur du premier élément.

Q2) Avant de commencer à implémenter les méthodes demandées, il est utile de compter les éléments d'une liste : 

```go
func (l *Liste) compte_elem() int {
    if l.deb == nil {
        return 0
    }else{
        comptage := 1
        cpy := l.deb
        for cpy.suiv != nil {
            cpy = cpy.suiv
            comptage ++
        }
        return comptage
    }
}
```

Cela permettra de gérer correctement les cas possibles pour la suppression et l'affichage notamment. 
Les méthodes d'insertion sont les suivantes : 

```go
func (l *Liste) insere_deb(val int) {
    nouveau := &Element{
        suiv: l.deb,
        val:  val,
    }
    l.deb = nouveau
}

func (l *Liste) insere_fin(val int) {
    ancienne := l.deb
    nouveau := &Element{
        suiv: nil,
        val:  val,
    }
    for ancienne.suiv != nil {
        ancienne = ancienne.suiv
    }
    ancienne.suiv=nouveau
}
```
Il est important de noter qu'il est nécessaire de créer un nouvel élément `nouveau` puis de l'insérer. 

Les méthodes de suppression sont les suivantes : 

```go
func (l *Liste) supprime_deb(){
	comptage := l.compte_elem()
    if comptage == 0 {
    }else if comptage == 1 {
        l.deb = nil
    }else {
        ancienne := l.deb
        l.deb = ancienne.suiv
        ancienne.suiv = nil
    }
}

func (l *Liste) supprime_fin() {
    comptage := l.compte_elem()
    if comptage == 0 {
    }else if comptage == 1 {
        l.deb = nil
    }else {
        ancienne := l.deb
        for ancienne.suiv.suiv != nil {
            ancienne = ancienne.suiv
        }
        ancienne.suiv = nil
    }
}
```

Il existe 3 cas possibles : un cas où la liste est vide et on ne fait rien, un cas où la liste contient 1 élément et on supprime l'élément au début de celle-ci, et un cas où la liste possède plusieurs éléments et il faut supprimer l'élément de début ou de fin. 

La méthode d'affichage est la suivante : 

```go
func (l *Liste) affiche() {
    if l.deb == nil {
        fmt.Println("[Liste Vide]")
    }else{
        list := l.deb
        fmt.Print("[")
        for list.suiv != nil {
            fmt.Printf("%d -> ", list.val)
            list = list.suiv
        }
        fmt.Printf("%d]\n", list.val)
    }
}
```

Pour avoir le code complet : [exo2.3](exo2.3.zip)

