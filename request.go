package main

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "regexp"
)

const baseUrl = "https://beta.parliament.uk"
const routeSource = "https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv"

func RetrieveRouteList() {
  // Get http response with links
  linksResponse, err := http.Get(routeSource)
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

func RecordRouteStatus(routes []string){
  // Create a new file, result.txt (if it doesn't already exist)
  resultFile, err := os.Create("results.txt")
  checkError(err)
  defer resultFile.Close()

  fmt.Println("Checking route responses\n")

  for _, route := range routes {
    // Create new Route object
    r := Route{url: route}

    // Visit link
    response, err := http.Get(baseUrl + r.url)
    checkError(err)
    defer response.Body.Close()

    // Get response code
    r.code = response.StatusCode

    // Write Response to file
    writer := bufio.NewWriter(resultFile)
    fmt.Printf("Route: %v, Status Code: %v\n", r.url, r.code)
    fmt.Fprintf(writer, "%v, %v\n", r.url, r.code)

    writer.Flush()
  }
}
