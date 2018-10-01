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
