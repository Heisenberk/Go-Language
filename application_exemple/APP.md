# Application exemple

Ce projet consiste à créer une application en ligne de commande pour le chiffrement et le déchiffrement cryptographique de fichiers. La méthode de chiffrement sera choisi par le développeur selon le package [`crypto`](https://golang.org/pkg/crypto/) de Golang. 

## Description de l'application 

L'application proposera les fonctionnalités suivantes : 

- chiffrer (`--encrypt` ou `-e`) un ensemble de fichiers et/ou de contenus de dossiers. 
- déchiffrer (`--decrypt` ou `-d`) un ensemble de fichiers et/ou de contenus de dossiers. 

L'utilisateur devra choisir un mot de passe précédé de `--password` ou `-p`. 

## Exemple d'utilisation

- `./goshield --encrypt -p projet dossier1/sous-dossier2/ fichier1.txt image.png` chiffre le contenu de `sous-dossier2/`, le fichier texte `fichier1.txt` et le fichier image `image.png` avec le mot de passe `projet`. 
- `./goshield --decrypt -p projet dossier1/sous-dossier2/ fichier1.txt.gs image.png.gs` déchiffre le contenu de `sous-dossier2/`, le fichier texte chiffré `fichier1.txt.gs` et le fichier image chiffré `image.png.gs` avec le mot de passe `projet`. 

## Contraintes techniques 

- La construction du projet devra être dans le package `goshield`.
- La plupart des fonctions, méthodes devront être associées à des tests unitaires.
- Le projet devra comporter une documentation sous la forme d'un fichier `README.md` et de commentaires `Godoc` dans le code source. 
- La documentation devra décrire l'usage de l'application (manuel utilisateur) ainsi que la conception (manuel technique). 
- La gestion de la concurrence est exigée pour réaliser l'application. 


Pour voir une solution intégrale, cliquer [ici](http://www.github.com/Heisenberk/GoShield/)