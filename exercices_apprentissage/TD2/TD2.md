# TD2 : Manipulation de structures complexes

__But :__ Apprendre à manipuler les tableaux, les slices et les maps. 

__Difficulté :__ Intermédiaire

__Connaissances à avoir :__ Structures, Tableaux, Slices, Maps, Méthodes, Pointeur.

__Durée du TD :__ 1h30 - 1h45

## Exercice 1 : Tableau, Slice et Tri

1. Créer une fonction qui fait le tri par propagation (tri à bulles) d'un tableau en paramètre. 

2. Créer une fonction qui fait de manière récursive le tri fusion d'un tableau en paramètre. 

Pour voir la solution intégrale, cliquer [ici](exo2.1/solution_exo2.1.md)

## Exercice 2 : Implémentation d'un annuaire

On souhaite implémenter un annuaire électronique, qui donne pour chaque nom de contact ses coordonnées (numéro de téléphone et adresse). 

1. Définir la structure Coordonnées. 

2. Définir la structure Annuaire qui gère l'annuaire électronique. 

3. Définir les fonctions suivantes : l'ajout d'un nouveau contact, la modification des coordonnées d'un contact, la suppression d'un contact et l'affichage de la liste des contacts de l'annuaire. 

Pour voir la solution intégrale, cliquer [ici](exo2.2/solution_exo2.2.md)

## Exercice 3 : Implémentation d'une liste chaînée

On souhaite implémenter une liste d'entiers en langage Go sans utiliser la bibliothèque standard `list`.

1. Définir la structure Liste

2. Définir les méthodes suivantes : l'ajout d'un élément au début de la liste, l'ajout d'un élément à la fin de la liste, la suppression du premier élément de la liste, la suppression du dernier élément de la liste et l'affichage de la liste. 

Pour voir la solution intégrale, cliquer [ici](exo2.3/solution_exo2.3.md)

## Exercice 4 

On souhaite comparer le temps de calcul de la fonction non récursive de la factorielle avec et sans goroutines. 

1. Ecrire une fonction qui calcule de façon non récursive la factorielle d'un grand nombre sans goroutines pour l'instant. 

2. Ecrire une fonction qui calcule de façon non récursive la factorielle d'un grand nombre avec l'implémentation des goroutines. Vous pouvez vous aider avec le lien [suivant](https://fr.wikibooks.org/wiki/Programmation_en_Go/Goroutines). 

Pour voir la solution intégrale, cliquer [ici](exo2.4/solution_exo2.4.md)
