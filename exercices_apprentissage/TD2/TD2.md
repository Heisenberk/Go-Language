# TD2 : Manipulation de structures complexes

But : Apprendre à manipuler les tableaux, les slices et les maps. 

Difficulté : Intermédiaire

Connaissances à avoir : Structures, Tableaux, Slices, Maps.

Durée du TD : -----

## Exercice 1 : Tableau, Slice et Tri

1. Créer une fonction qui fait le tri par propagation (tri à bulles) d'un tableau en paramètre. 

2. Créer une fonction qui fait de manière récursive le tri fusion d'un tableau en paramètre. 

Solution : 

1. Le tri à bulles consiste à parcourir le tableau et échanger les valeurs consécutives mal triées du tableau jusqu'à ce que toutes les valeurs soient triées. Pour cela, il faut créer une fonction qui échange les valeurs du tableau, en plus de la fonction de tri à bulles : 

```go
package main

import "fmt"

// fonction qui échange les valeurs du tableau qui sont aux positions x et y
func  echange(tab []int, x int, y int)  {
	tmp := tab[x]
	tab[x]= tab[y]
	tab[y]=tmp
}

// pour chaque couple de valeurs, si elles sont dans le mauvais ordre, on les échange
func bulle (tab []int) []int {
	for i:=0 ;i<len(tab) ; i++ {
		for j:=i+1 ;j<len(tab) ; j++ {
			if tab[j]< tab[i]{
				echange(tab , i , j)
			}
		}
	}
	return tab
}

func main() {
	tab := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Tableau avant le tri : ", tab)
	bulle(tab)
	fmt.Println("Tableau après le tri : ", tab)
}
```

2. Le tri fusion consiste à appeler de façon récursive au tri fusion `triFusion` pour la première moitié du tableau ainsi que la deuxième. Ensuite, il faut fusionner `fusionner` ces deux parties triées : 

```go
package main

import "fmt"

func fusionner(gauche []int, droite []int) []int {

  taille := len(gauche) + len(droite)
  var tab = make([]int, taille)
   i := 0
   j := 0

  for i < len(gauche) && j < len(droite) {
    if gauche[i] <= droite[j] {
      tab[i+j] = gauche[i]
      i++
    } else {
      tab[i+j] = droite[j]
      j++
    }
  }
  
  for i < len(gauche) { 
    tab[i+j] = gauche[i] 
    i++ 
  }

  for j < len(droite) {
   tab[i+j] = droite[j]
   j++ 
 }

  return tab
}



func triFusion(tab []int) []int {

  if len(tab) < 2 {
    return tab
  }

  var milieu = len(tab) / 2
  gauche := triFusion(tab[:milieu])
  droite := triFusion(tab[milieu:])

  return fusionner(gauche, droite)
}

func main () {

  tab := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
  fmt.Println("Tableau avant le tri : ", tab)
  result := triFusion(tab)
  fmt.Println("Tableau après le tri : ", result)

}
```

Il est important de noter ici que l'on a utilisé une slice pour le tableau fusionné dans la fonction `fusionner`. En effet, il n'est pas possible de déclarer un tableau avec une taille variable (en fonction de l'entrée de la fonction), d'où l'utilisation d'une slice `var tab = make([]int, taille)`. 

Pour avoir le code complet : [exo2.1](exo2.1.zip)

## Exercice 2

On souhaite implémenter un annuaire électronique, qui donne pour chaque nom de contact ses coordonnées (numéro de téléphone et adresse). 

1. Définir la structure Coordonnées. 

2. Définir la structure Annuaire qui gère l'annuaire électronique. 

3. Définir les fonctions suivantes : l'ajout d'un nouveau contact, la modification des coordonnées d'un contact, la suppression d'un contact et l'affichage de la liste des contacts de l'annuaire. 

Solution : 

1. La structure `Coordonnees` comporte 2 champs : `telephone` et `adresse`. On crée également une fonction pour changer les valeurs des coordonnées et une fonction d'affichage. 

```go
// Structure Coordonnees
type Coordonnees struct {
  telephone int
  adresse string
}

// Change le telephone et l'adresse de la coordonnee
func (c Coordonnees) changeCoordonnees(nouveauTelephone int, nouvelleAdresse string) Coordonnees {
  c.adresse=nouvelleAdresse
  c.telephone=nouveauTelephone
  return c
}

// Affiche les coordonnees 
func (c Coordonnees) afficheCoordonnees() string{
  return "Tel: " + strconv.Itoa(c.telephone) + "; Adresse: " + c.adresse;
}
```

2. La structure `Annuaire` correspond à une Map dont le nom du contact est la clé. 

```go
// Structure Annuaire
type Annuaire struct {
  contacts map[string]Coordonnees
}

// Initialise l'annuaire
func initialiseAnnuaire(annuaire *Annuaire) {
  annuaire.contacts = make(map[string]Coordonnees)
}
```

3. 

```go
// Ajout d'un nouveau contact à l'annuaire
// Lance une exception si l'annuaire n'a pas été initialisé
func (a Annuaire) ajouteContact(nom string, coordonnees Coordonnees){
  if(a.contacts == nil){
    panic(errors.New("L'annuaire doit être initialisé"))
  }
  a.contacts[nom] = coordonnees
}

// Affichage de la liste des contacts de l'annuaire
func (a Annuaire) afficheAnnuaire() {
  i := 0
  for nom, coordonnees := range a.contacts {
    i++
      fmt.Println("Contact ", i, "--> Nom: ", nom, coordonnees.afficheCoordonnees())
  }
}

// Modification des coordonnées d'un contact de l'annuaire
func (a Annuaire) modifieCoordonneesContact (nom string, nouvellesCoordonnees Coordonnees) {
  a.contacts[nom]=nouvellesCoordonnees
} 

// Suppression d'un contact 
func (a Annuaire) supprimeContact (nom string) {
  delete(a.contacts, nom)
}
```

A noter que `ajouteContact` nécessite que la Map soit initialisée : nous avons donc utilisé une exception pour arrêter le programme en cas de problème. 

Pour avoir le code complet : [exo2.2](exo2.2.zip)