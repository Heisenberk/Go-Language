package main

import (
    "image"
    "image/color"
    "image/draw"
    "os"
    "image/png"
    "math/cmplx"
)

func mandelbrot(c complex128) color.Color {

   var z complex128

               for n := 0; n< 50; n++ {
            z=z*z+c
            if cmplx.Abs(z)>2 {
            
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


      for x := 0 ; x<700 ; x++{
        i:=float64(x)/700*4-2
        for y := 0 ; y<700 ;y++{
            j:=float64(y)/700*4-2
            c := complex(i,j)
            myimage.Set(x,y,mandelbrot(c))


    }
}
           
        
    
  



    myfile, _ := os.Create(new_png_file)    
    png.Encode(myfile, myimage)
}
