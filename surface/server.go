package main

import (
    "./surf"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strings"
)

type EchoResponse struct {
    Method string
    URL string
    Headers map[string]string
}

func main() {
    http.HandleFunc("/", echo)
    http.HandleFunc("/surface", plotSurface)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
    headers := map[string]string{}
    for k, v := range r.Header {
        headers[k] = strings.Join(v, ",")
    }
    responseInfo := EchoResponse{Method:r.Method, URL:r.URL.String(), Headers:headers}
    content, err := json.MarshalIndent(responseInfo, ""," ")
    if err == nil {
        fmt.Fprintf(w, string(content))
    } else {
        fmt.Fprintf(w, "{'error': '%s'}", err)
    }
}

func plotSurface(w http.ResponseWriter, r *http.Request) {
    surface := surf.Surface{Function:surf.Wave, SurfaceConfig:surf.DefaultConfig}
    svg := surface.Plot("white",1.0)
    response := fmt.Sprintf("<html><head></head><body>%s</body></html>", svg.String())
    fmt.Fprintf(w, response)
}