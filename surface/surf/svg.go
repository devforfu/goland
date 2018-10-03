package surf

import (
    "fmt"
    "log"
    "os"
)

// Simple string object wrapper building SVG file with polygons
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
    if svg.CloseTag { svg.WriteLine("</svg>") }
    file, err := os.Create(filename)
    if err != nil { log.Fatal("Cannot save svg file") }
    file.WriteString(svg.buffer)
    file.Sync()
    file.Close()
}

type Polygon struct {
    //ax, ay, bx, by, cx, cy, dx, dy float64
    points []float64
    color string
}

func (p *Polygon) String() string {
    var buffer string
    for _, point := range p.points {
        buffer += fmt.Sprintf("%g", point)
    }
    return fmt.Sprintf("<polygon points='%s', style='fill:%s' />", buffer, p.color)

    //return fmt.Sprintf(
    //    "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g', style='fill:%s' />",
    //    p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, p.color)
}
