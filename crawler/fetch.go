package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr,"Cannot retrieve URL: %s", url)
            continue
        }
        content, err := readContent(resp)
        if err != nil {
            fmt.Fprintf(os.Stderr,"Cannot parse HTML response content")
            continue
        }
        fmt.Printf("%s", content)
    }
}

func readContent(resp *http.Response) (string, error) {
    defer resp.Body.Close()
    html, err := ioutil.ReadAll(resp.Body)
    if err != nil { return "", err }
    return string(html), nil
}
