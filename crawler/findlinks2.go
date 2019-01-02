package main

import (
    "bytes"
    "errors"
    "fmt"
    "golang.org/x/net/html"
    "io"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        links, err := findLinks(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
            continue
        }
        for _, link := range links {
            fmt.Println(link)
        }
    }
}

func DownloadPages(urls []string) {
    done := make(chan []byte, len(urls))
    errch := make(chan error, len(urls))
    for _, url := range urls {
        go func (url string) {

        }(url)
    }
}

func downloadPage(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil { return nil, err }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, resp.Status)
    }
    var data bytes.Buffer
    _, err = io.Copy(&data, resp.Body)
    if err != nil { return nil, err }
    return string(data.Bytes()), nil
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil { return nil, err }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
    return visitLinks(nil, doc), nil
}

// visit appends to links each link found in n and returns the result.
func visitLinks(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                links = append(links, a.Val)
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }
    return links
}
