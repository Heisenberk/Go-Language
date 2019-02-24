# TD1 Exercice 1 : Solution 

1. Aller sur le site de Golang pour récupérer le fichier `.tar.gz` (pour Linux). Par exemple, la dernière version disponible est `go1.11.5.linux-amd64.tar.gz`. Puis, taper la commande suivante pour installer l'environnement de `Go`: 
```shell
tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz
```
Ensuite, aller dans le fichier `/home/user/.profile` et ajouter à la fin du fichier la ligne `PATH="/usr/local/go/bin:$PATH"`. 
Puis, pour installer les commandes `go`, taper la commande `sudo apt-get install golang-go`.
Enfin, taper la commande `go env` et vérifier le chemin pour `$GOPATH`. Ce chemin va correspondre au répertoire où il faudra programmer en langage Go. 

2. Dans le dossier `$GOPATH/src/`, créer le dossier `hello` dans lequel il faudra créer le fichier `hello.go`. 

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world! \n")
}
```
Pour compiler, il suffit de taper `go build hello`, puis, pour exécuter, faire `./hello`. 

Pour avoir le code complet : [exo1.1](exo1.1.zip)
