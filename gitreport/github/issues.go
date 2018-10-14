package github

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
    "time"
)

const IssuesURL = "https://api.github.com/search/issues"

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

func (is *Issue) DaysAgo() string {
    days := int(time.Since(is.CreatedAt).Hours() / 24)
    return fmt.Sprintf("%d days", days)
}

// Search iIssues returns list of issues for a specific GitHub's repository.
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
