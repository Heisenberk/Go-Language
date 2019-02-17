package main

import "fmt"
import "errors"
import "strconv"

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


/////////////////////////////////////////////////////////////////////////

// Structure Annuaire
type Annuaire struct {
	contacts map[string]Coordonnees
}

// Initialise l'annuaire
func initialiseAnnuaire(annuaire *Annuaire) {
	annuaire.contacts = make(map[string]Coordonnees)
}

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

func main() {

	var annuaire Annuaire
	initialiseAnnuaire(&annuaire)

	annuaire.ajouteContact("Jean Dupont", Coordonnees{telephone: 0130405060, adresse: "12 av. des tulipes Versailles"})
	annuaire.ajouteContact("Jeanne Martin", Coordonnees{telephone: 0130405010, adresse: "1 rue des fraises Chenay"})

	annuaire.afficheAnnuaire()
	fmt.Println();

	annuaire.modifieCoordonneesContact("Jean Dupont", Coordonnees{telephone: 0230707070, adresse: "12 av. des tulipes Paris"})
	annuaire.afficheAnnuaire()
	fmt.Println();

	annuaire.supprimeContact ("Jean Dupont")
	annuaire.afficheAnnuaire()
	fmt.Println();
}