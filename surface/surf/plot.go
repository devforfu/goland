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
    xyscale, zscale float64
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

func (sf *Surface) Corner(i, j int) (float64, float64, float64){
    x := sf.XYRange * (float64(i)/float64(sf.Cells) - 0.5)
    y := sf.XYRange * (float64(j)/float64(sf.Cells) - 0.5)
    z := sf.Function(x, y)
    return x, y, z
}

func (sf *Surface) Projection(x, y, z float64) (float64, float64) {
    sx := float64(sf.Width)/2 + (x-y)*sf.Cos()*sf.xyscale
    sy := float64(sf.Height)/2 + (x+y)*sf.Sin()*sf.xyscale - z*sf.zscale
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

func (sf *Surface) Plot() []*Polygon {
    if sf.colorBar == nil {
        sf.CreateColorBar()
    }
    n := sf.Cells
    polygons := make([]*Polygon, n*n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            xs := [4]int{i + 1, i,     i, i + 1}
            ys := [4]int{j    , j, j + 1, j + 1}
            polygons[i*n + j] = sf.createPolygon(xs, ys)
       }
   }
    return polygons
}

func (sf *Surface) createPolygon(xs, ys [4]int) *Polygon {
    var average float64
    var points [8]float64
    for i := 0; i < 8; i += 2 {
        x, y, z := sf.Corner(xs[i], ys[i])
        sx, sy := sf.Projection(x, y, z)
        if math.IsNaN(z) { return nil }
        points[i], points[i+1] = sx, sy
        average += z
    }
    average /= 4
    alpha := sf.colorBar.Normalize(average)
    red, blue := int(255 * alpha), int(255 * (1-alpha))
    color := fmt.Sprintf("#%x0%x", red, blue)
    return &Polygon{points[:], color}
}