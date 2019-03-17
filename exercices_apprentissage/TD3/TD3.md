# TD3 : Manipulation avancée de bibliothèques Go

__But :__ Apprendre à manipuler les packages, les erreurs et les tests unitaires. 

__Difficulté :__ Difficile

__Connaissances à avoir :__ Bibliothèques du langage Go.

__Durée du TD :__ 1h15 - 1h30

## Exercice 1 : Manipulation de la bibliothèque `image`

Pour cet exercice, il est utile d'étudier la bibliothèque [`image`](https://golang.org/pkg/image/) pour faire un dessin et la bibliothèque [`os`](https://golang.org/pkg/os/). Il faudra que le fichier soit dans un chemin que l'utilisateur aura choisi. 

1. Implémenter une fonction qui dessine un damier 4x4. 

2. Implémenter une fonction qui dessine une fractale de [Mandelbrot](http://sdz.tdct.org/sdz/dessiner-la-fractale-de-mandelbrot.html). 

Pour voir la solution intégrale, cliquer [ici](exo3.1/solution_exo3.1.md)

## Exercice 2 : Manipulation de la bibliothèque `io`

Pour les questions suivantes, il faut utiliser les fonctions du package [`io`](https://golang.org/pkg/io/). 

1. Implémenter une fonction qui copie le contenu d'un fichier (x.txt) vers un autre (y.txt). Pour cela, l'utilisateur devra taper la commande `./copy ./x.txt ./y.txt`. 

2. Implémenter une fonction qui va lister les éléments contenus dans un dossier dont son chemin a été choisi par l'utilisateur en paramètre. Il faudra afficher les droits, la taille, la date de dernière modification, le fait qu'il soit un dossier ou non et le nom de l'élément en question. 

Pour voir la solution intégrale, cliquer [ici](exo3.2/solution_exo3.2.md)