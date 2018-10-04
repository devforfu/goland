package surf

import (
    "fmt"
    "math"
)

type SurfaceFunction func (float64, float64) float64

type ColorBar struct {
    Min, Max float64
}

func (cb *ColorBar) Normalize(value float64) float64 {
    return (value - cb.Min)/(cb.Max - cb.Min)
}

type SurfaceConfig struct {
    Width, Height int
    Cells int
    XYRange float64
    Angle float64
}

type Surface struct {
    Function SurfaceFunction
    SurfaceConfig
    colorBar *ColorBar
}

var DefaultConfig = SurfaceConfig{
   Width:600, Height:320,
   Cells:100, XYRange:30.0,
   Angle:math.Pi/6,
}

func (sf *Surface) Sin() float64 { return math.Sin(sf.Angle) }

func (sf *Surface) Cos() float64 { return math.Cos(sf.Angle) }

func (sf *Surface) XYScale() float64 { return float64(sf.Width)/2/sf.XYRange }

func (sf *Surface) ZScale() float64 { return float64(sf.Height)*0.4 }

func (sf *Surface) Corner(i, j int) (float64, float64, float64) {
    x := sf.XYRange * (float64(i)/float64(sf.Cells) - 0.5)
    y := sf.XYRange * (float64(j)/float64(sf.Cells) - 0.5)
    z := sf.Function(x, y)
    return x, y, z
}

func (sf *Surface) Projection(x, y, z float64) (float64, float64) {
    sx := float64(sf.Width)/2 + (x-y)*sf.Cos()*sf.XYScale()
    sy := float64(sf.Height)/2 + (x+y)*sf.Sin()*sf.XYScale()- z*sf.ZScale()
    return sx, sy
}

func (sf *Surface) CreateColorBar() {
    lo, hi := math.Inf(1), math.Inf(-1)
    for i := 0; i < sf.Cells; i++ {
        for j := 0; j < sf.Cells; j++ {
            _, _, z := sf.Corner(i, j)
            if math.IsNaN(z) { continue }
            lo = math.Min(lo, z)
            hi = math.Max(hi, z)
        }
    }
    sf.colorBar = &ColorBar{lo,hi}
}

func (sf *Surface) DefaultPlot(filename string) {
    sf.Plot(filename, "while", 1)
}

func (sf *Surface) Plot(filename string, strokeColor string, strokeWidth int) {
    if sf.colorBar == nil {
        sf.CreateColorBar()
    }
    n := sf.Cells
    polygons := make([]*Polygon, n*n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            xs := [4]int{i + 1, i,     i, i + 1}
            ys := [4]int{j    , j, j + 1, j + 1}
            polygons[i*n + j] = sf.createPolygon(xs[:], ys[:])
       }
    }
    svg := SVG{
        Stroke:strokeColor, Fill:"white",
        StrokeWidth:strokeWidth,
        Width:sf.Width, Height:sf.Height,
        CloseTag:true,
    }
    svg.CreatePreamble()
    for _, polygon := range polygons {
        if polygon != nil {
            svg.WriteLine(polygon.String())
        }
    }
    svg.Save(filename)
}

func (sf *Surface) createPolygon(xs, ys []int) *Polygon {
    var average float64
    var points [8]float64
    for i, j := 0, 0; i < 4; i, j = i+1, j+2 {
        x, y, z := sf.Corner(xs[i], ys[i])
        sx, sy := sf.Projection(x, y, z)
        if math.IsNaN(z) { return nil }
        points[j], points[j+1] = sx, sy
        average += z
    }
    average /= 4
    alpha := sf.colorBar.Normalize(average)
    red, blue := int(255 * alpha), int(255 * (1-alpha))
    color := fmt.Sprintf("#%02x00%02x", red, blue)
    return &Polygon{points[:], color}
}