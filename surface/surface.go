package main

import (
    "./surf"
    "flag"
    "fmt"
    "os"
    "reflect"
    "strings"
)

var width = flag.Int("width", 1024,"canvas width")
var height = flag.Int("height", 800, "canvas height")
var xyRange = flag.Float64("range",30.0, "XY plane values range")
var strokeColor = flag.String("color", "white", "polygons stroke color")
var strokeWidth = flag.Int("stroke",1,"polygons stroke width")
var functionName = flag.String("func","wave", "name of 3D surface to plot")
var filename = flag.String("out", "plot.html", "output file with SVG plot")

func init() {
    flag.Parse()
}

func main() {
    flag.Parse()

    functions := map[string]surf.SurfaceFunction {
        "wave": surf.Wave,
        "smoothed": surf.SmoothedWave,
    }

    function := functions[*functionName]
    if function == nil {
        n := len(functions)
        reflectedKeys := reflect.ValueOf(functions).MapKeys()
        keys := make([]string, n)
        for i := 0; i < n; i++ {
            keys[i] = reflectedKeys[i].String()
        }
        fmt.Printf("Unexpected surface function: %s\n", *functionName)
        fmt.Printf("Available choices are {%s}\n", strings.Join(keys, ", "))
        os.Exit(1)
    }

    config := surf.DefaultConfig
    config.Width = *width
    config.Height = *height
    config.XYRange = *xyRange
    surface := surf.Surface{Function:function, SurfaceConfig:config}
    svg := surface.Plot(*strokeColor, *strokeWidth)
    svg.Save(*filename)
}
