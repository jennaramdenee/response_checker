package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "regexp"
  "strings"

  "gopkg.in/cheggaaa/pb.v1"
)

// const baseUrl = "http://varnish-lb-327117095.eu-west-1.elb.amazonaws.com"
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

func RecordRouteStatus(routes []string) []Route{
  fmt.Println("Checking route responses\n")

  // Create and start progress bar
  progressBar := pb.StartNew(len(routes))

  // Create array for Route objects
  routesObjectsArray := []Route{}

  for _, route := range routes {
    // Create new Route object
    r := Route{Url: route}

    // Create custom http client instance that does not follow redirects
    client := &http.Client{
      CheckRedirect: func(request *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
      },
    }

    // Visit link
    response, err := client.Get(baseUrl + r.Url)
    checkError(err)

    // Get response code
    r.Code = response.StatusCode

    // Add new Route object to array
    routesObjectsArray = append(routesObjectsArray, r)

    // Update progress bar
    // statusString := "Route: " + r.url + "Status Code: " + strconv.Itoa(r.code) + "\n"
    // progressBar.Prefix(statusString)
    progressBar.Increment()

    defer response.Body.Close()
  }
  // Finish progress bar
  progressBar.FinishPrint("Finished")

  // Generate report
  generateHTMLReport(routesObjectsArray)
  return routesObjectsArray
}
