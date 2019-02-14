package main

import (
    "image"
    "image/color"
    "image/draw"
    "os"
    "image/png"
    "math/cmplx"
)

const(

    Bordure = 600

    Xmin = -2

    Xmax = 2

    Ymin = -2

    Ymax = 2

    ItérationMax = 100

    BorneModule = 2
)

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


func main() {
new_png_file := "/home/user/Bureau/Go-Language/exercice/dessin/mandelbrot.png"

    bordure:= 600

    myimage := image.NewRGBA(image.Rect(0, 0, bordure, bordure))

    var couleur color.RGBA

    couleur=color.RGBA{255,255,255,255}

    draw.Draw(myimage, image.Rect(0, 0, bordure,bordure), &image.Uniform{couleur}, image.ZP, draw.Src)


    for x := 0 ; x<Bordure ; x++{

        i:=float64(x)/Bordure*(Xmax - Xmin) + Xmin

        for y := 0 ; y<Bordure ;y++{

            j:=float64(y)/Bordure*(Ymax - Ymin) + Ymin

            c := complex(i,j)
            
            myimage.Set(x,y,mandelbrot(c))


    }
}
           
    myfile, _ := os.Create(new_png_file)

    png.Encode(myfile, myimage)

}
