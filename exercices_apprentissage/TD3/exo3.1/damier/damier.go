package main

import "image"
import "image/color"
import "image/draw"
import "image/png"
import "os"
import "fmt"

// Fonction qui dessine un damier
func dessineDamier(path string) int{
    // on va utiliser 2 couleurs pour le damier : blanc et noir
    var couleur[2] color.RGBA
    couleur[0]=color.RGBA{0,0,0,255}
    couleur[1]=color.RGBA{255,255,255,255}

    bordure:= 250
    //déclaration de l'image qui sera de taille 250x250 pixels
    damier := image.NewRGBA(image.Rect(0, 0, bordure, bordure))

    var swap_color int = 0
    // dessin des cases du damier qui auront une taille de 50 pixels
    for x := 0; x< 250; x=x+50 {
        for y :=0 ; y<250; y=y +50 {
            draw.Draw(damier, image.Rect(x, y, x+50,y+50 ), &image.Uniform{couleur[swap_color]}, image.ZP, draw.Src)
            swap_color=1 - swap_color
        }
    }

    // création du fichier avec le chemin choisi par l'utilisateur
    myfile, err1 := os.Create(path)
    if err1 != nil {
        return 1
    }

    // écriture de l'image dans le fichier
    err2 := png.Encode(myfile, damier)
    if err2 != nil {
        return 1
    }

    return 0
}

func main(){
    arg := os.Args[1:]
    if len(arg) == 0 {
        fmt.Println("Aucun argument. ")
    }else {
        err := dessineDamier(arg[0])
        if err == 1 {
            fmt.Println("Impossible de dessiner le damier. ")
        }else{
            fmt.Println("Dessin du damier effectué. ")
        }
    }
}

