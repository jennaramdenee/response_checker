package main

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "regexp"
  "strings"
  // "strconv"

  "gopkg.in/cheggaaa/pb.v1"
)

const baseUrl = "https://beta.parliament.uk"
const routeSource = "https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv"

func RetrieveRouteList() []string {
  // Get http response with links
  linksResponse, err := http.Get(routeSource)
  checkError(err)
  defer linksResponse.Body.Close()

  body, err := ioutil.ReadAll(linksResponse.Body)
  bodyString := string(body)

  // Replace carriage return with new line
  var r = regexp.MustCompile("\r")
  s := r.ReplaceAllString(bodyString, "\n")

  routesReader := strings.NewReader(s)
  return ParseRoutes(routesReader)
}

func RecordRouteStatus(routes []string){
  // Create a new file, result.txt (if it doesn't already exist)
  resultFile, err := os.Create("results.txt")
  checkError(err)
  defer resultFile.Close()

  fmt.Println("Checking route responses\n")

  // Create and start progress bar
  progressBar := pb.StartNew(len(routes))


  for _, route := range routes {
    // Create new Route object
    r := Route{url: route}

    // Create custom http client instance that does not follow redirects
    client := &http.Client{
      CheckRedirect: func(request *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
      },
    }

    // Visit link
    response, err := client.Get(baseUrl + r.url)
    checkError(err)
    defer response.Body.Close()

    // Get response code
    r.code = response.StatusCode

    // Write Response to file
    writer := bufio.NewWriter(resultFile)
    // fmt.Printf("Route: %v, Status Code: %v\n", r.url, r.code)
    fmt.Fprintf(writer, "%v, %v\n", r.url, r.code)

    writer.Flush()

    // Update progress bar
    // statusString := "Route: " + r.url + "Status Code: " + strconv.Itoa(r.code) + "\n"
    // progressBar.Prefix(statusString)
    progressBar.Increment()

  }

  // Finish progress bar
  progressBar.FinishPrint("Finished")
}
