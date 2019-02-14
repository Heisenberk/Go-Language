package main

import "fmt"
import "strconv"

func conversion (val int) string {
	var heure int = val / 3600 
	var minute int = (val % 3600) / 60
	var seconde int = val % 60

	var phrase string = " correspond à "

	var heure_str string;
	var minute_str string;
	var seconde_str string;

	if heure <= 1 {
		heure_str = " heure, "
	}else{
		heure_str = " heures, "
	} 

	if minute <= 1 {
		minute_str = " minute, "
	}else{
		minute_str = " minutes, "
	} 

	if seconde <= 1 {
		seconde_str = " seconde. "
	}else{
		seconde_str = " secondes. "
	} 
	
	return strconv.Itoa(val) + phrase + strconv.Itoa(heure) + heure_str + strconv.Itoa(minute) + minute_str + strconv.Itoa(seconde) + seconde_str;
}


func main() {
	fmt.Println(conversion(3620))
}