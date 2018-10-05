package surf

import (
    "fmt"
    "log"
    "os"
    "strings"
)

// Simple string object wrapper building SVG file with polygons
type SVG struct {
    Stroke, Fill string
    StrokeWidth int
    Width, Height int
    CloseTag bool
    buffer string
}

func (svg *SVG) CreatePreamble() {
    preamble := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' " +
        "style='stroke: %s; fill: %s; stroke-width: %d' " +
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

func (svg *SVG) String() string {
    return svg.buffer
}

type Polygon struct {
    points []float64
    color string
}

func (p *Polygon) String() string {
    var buffer = make([]string, len(p.points))
    for i, point := range p.points {
        buffer[i] = fmt.Sprintf("%g", point)
    }
    joined := strings.Join(buffer, ",")
    return fmt.Sprintf("<polygon points='%s', style='fill:%s' />", joined, p.color)
}
