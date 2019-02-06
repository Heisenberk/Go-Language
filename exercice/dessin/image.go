package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"

)

func main() {



    new_png_file := "/home/user/Bureau/Go-Language/exercice/dessin/deux_rectangles.png"

    myimage := image.NewRGBA(image.Rect(0, 0, 220, 220)) 
    mygreen := color.RGBA{0, 100, 0, 255}  

    
    draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)

    red_rect := image.Rect(60, 80, 120, 160) 
    myred := color.RGBA{200, 0, 0, 255}



    // create a red rectangle atop the green surface
    draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)

    myfile, _ := os.Create(new_png_file)     // ... now lets save imag
    png.Encode(myfile, myimage)
}


