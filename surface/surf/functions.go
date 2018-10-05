package surf

import "math"

func Wave(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

func SmoothedWave(x, y float64) float64 {
    w := Wave(x, y)
    return w*w
}
