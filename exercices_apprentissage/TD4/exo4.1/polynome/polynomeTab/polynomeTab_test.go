// Package main contient la définition de la structure PolynomeTab 
// et ses fonctions et méthodes. 
package polynomeTab

import "testing"

// Teste si un polynomeTab n'a pas été initialisé. 
func TestInitPolynomeTabFaux(t *testing.T) {
    p1 := PolynomeTab{}

    calcul1 := p1.TestInitPolynomeTab()
    result1 := false

    if result1 != calcul1 {
        t.Errorf("erreur dans TestInitPolynomeTabFaux")
    }
}

// Teste si un polynomeTab a bien été initialisé. 
func TestInitPolynomeTabVrai(t *testing.T) {
    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)

    calcul2 := p2.TestInitPolynomeTab()
    result2 := true

    if result2 != calcul2 {
        t.Errorf("erreur dans TestInitPolynomeTabVrai")
    }
}

// Teste si l'insertion dans un polynome non initialisé échoue. 
func TestInsertPolynomeTabErreur1(t *testing.T) {
    p2 := PolynomeTab{}
    errCalcul := p2.InsertPolynomeTab(2,1)

    if errCalcul == nil { // doit lancer une erreur
        t.Errorf("erreur dans TestInsertPolynomeTabErreur1")
    }
}

// Teste si l'insertion d'un trop grand degré échoue. 
func TestInsertPolynomeTabErreur2(t *testing.T) {
    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    errCalcul := p2.InsertPolynomeTab(1,60)

    if errCalcul == nil { // doit lancer une erreur
        t.Errorf("erreur dans TestInsertPolynomeTabErreur2")
    }
}

// Teste si l'insertion d'un degré négatif échoue. 
func TestInsertPolynomeTabErreur3(t *testing.T) {
    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    errCalcul := p2.InsertPolynomeTab(1,-2)

    if errCalcul == nil { // doit lancer une erreur
        t.Errorf("erreur dans TestInsertPolynomeTabErreur3")
    }
}

// Teste si le degré du polynome est bien égal à 2
func TestGetDegreMax1(t *testing.T){
    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    p2.InsertPolynomeTab(1,2)
    degre,_ := p2.GetDegreMax()
    if degre != 2 { 
        t.Errorf("erreur dans TestGetDegreMax1")
    }
}

// Teste si la non initialisation provoque bien une erreur. 
func TestGetDegreMax2(t *testing.T){
    p2 := PolynomeTab{}
    err := p2.InsertPolynomeTab(1,2)
    degre,_ := p2.GetDegreMax()

    if degre != -1 && err == nil{ 
        t.Errorf("erreur dans TestGetDegreMax2")
    }
}

// Teste si le degre d'un polynome sans valeur est VIDE.
func TestGetDegreMax3(t *testing.T) {
    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    degre,_ := p2.GetDegreMax()
    if degre != VIDE { // doit lancer une erreur
        t.Errorf("erreur dans TestGetDegreMax3")
    }
}

// Teste l'affichage d'un polynome non initialisé. 
func TestPrintPolynomeTab1(t *testing.T){
    p2 := PolynomeTab{}
    str := p2.PrintPolynomeTab()
    if str != "PolynomeTab non initialisé." {
        t.Errorf("erreur dans TestPrintPolynomeTab1")
    }
}

// Teste l'affichage d'un polynome vide.
func TestPrintPolynomeTab2(t *testing.T){
    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    str := p2.PrintPolynomeTab()
    if str != "0" {
        t.Errorf("erreur dans TestPrintPolynomeTab2")
    }
}

// Teste si l'affichage d'un polynome. 
func TestPrintPolynomeTab3(t *testing.T){
    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    _ = p2.InsertPolynomeTab(1,0)
    _ = p2.InsertPolynomeTab(2,3)
    _ = p2.InsertPolynomeTab(4,8)

    str := p2.PrintPolynomeTab()
    if str != "1 + 2X^3 + 4X^8" {
        t.Errorf("erreur dans TestPrintPolynomeTab3")
    }
}

// Teste si l'addition de deux polynomes fonctionne. 
func TestAddPolynomeTab1(t *testing.T){
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(3,0)
    _ = p1.InsertPolynomeTab(1,2)
    _ = p1.InsertPolynomeTab(1,3)
    _ = p1.InsertPolynomeTab(4,7)
    _ = p1.InsertPolynomeTab(4,8)

    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    _ = p2.InsertPolynomeTab(1,0)
    _ = p2.InsertPolynomeTab(2,3)
    _ = p2.InsertPolynomeTab(4,8)

    p3, _ := AddPolynomeTab(p1,p2)

    str := p3.PrintPolynomeTab() 
    if str != "4 + 1X^2 + 3X^3 + 4X^7 + 8X^8"{
        t.Errorf("erreur dans TestAddPolynomeTab1")
    }
}

