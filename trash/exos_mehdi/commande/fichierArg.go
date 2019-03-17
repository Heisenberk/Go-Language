package main
import "fmt"
import "os"
import "io/ioutil"
import "strconv"

func check(e error) {

    if e != nil {

        panic(e)

    }
}

func main() {

	fmt.Println()

    dat, err := ioutil.ReadFile("/home/user/Bureau/"+os.Args[1])	

    check(err)

    secondes , err1 := strconv.Atoi(string(dat))

    check(err1)

	fmt.Println("Entrée",secondes,"secondes")
		
	var heure int = secondes/3600 
	var minute int = (secondes % 3600)/60
	var seconde int = secondes%60

	fmt.Println("heure :",heure)
	fmt.Println("minute :",minute)
	fmt.Println("seconde :", seconde);

    

}