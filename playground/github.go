package main

import (
    "./github.com/hako/durafmt"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"
)

const IssuesURL = "https://api.github.com/search/issues"

func main() {
    result, err := SearchIssues(os.Args[1:])
    if err != nil {
       log.Fatal(err)
    }
    fmt.Printf("%d issues:\n", result.TotalCount)
    fmt.Printf("%6s\t%-9s\t%-9s\t%5s\n", "Number", "Login", "Age", "Title")
    println(strings.Repeat("-",96))
    for _, item := range result.Items {
       fmt.Printf("#%-5d\t%9.9s\t%-9s\t%.55s\n",
           item.Number, item.User.Login, item.Age(), item.Title)
    }
}

type IssuesSearchResult struct {
    TotalCount int `json:"total_count"`
    Items []*Issue
}

type Issue struct {
    Number int
    HTMLURL string `json:"html_url"`
    Title string
    State string
    User *User
    CreatedAt time.Time `json:"created_at"`
    Body string
}

type User struct {
    Login string
    HTMLURL string `json:"html_url"`
}

func (is *Issue) Age() string {
    return durafmt.ParseShort(time.Since(is.CreatedAt)).String()
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
    q := url.QueryEscape(strings.Join(terms, " "))
    resp, err := http.Get(fmt.Sprintf("%s?q=%s", IssuesURL, q))
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("search query failed: %s", resp.Status)
    }

    var result IssuesSearchResult
    defer resp.Body.Close()
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }
    return &result, nil
}