// Teste si l'addition échoue avec un paramètre non initialisé. 
func TestAddPolynomeTab2(t *testing.T){
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(3,0)
    _ = p1.InsertPolynomeTab(1,2)
    _ = p1.InsertPolynomeTab(1,3)
    _ = p1.InsertPolynomeTab(4,7)
    _ = p1.InsertPolynomeTab(4,8)

    p2 := PolynomeTab{}
    _, err := AddPolynomeTab(p1,p2)

    if err == nil {
        t.Errorf("erreur dans TestAddPolynomeTab2")
    }
}

// Teste si l'addition échoue avec un paramètre non initialisé. 
func TestAddPolynomeTab3(t *testing.T){
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(3,0)
    _ = p1.InsertPolynomeTab(1,2)
    _ = p1.InsertPolynomeTab(1,3)
    _ = p1.InsertPolynomeTab(4,7)
    _ = p1.InsertPolynomeTab(4,8)

    p2 := PolynomeTab{}
    _, err := AddPolynomeTab(p2,p1)

    if err == nil {
        t.Errorf("erreur dans TestAddPolynomeTab3")
    }
}

// Teste l'addition avec un paramètre vide. 
func TestAddPolynomeTab4(t *testing.T){
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(3,0)
    _ = p1.InsertPolynomeTab(1,2)
    _ = p1.InsertPolynomeTab(1,3)
    _ = p1.InsertPolynomeTab(4,7)
    _ = p1.InsertPolynomeTab(4,8)

    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    _, err := AddPolynomeTab(p1,p2)

    if err != nil {
        t.Errorf("erreur dans TestAddPolynomeTab4")
    }
}

// Teste l'addition avec un paramètre vide. 
func TestAddPolynomeTab5(t *testing.T){
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(3,0)

    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    _, err := AddPolynomeTab(p2,p1)

    if err != nil {
        t.Errorf("erreur dans TestAddPolynomeTab5")
    }
}

// Teste la multiplication de deux polynomes. 
func TestMultPolynomeTab1(t * testing.T) {
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(2,0)
    _ = p1.InsertPolynomeTab(5,1)
    _ = p1.InsertPolynomeTab(2,3)

    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    _ = p2.InsertPolynomeTab(3,0)
    _ = p2.InsertPolynomeTab(2,2)
    _ = p2.InsertPolynomeTab(2,4)

    p3, _ := MultPolynomeTab(p1,p2)

    str := p3.PrintPolynomeTab()

    if str != "6 + 15X^1 + 4X^2 + 16X^3 + 4X^4 + 14X^5 + 4X^7"{
        t.Errorf("erreur dans TestMultPolynomeTab1")
    }
}

// Teste la multiplication avec un paramètre non initialisé. 
func TestMultPolynomeTab2(t *testing.T){
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(3,0)
    _ = p1.InsertPolynomeTab(1,2)
    _ = p1.InsertPolynomeTab(1,3)
    _ = p1.InsertPolynomeTab(4,7)
    _ = p1.InsertPolynomeTab(4,8)

    p2 := PolynomeTab{}
    _, err := MultPolynomeTab(p1,p2)

    if err == nil {
        t.Errorf("erreur dans TestMultPolynomeTab2")
    }
}

// Teste la multiplication avec un paramètre non initialisé.  
func TestMultPolynomeTab3(t *testing.T){
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(3,0)
    _ = p1.InsertPolynomeTab(1,2)
    _ = p1.InsertPolynomeTab(1,3)
    _ = p1.InsertPolynomeTab(4,7)
    _ = p1.InsertPolynomeTab(4,8)

    p2 := PolynomeTab{}
    _, err := MultPolynomeTab(p2,p1)

    if err == nil {
        t.Errorf("erreur dans TestMultPolynomeTab3")
    }
}

// Teste la multiplication avec un paramètre vide. 
func TestMultPolynomeTab4(t *testing.T){
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(3,0)
    _ = p1.InsertPolynomeTab(1,2)
    _ = p1.InsertPolynomeTab(1,3)
    _ = p1.InsertPolynomeTab(4,7)
    _ = p1.InsertPolynomeTab(4,8)

    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    _, err := MultPolynomeTab(p1,p2)

    if err != nil {
        t.Errorf("erreur dans TestMultPolynomeTab4")
    }
}

// Teste la multiplication avec un paramètre vide. 
func TestMultPolynomeTab5(t *testing.T){
    p1 := PolynomeTab{}
    InitPolynomeTab(&p1)
    _ = p1.InsertPolynomeTab(3,0)

    p2 := PolynomeTab{}
    InitPolynomeTab(&p2)
    _, err := MultPolynomeTab(p2,p1)

    if err != nil {
        t.Errorf("erreur dans TestMultPolynomeTab5")
    }
}