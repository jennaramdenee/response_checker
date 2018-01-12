package main

import (
  "encoding/csv"
  "encoding/json"
  "io"
  "io/ioutil"
  "regexp"
  "strings"
)

type Route struct {
  Url   string
  Code  int
}

func ParseRoutes(routesReader io.Reader) []string{
  csvReader := csv.NewReader(routesReader)

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
  return routeArray
}

func ReplaceResourceId(route string) []string {
  resourceMapFile, err := ioutil.ReadFile("./resource_map.json")
  checkError(err)

  var resourceMap map[string]string

  if err := json.Unmarshal(resourceMapFile, &resourceMap); err != nil {
    checkError(err)
  }

  // Set up array to hold all routes
  routeArray := []string{}
  var r = regexp.MustCompile(":letters")

  // Replace with valid ids
  for id, value := range resourceMap {
    if strings.Contains(route, id) {
      route = strings.Replace(route, id, value, -1)
    }
  }

  // If any route contains :letters, generate 26 routes for each letter
  if strings.Contains(route, ":letters") {
    alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
      "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

    for _, letter := range alphabet {
      s := r.ReplaceAllString(route, letter)
      routeArray = append(routeArray, s)
    }

  } else {
    routeArray = append(routeArray, route)
  }
  return routeArray
}
