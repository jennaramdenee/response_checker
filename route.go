package main

import (
  "encoding/csv"
  "io"
  "os"
  "regexp"
  "strings"
)

type Route struct {
  url   string
  code  int
}

func ParseRoutes() {
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
  RecordRouteStatus(routeArray)
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
