# TD1 Exercice 3 : Solution 

- Exercice de Fibonacci :

```go
package main

import "fmt"
import "strconv"
import "os"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Aucun argument. ")
	}else {
		val, err := strconv.Atoi(arg[0])
		if err != nil {
			fmt.Println("Argument invalide. ")
		}else if val < 0 {
			fmt.Println("Impossible de calculer la factorielle d'un nombre négatif. ")
		}else {
			fmt.Println(fibonacci(val))
		}
	}
}
```

- Exercice de Factorielle : 

```go
package main

import "fmt"
import "strconv"
import "os"

func factorielle (n int) int {
    if n == 0 {
        return 1
    }
    return n * factorielle(n-1)
}

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("Aucun argument. ")
	}else {
		val, err := strconv.Atoi(arg[0])
		if err != nil {
			fmt.Println("Argument invalide. ")
		}else if val < 0 {
			fmt.Println("Impossible de calculer la factorielle d'un nombre négatif. ")
		}else {
			fmt.Println(factorielle(val))
		}
	}
}
```

Pour avoir le code complet : [exo1.3](exo1.3.zip)