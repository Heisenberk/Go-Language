
# TD1 : Bases du langage Go

But : Apprendre l'environnement et manipuler les différentes bases du langage de programmation Go. 
Difficulté : Facile

## Exercice 1

Avant de programmer en langage Go, il faut installer l'environnement. Pour cela, allez sur la documentation officielle de [Golang](https://golang.org/doc/install).

1. Installez l'environnement pour programmer en langage Go. 

2. Vérifiez que la compilation et l'exécution d'un programme basique "Hello, world!" fonctionne correctement. 

Solutions: 

1. Allez sur le site de Golang pour récupérer le fichier `.tar.gz` (pour Linux). Par exemple, la dernière version disponible est `go1.11.5.linux-amd64.tar.gz`. Puis, tapez la commande suivante pour installer l'environnement de `Go`: 
```shell
tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz
```
Ensuite, allez dans le fichier `/home/user/.profile` et ajoutez à la fin du fichier la ligne `PATH="/usr/local/go/bin:$PATH"`. 
Puis, pour installer les commandes `go`, tapez la commande `sudo apt-get install golang-go`.
Enfin, tapez la commande `go env` et vérifier le chemin pour `$GOPATH`. Ce chemin va correspondre au répertoire où il faudra programmer en langage Go. 

2. Dans le dossier `$GOPATH/src/`, créez le dossier `hello` dans lequel il faudra créer le fichier `hello.go`. 

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world! \n")
}
```
Pour compiler, il suffit de taper `go build hello`, puis, pour exécuter, faire `./hello`. 

