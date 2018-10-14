package github

import (
    "fmt"
    "html/template"
    "io"
    "log"
    "os"
)

const defaultTemplate = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.DaysAgo}} ago
{{end}}`

// SimpleReport prints a list of retrieved issues in tabular format.
func (isr *IssuesSearchResult) SimpleReport() {
    fmt.Printf("Total issues: %d\n", isr.TotalCount)
    fmt.Printf("%6s\t%9s\t%5s\n", "Number", "Login", "Title")
    for _, item := range isr.Items {
        fmt.Printf("#%-5d\t%9.9s\t%.55s\n",
            item.Number, item.User.Login, item.Title)
    }
}

func (isr *IssuesSearchResult) FormattedReport() {
    isr.TemplateReport(os.Stdout, defaultTemplate)
}

func (isr *IssuesSearchResult) TemplateReport(io io.Writer, templ string) {
    report, err := template.New("report").Parse(templ)
    if err != nil {
        log.Fatal(err)
    }
    if err := report.Execute(io, isr); err != nil {
        log.Fatal(err)
    }
}