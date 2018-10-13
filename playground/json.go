package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Movie struct {
    Title string
    Year int `json:"released"`
    Color bool `json:"color,omitempty"`
    Actors []string
}

func main() {
    movies := []Movie{
        {"Casablanca",1942,false,[]string{"Humphrey Bogart", "Ingrid Bergman"}},
        {"Cool Hand Luke",1967,true,[]string{"Paul Newman"}},
        {"Bullitt", 1968, true,[]string{"Steve McQueen", "Jacqueline Bisset"}},
    }
    data, err := json.MarshalIndent(movies, ""," ")
    if err != nil {
        log.Fatalf("JSON marshaling failed: %s", err)
    }
    fmt.Printf("%s\n", data)
    var titles []struct{ Title string }
    if err := json.Unmarshal(data, &titles); err != nil {
        log.Fatalf("JSON unmarshaling failed: %s", err)
    }
    fmt.Println(titles)
}
