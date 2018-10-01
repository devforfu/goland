package main

import (
	"os"
	"fmt"
	"log"
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

type SurfaceFunction func (float64, float64) float64

type SVG struct {
	Stroke, Fill string
	StrokeWidth float64
	Width, Height int
	CloseTag bool

	buffer string
}

func (svg *SVG) CreatePreamble() {
	preamble := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: %s; fill: %s; stroke-width: %f' " +
		"width='%d' height='%d'>",
		svg.Stroke, svg.Fill, svg.StrokeWidth, svg.Width, svg.Height)
	svg.WriteLine(preamble)
}

func (svg *SVG) WriteLine(s string) {
	svg.buffer += fmt.Sprintf("%s\n", s)
}

func (svg *SVG) Save(filename string) {
	if svg.CloseTag {
		svg.WriteLine("</svg>")
	}
	file, err := os.Create(filename)
	if err != nil { log.Fatal("Cannot save svg file") }
	file.WriteString(svg.buffer)
	file.Sync()
	file.Close()
}

func corner(i, j int, f SurfaceFunction) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func z(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func main() {
	out := SVG{
		Width:width, Height:height,
		Stroke:"grey", StrokeWidth:0.7,
		Fill:"white",
		CloseTag:true}
	out.CreatePreamble()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1, j,     z)
			bx, by := corner(i,     j,     z)
			cx, cy := corner(i,     j + 1, z)
			dx, dy := corner(i + 1, j + 1, z)
			poly := fmt.Sprintf(
				"<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>",
				ax, ay, bx, by, cx, cy, dx, dy)
			out.WriteLine(poly)
		}
	}
	out.Save("plot.html")
}