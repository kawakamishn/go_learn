// クエリストリングでwidthやheightを変更できる。 http://localhost:8000/?width=400&height=900
// POSTしたければcurl -XPOST -d 'width=600&height=320' localhost:8000

package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

var width float64  //外から受け取る
var height float64 //外から受け取る
var xyscale float64
var zscale float64

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) { // ハンドラ
		r.ParseForm()
		setParameter(r)
		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func setParameter(r *http.Request) {
	width, _ = strconv.ParseFloat(r.Form.Get("width"), 64)   //外から受け取る
	height, _ = strconv.ParseFloat(r.Form.Get("height"), 64) //外から受け取る
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
}

func surface(out io.Writer) {
	var s string
	s = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	out.Write([]byte(s))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			s = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			out.Write([]byte(s))
		}
	}
	s = fmt.Sprintln("</svg>")
	fmt.Sprintln("width", width)
	out.Write([]byte(s))
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r != 0 {
		return math.Sin(r) / r
	} else {
		return 0
	}
}
