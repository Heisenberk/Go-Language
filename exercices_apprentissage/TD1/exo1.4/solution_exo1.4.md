# TD1 Exercice 4 : Solution 

Tester de manière naïve la primalité d'un nombre consiste à tester s'il est divisible par un nombre plus petit que lui et différent de 2. De cette façon, créer la fonction `primalite`. De plus, il faut interpréter le résultat de cette dernière par le biais de `interpreter`. 

```go
package main

import "fmt"
import "os"
import "strconv"

func primalite (val int) int {
	if val <= 1 {
		return 0
	}else if val == 2 {
		return 1
	}else {
		for i := 2 ; i < val ; i++ {
			if val % i == 0 {
				return 0
			}
		}
	}
	return 1
}

func interpreter (val int) string {
	var test int = primalite(val)

	if test == 0 {
		return "n n'est pas premier. "
	}else{
		return "n est premier. "
	}
}

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Aucun argument. ")
	}else {
		val, err := strconv.Atoi(arg[0])
		if err != nil {
			fmt.Println("Argument invalide");
		}else {
			fmt.Println(interpreter(val))
		}
	}
}
```

Pour avoir le code complet : [exo1.4](exo1.4.zip)