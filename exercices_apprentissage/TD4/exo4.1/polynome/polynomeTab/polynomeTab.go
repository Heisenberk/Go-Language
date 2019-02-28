// Package main contient la définition de la structure PolynomeTab 
// et ses fonctions et méthodes. 
package polynomeTab

import "strconv"
import "errors"

// Constante représentant la taille maximale du polynome. 
const TAILLEMAX int = 50

const VIDE int = -1

// Structure PolynomeTab qui stocke un polynome sous forme d'un tableau. 
type PolynomeTab struct {
    Coeff []int 
    Degre int 
    init bool
}

// Fonction d'initialisation d'un polynome. 
func InitPolynomeTab(p *PolynomeTab){
	p.Coeff = make([]int, TAILLEMAX)
	for i:=0 ; i<TAILLEMAX ; i++ {
		p.Coeff[i]=0
	}
	p.Degre = VIDE
	p.init = true
}

// Méthode qui teste si le polynome a bien été initialisé. 
func (p PolynomeTab) TestInitPolynomeTab() bool {
	if p.init == false {
		return false
	}else {
		return true
	}
}

// Méthode permettant d'insérer un nouveau coefficient. 
func (p *PolynomeTab) InsertPolynomeTab (coef int, degre int) error{
	if p.TestInitPolynomeTab() == false {
		return errors.New("Insertion Impossible : PolynomeTab non initialisé. ")
	}
	if degre >= 50 {
		return errors.New("Insertion Impossible : Degré de l'insertion trop élevé. ")
	}
	if degre < 0 {
		return errors.New("Insertion Impossible : Degré de l'insertion négatif. ")
	}
	p.Coeff[degre] = coef;
	if (p.Degre < degre){
		p.Degre = degre
	} 
	return nil
}

// Méthode qui renvoie le degré maximal du polynome. 
func (p PolynomeTab) GetDegreMax() (int, error) {
	if p.TestInitPolynomeTab() == false {
		return -1, errors.New("Degre Invalide. ")
	}
	return p.Degre, nil
}

// Méthode qui affiche le polynome. 
func (p PolynomeTab) PrintPolynomeTab() string{
	//fmt.Println(strconv.Itoa(p.Degre))
	if p.TestInitPolynomeTab() == false {
		return "PolynomeTab non initialisé."
	}else if p.Degre == VIDE {
		return "0"
	}

	var str string = ""

	for i:=0 ; i<=p.Degre ; i++ {
		if p.Coeff[i] != 0 {
			str += strconv.Itoa(p.Coeff[i])
			if i != 0 {
				str += "X^"
				str += strconv.Itoa(i)
			}
			if i != p.Degre {
				str += " + "
			}
		}
	}

	return str
}

// Fonction qui additionne deux polynomes. 
func AddPolynomeTab(p1 PolynomeTab, p2 PolynomeTab) (PolynomeTab, error) {
	p := PolynomeTab{}
	if p1.TestInitPolynomeTab() == false {
		return p, errors.New("Addition impossible : p1 non initialisé. ")
	}else if p2.TestInitPolynomeTab() == false {
		return p, errors.New("Addition impossible : p2 non initialisé. ")
	}else if p1.Degre == VIDE {
		return p2, nil
	}else if p2.Degre == VIDE {
		return p1, nil
	}
	
    InitPolynomeTab(&p)
    for i:=0 ; i<TAILLEMAX ; i++ {
    	p.InsertPolynomeTab(p1.Coeff[i]+p2.Coeff[i], i)
    }
    if p1.Degre > p2.Degre {
    	p.Degre = p1.Degre
    }else {
    	p.Degre = p2.Degre
    }
    return p, nil
}

// Fonction qui multiplie deux polynomes. 
func MultPolynomeTab(p1 PolynomeTab, p2 PolynomeTab) (PolynomeTab, error) {
	p := PolynomeTab{}
	if p1.TestInitPolynomeTab() == false {
		return p, errors.New("Addition impossible : p1 non initialisé. ")
	}else if p2.TestInitPolynomeTab() == false {
		return p, errors.New("Addition impossible : p2 non initialisé. ")
	}else if p1.Degre == VIDE {
		return p2, nil
	}else if p2.Degre == VIDE {
		return p1, nil
	}

	InitPolynomeTab(&p)
	for i:=0 ; i<TAILLEMAX ;i++ {
		if p1.Coeff[i] != 0 {
			for j:=0 ; j<TAILLEMAX; j++ {
				if p2.Coeff[j] != 0 {
					p.Coeff[i+j] += p1.Coeff[i]*p2.Coeff[j]
				}
			}
		}
	}

	p.Degre = p1.Degre + p2.Degre
	return p, nil
}

