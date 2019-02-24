# TD2 Exercice 2 : Solution 

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