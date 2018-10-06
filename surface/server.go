package main

import (
    "./surf"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "strconv"
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
    params := parseQuery(r.URL)
    config := surf.DefaultConfig
    config.Width = parseInt(params, "width", config.Width)
    config.Height = parseInt(params, "height", config.Height)
    surface := surf.Surface{Function:surf.Wave, SurfaceConfig:config}
    svg := surface.Plot("white",1.0)
    response := fmt.Sprintf("<html><head></head><body>%s</body></html>", svg.String())
    fmt.Fprintf(w, response)
}

func parseQuery(url *url.URL) map[string]string {
    params := map[string]string{}
    for k, v := range url.Query() {
        params[k] = strings.Join(v, "")
    }
    return params
}

func parseInt(params map[string]string, key string, defaultValue int) int {
    value, ok := params[key]
    if !ok {
        return defaultValue
    }
    parsed, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
        return defaultValue
    }
    return int(parsed)
}