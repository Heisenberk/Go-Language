package main

import (
	"image"
	"image/color"
	"image/draw"
	"os"
	"image/png"
)

func main() {
new_png_file := "/home/user/Bureau/Go-Language/exercice/dessin/damier.png"
    bordure:= 250

    myimage := image.NewRGBA(image.Rect(0, 0, bordure, bordure))

    var couleur [2]color.RGBA
    couleur[0]=color.RGBA{0,0,0,255}
    couleur[1]=color.RGBA{255,255,255,255}

    var swap_color int =0

    for x := 0; x< 250; x=x+50 {
    	for y :=0 ; y<250; y=y +50 {
    		draw.Draw(myimage, image.Rect(x, y, x+50,y+50 ), &image.Uniform{couleur[swap_color]}, image.ZP, draw.Src)
    		swap_color=1 - swap_color
    	}
    }



    myfile, _ := os.Create(new_png_file)    
    png.Encode(myfile, myimage)
}
