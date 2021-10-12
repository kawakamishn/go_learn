package main

import (
	"fmt"
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

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	var ax, ay, bx, by, cx, cy, dx, dy float64 // グリッド値をここで事前に定義しておく
	var isInvalid bool                         // 出力がInfまたはNanならTrue

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	pts := plotter.XYs{} // 点の集まりインスタンス
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
			pts = append(pts, plotter.XY{X: ax, Y: ay}) // ここでaxとayのみ追加
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}

	}
	fmt.Println("</svg>")
	p := plot.New()
	p.Title.Text = "picture"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	s, _ := plotter.NewScatter(pts)

	p.Add(s, plotter.NewGrid())
	p.Save(4*vg.Inch, 4*vg.Inch, "plot.png")
}

func corner(i, j int) (float64, float64, bool) { //出力にisInvalidを加えた
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale + 2*z*zscale
	return sx, sy, math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy) //出力がInfまたはNanならTrue
}

func f(x, y float64) float64 { // sin(r)/r
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func fEggBox(x, y float64) float64 { // 卵の箱
	segment := 5
	x = math.Abs(float64(int(x*10.0)%(segment*10))) / 10.0 // 0~4.7
	y = math.Abs(float64(int(y*10.0)%(segment*10))) / 10.0
	//fmt.Println(x, y)
	half := float64(segment) / 2.0
	r := (x-half)*(x-half) + (y-half)*(y-half)
	return r / 20.0
}
