package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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
	var ax, ay, bx, by, cx, cy, dx, dy float64
	var isInvalid bool
	var isPeak bool

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	ptsUp := plotter.XYs{}   // z>=0の点の集まりインスタンス
	ptsDown := plotter.XYs{} // z<0の点の集まりインスタンス
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if ax, ay, isInvalid, isPeak = corner(i+1, j); isInvalid {
				continue
			}
			if bx, by, isInvalid, _ = corner(i, j); isInvalid {
				continue
			}
			if cx, cy, isInvalid, _ = corner(i, j+1); isInvalid {
				continue
			}
			if dx, dy, isInvalid, _ = corner(i+1, j+1); isInvalid {
				continue
			}
			if isPeak { // zが0より大きいか小さいかでインスタンスを分ける。
				ptsUp = append(ptsUp, plotter.XY{X: ax, Y: ay})
			} else {
				ptsDown = append(ptsDown, plotter.XY{X: ax, Y: ay})
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}

	}
	fmt.Println("</svg>")
	p := plot.New()
	p.Title.Text = "picture"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	sUp, _ := plotter.NewScatter(ptsUp)
	sUp.GlyphStyle.Color = color.RGBA{R: 128, G: 255, B: 255, A: 128} // 色を変える

	sDown, _ := plotter.NewScatter(ptsDown)
	sDown.GlyphStyle.Color = color.RGBA{R: 255, G: 255, B: 128, A: 128}

	p.Add(sDown, sUp, plotter.NewGrid())
	p.Save(4*vg.Inch, 4*vg.Inch, "plot.png")
}

func corner(i, j int) (float64, float64, bool, bool) { //z>=0でTrueになるisIsPeakを出力に加えた
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale + 2*z*zscale

	var isValid bool = math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy)

	var isPeak bool //上半分が赤で下半分が青
	if z >= 0 {
		isPeak = true
	} else {
		isPeak = false
	}

	return sx, sy, isValid, isPeak
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
