package main

import (
	"fmt"
  "log"
  "encoding/json"
  "time"

  "net/http"
  "net/url"
  "strings"

  "os"
)

/**
 * JSON Marshall/Unmarshall start
 */
type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
  {Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
  {Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
  {Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func jsonEncode(movies []Movie) []byte {
  data, err := json.Marshal(movies)
  if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
  }
  return data
}

func jsonEncodeFormatted(movies []Movie) []byte {
  data, err := json.MarshalIndent(movies, "", " ")
  if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
  }
  return data
}

func jsonDecode(data []byte) []Movie {
  // var movies []struct{ Title string }
  var movies []Movie
  if err := json.Unmarshal(data, &movies); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
  }
  return movies
}
/**
 * JSON Marshall/Unmarshall end
 */

/**
 * JSON Github example decode/encode start
 */
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
  Body string // in Markdown format
}

type User struct {
  Login string
  HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
  q := url.QueryEscape(strings.Join(terms, " "))
  fullUrl := IssuesURL + "?q=" + q

  resp, err := http.Get(fullUrl)
  if err != nil {
    return nil, err
  }

  // We must close resp.Body on all execution paths.
  // (Chapter 5 presents 'defer', which makes this simpler.)
  if resp.StatusCode != http.StatusOK {
    resp.Body.Close()
    return nil, fmt.Errorf("search query failed: %s", resp.Status)
  }

  var result IssuesSearchResult
  if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
    resp.Body.Close()
    return nil, err
  }

  resp.Body.Close()
  return &result, nil
}

func executeSearch(terms []string) {
  result, err := SearchIssues(terms)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%d issues:\n", result.TotalCount)
  for _, item := range result.Items {
    fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
  }
}

/**
 * JSON Github example decode/encode end
 */

func main() {
  // data := jsonEncode(movies)
  // data := jsonEncodeFormatted(movies)
  // structData := jsonDecode(data)
  // // fmt.Printf("%s\n", data)
  // fmt.Println(structData)

  executeSearch(os.Args[1:])
}