package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	var ax, ay, bx, by, cx, cy, dx, dy float64 // グリッド値をここで事前に定義しておく
	var isInvalid bool                         // 出力がInfまたはNanならTrue

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if ax, ay, isInvalid = corner(i+1, j); isInvalid {
				continue
			}
			if bx, by, isInvalid = corner(i, j); isInvalid {
				continue
			}
			if cx, cy, isInvalid = corner(i, j+1); isInvalid {
				continue
			}
			if dx, dy, isInvalid = corner(i+1, j+1); isInvalid {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) { //出力にisInvalidを加えた
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy) //出力がInfまたはNanならTrue
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r != 0 {
		return math.Sin(r) / r
	} else {
		return 0
	}
}
