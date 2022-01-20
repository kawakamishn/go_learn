// GOMAXPROCS=8 ./mandelbrot_parallel_09_06

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
)

func main() {
	nParallel := 2
	main_(nParallel)
}

func main_(nParallel int) {
	var wg sync.WaitGroup
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for j := 0; j < nParallel; j++ { // ループをnParallelに並列化
		wg.Add(1)
		go pixel(j, nParallel, img, &wg)
	}
	wg.Wait()
	Output, _ := os.Create("./new.png")
	png.Encode(Output, img) //注意：エラーを無視
	secs := time.Since(start).Seconds()
	fmt.Println("GOMAXPROCS", os.Getenv("GOMAXPROCS"), ":", secs)
}

func pixel(i int, nParallel int, img *image.RGBA, wg *sync.WaitGroup) {
	defer wg.Done()
	for py := i; py < height; py += nParallel {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
}

func mandelbrot(z complex128) color.Color {

	const (
		iterations = 200
		contrast   = 15
	)

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{100, 255 - contrast*n, 255 - contrast*n, 100} // ここを教科書から変えた
		}
	}

	return color.RGBA{255, 0, 50, 50} // ここを教科書から変えた
}
