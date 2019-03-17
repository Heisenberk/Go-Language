package main

import "image"
import "image/color"
import "image/draw"
import "image/png"
import "os"
import "math/cmplx"
import "fmt"

const Bordure = 600
const Xmin = -2
const Xmax = 2
const Ymin = -2
const Ymax = 2
const ItérationMax = 100
const BorneModule = 2

// Dessin des pixels de la fractale de Mandelbrot 
func mandelbrot(c complex128) color.Color {
   var z complex128
    for n := 0; n< ItérationMax; n++ {
        z=z*z+c
        if cmplx.Abs(z)> BorneModule {
            return color.RGBA{255,255,255,255}
        }
    }
    return color.Black
}

// Réalise le dessin
func dessineMandelbrot(path string) int{
    // Initialisation du dessin
    bordure:= 600
    dessin_png := image.NewRGBA(image.Rect(0, 0, bordure, bordure))
    var couleur color.RGBA
    couleur=color.RGBA{255,255,255,255}
    draw.Draw(dessin_png, image.Rect(0, 0, bordure,bordure), &image.Uniform{couleur}, image.ZP, draw.Src)

    // Dessin de chaque pixel
    for x := 0 ; x<Bordure ; x++{
        i:=float64(x)/Bordure*(Xmax - Xmin) + Xmin
        for y := 0 ; y<Bordure ;y++{
            j:=float64(y)/Bordure*(Ymax - Ymin) + Ymin
            c := complex(i,j)
            dessin_png.Set(x,y,mandelbrot(c))
        }
    }
    
    // Ecriture dans le fichier
    fichier, err := os.Create(path)
    if err != nil {
        return 1
    }
    png.Encode(fichier, dessin_png)
    return 0
}

// Fonction principale
func main() {
    arg := os.Args[1:]
    if len(arg) == 0 {
        fmt.Println("Aucun argument. ")
    }else {
        err := dessineMandelbrot(arg[0])
        if err == 1 {
            fmt.Println("Impossible de dessiner la fractale. ")
        }else{
            fmt.Println("Dessin de la fractale effectué. ")
        }
    }

}