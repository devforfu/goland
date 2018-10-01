package surf

import (
    "fmt"
    "math"
)

type Polygon struct {
    ax, ay, bx, by, cx, cy, dx, dy float64
    color string
}

func (p *Polygon) String() string {
    return fmt.Sprintf(
        "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g', style='fill:%s' />",
        p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, p.color)
}

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

//func (sf *Surface) Plot() []Polygon {
//    for i := 0; i < sf.Cells; i++ {
//        for j := 0; j < sf.Cells; j++ {
//
//        }
//    }
//}