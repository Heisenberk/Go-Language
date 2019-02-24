
# TD1 : Bases du langage Go

But : Apprendre l'environnement et manipuler les différentes bases du langage de programmation Go. 

Difficulté : Facile

Connaissances à avoir : Variables, Conditions, Fonctions,  Bibliothèques standards, Ligne de commande et Arguments, Structures.

Durée du TD : 1h15 - 1h30

## Exercice 1 : Installation

Avant de programmer en langage Go, il faut installer l'environnement. Pour cela, aller sur la documentation officielle de [Golang](https://golang.org/doc/install).

1. Installer l'environnement pour programmer en langage Go. 

2. Vérifier que la compilation et l'exécution d'un programme basique "Hello, world!" fonctionne correctement. 

Pour voir la solution intégrale, cliquer [ici](exo1.1/solution_exo1.1.md)


## Exercice 2 : Conversions

Ecrire un programme qui convertit un temps donné en secondes, minutes et secondes (avec l'accord des pluriels). Le calcul doit se faire dans une fonction et doit retourner, par exemple, `3620 correspond à 1 heure, 0 minute et 20 secondes`. 

Note : Utiliser la bibliothèque standard `strconv` pour convertir un entier en string. 

Pour voir la solution intégrale, cliquer [ici](exo1.2/solution_exo1.2.md)


## Exercice 3 : Récursion

Il est possible de faire des appels récursifs en langage Go. Fibonacci et Factorielle sont deux fonctions illustrant la récursivité. 

1. Ecrire un programme récursif qui permet de calculer la `n` iéme valeur de la suite de Fibonacci. 

2. Ecrire un programme récursif qui permet de calculer la factorielle de `n`.

L'utilisateur du programme doit pouvoir entrer en paramètre d'arguments le nombre à calculer. 

Pour voir la solution intégrale, cliquer [ici](exo1.3/solution_exo1.3.md)


## Exercice 4 : Primalité

Ecrire un fonction `primalite` qui teste de la manière la plus naïve possible la primalité d'un nombre, et une fonction `interpreter` qui retourne la chaîne de caractères affichée sur le terminal. L'utilisateur du programme doit pouvoir entrer en paramètre le nombre à tester. 

Pour voir la solution intégrale, cliquer [ici](exo1.4/solution_exo1.4.md)


## Exercice 5 : Nombres Amicaux

Soit `n` et `m`, deux entiers positifs. `n` et `m` sont dit amis si la somme de tous les diviseurs de `n` (sauf `n` lui-même) est égale à `m` et si la somme de tous les diviseur de `m` (sauf `m` lui-même) est égale à `n`. 
Ecrire un programme qui teste si deux entiers sont des nombres amis ou non. L'utilisateur du programme doit pouvoir entrer en paramètre le nombre à tester. 

Pour voir la solution intégrale, cliquer [ici](exo1.5/solution_exo1.5.md)

## Exercice 6 : Fraction

1. Définir une structure de données `Fraction` permettant de représenter en forme fractionnaire un nombre flottant. De plus, la fraction doit toujours être irréductible. 

2. Ecrire les fonctions qui permettent d'initialiser, d'afficher une fraction, d'additionner, de soustraire, de multiplier et de diviser 2 fractions. 

Pour voir la solution intégrale, cliquer [ici](exo1.6/solution_exo1.6.md)
