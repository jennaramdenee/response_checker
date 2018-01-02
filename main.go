package main

import (
  "bufio"
  "encoding/csv"
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "regexp"
  "strings"
)

const baseUrl = "https://beta.parliament.uk"

func checkError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

type Link struct {
  url   string
  code  int
}

func main(){
  RetrieveLinks()
  ParseLinks()
}

func RetrieveLinks() {
  // Get http response with links
  linksResponse, err := http.Get("https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv")
  checkError(err)
  defer linksResponse.Body.Close()

  body, err := ioutil.ReadAll(linksResponse.Body)
  bodyString := string(body)

  // Create a new file, output.csv (if it doesn't already exist) to write results to
  outputFile, err := os.Create("output.txt")
  checkError(err)
  defer outputFile.Close()

  // Replace carriage return with new line
  var r = regexp.MustCompile("\r")
  s := r.ReplaceAllString(bodyString, "\n")

  // Write response to outputFile
  writer := bufio.NewWriter(outputFile)
  fmt.Fprintf(writer, "%v", s)
}

func ParseLinks() {
  // Open output file and parse each comma separated value
  outputFile, err := os.Open("output.txt")
  checkError(err)
  defer outputFile.Close()

  csvReader := csv.NewReader(outputFile)

  routeArray := []string{}

  for {
    routeInfo, err := csvReader.Read()
    if err == io.EOF {
      break
    }
    checkError(err)

    // Logic to separate out each link
    for i, route := range routeInfo {
      // Find links which are live on beta and not column heading (i.e. the ones we care about)
      if i % 4 == 0 && string(route) != "" && string(routeInfo[i+1]) != "Route" {
        routeArray = append(routeArray, ReplaceResourceId(routeInfo[i + 1])...)
      }
    }
  }

  RecordLinkStatus(routeArray)

}

func RecordLinkStatus(routes []string){
  // Create a new file, result.txt (if it doesn't already exist)
  resultFile, err := os.Create("results.txt")
  checkError(err)
  defer resultFile.Close()

  for _, route := range routes {
    // Create new Link object
    link := Link{url: route}

    // Visit link
    response, err := http.Get(baseUrl + link.url)
    checkError(err)
    defer response.Body.Close()

    // Get response code
    link.code = response.StatusCode

    // Write Response to file
    writer := bufio.NewWriter(resultFile)
    fmt.Fprintf(writer, "%v, %v\n", link.url, link.code)

    writer.Flush()
  }
}

func ReplaceResourceId(route string) []string {
  resourceIdMap := map[string]string {
    ":source": "mnisId",
    ":id": "3299",
    ":person": "TyNGhslR",
    ":contact-point": "wk1atnfh",
    ":constituency": "3WLS0fFd",
    ":party": "DIifZMjq",
    ":house": "1AFu55Hs",
    ":parliament": "b0t56VVL",
    ":place": "E15000006",
    ":postcode": "SW1A 0AA",
    ":medium": "3UJ7otWM",
    ":resource": "S70cUJGM",
  }

  routeArray := []string{}
  var r = regexp.MustCompile(":letters")

  // Replace with valid ids
  for id, value := range resourceIdMap {
    if strings.Contains(route, id) {
      route = strings.Replace(route, id, value, -1)
    }
  }

  // If any route contains :letters, generate 26 routes for each letter
  if strings.Contains(route, ":letters") {
    alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
      "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z" }

    for _, letter := range alphabet {
      s := r.ReplaceAllString(route, letter)
      routeArray = append(routeArray, s)
    }

  } else {
    routeArray = append(routeArray, route)
  }
  return routeArray
}
