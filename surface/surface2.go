package main

import (
    "math"
    "./surf"
)

func main() {
    surface := surf.Surface{Function:z, SurfaceConfig:surf.DefaultConfig}
    surface.Plot("plot.html")
}

func z(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}
