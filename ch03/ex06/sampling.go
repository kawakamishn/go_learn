// ./sampling

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {

	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	xUnit := float64(xmax-xmin) / width
	yUnit := float64(ymax-ymin) / height

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)*yUnit + ymin
		for px := 0; px < width; px++ {
			x := float64(px)*xUnit + xmin
			xPlus := x + xUnit/2 //半単位分だけずらす
			yPlus := y + yUnit/2 //半単位分だけずらす
			z1, z2, z3, z4 := complex(x, y), complex(x, yPlus), complex(xPlus, y), complex(xPlus, yPlus)
			mb1, mb2, mb3, mb4 := mandelbrot(z1), mandelbrot(z2), mandelbrot(z3), mandelbrot(z4)
			RAve := (mb1.R + mb2.R + mb3.R + mb4.R)
			GAve := (mb1.G + mb2.G + mb3.G + mb4.G)
			BAve := (mb1.B + mb2.B + mb3.B + mb4.B)
			AAve := (mb1.A + mb2.A + mb3.A + mb4.A)
			mbAve := color.RGBA{R: RAve, G: GAve, B: BAve, A: AAve}
			img.Set(px, py, mbAve)
		}
	}
	Output, _ := os.Create("./new.png")
	png.Encode(Output, img) //注意：エラーを無視
}

func mandelbrot(z complex128) color.RGBA {

	const (
		iterations = 200
		contrast   = 15
	)

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{R: 100, G: 255 - contrast*n, B: 255 - contrast*n, A: 100}
		}
	}

	return color.RGBA{R: 255, G: 0, B: 50, A: 50}
}
