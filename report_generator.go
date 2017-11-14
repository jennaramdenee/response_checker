package main

import (
  // "fmt" 
  "html/template"
  "os"
)

func generateHTMLReport(routesObjectsArray []Route) {
  // Create report HTML file
  reportHTMLFile, err := os.Create("report.html")
  checkError(err)
  defer reportHTMLFile.Close()

  t, err := template.ParseFiles("_report_template.html")
  checkError(err)

  sortedRouteObjects := sortRoutes(routesObjectsArray)

  // fmt.Println("Generating report")
  err = t.Execute(reportHTMLFile, sortedRouteObjects)
}

func sortRoutes(routesObjectsArray []Route) map[int][]Route {
  // Create an empty map, which will hold all routes sorted by status code
  sortedRouteObjects := make(map[int][]Route)

  // For all routes objects
  for _, routeObject := range routesObjectsArray {
    sortedRouteObjects[routeObject.Code] = append(sortedRouteObjects[routeObject.Code], routeObject)
  }
  return sortedRouteObjects
}
