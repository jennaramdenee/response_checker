package main_test

import (
  // "fmt"
  "reflect"
  "strings"
  "testing"
  "."
)

// Request.go
// func TestRecordRouteStatusOK(t *testing.T) {
//   removeFiles("results.txt")
//   testUrl := []string{"/"}
//
//   main.RecordRouteStatus(testUrl)
//
//   testResultFileContents := readTestFiles("results.txt")
//
//   actualResult := string(testResultFileContents)
//   expectedResult := "/, 200"
//   if reflect.DeepEqual(actualResult, expectedResult) {
//     t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
//   }
// }
//
// func TestRecordRouteStatusNotFound(t *testing.T) {
//   removeFiles("results.txt")
//   testUrls := []string {"/someteststuff", "/someotherstuff"}
//
//   main.RecordRouteStatus(testUrls)
//
//   testResultFileContents := readTestFiles("results.txt")
//
//   actualResult := string(testResultFileContents)
//   expectedResult := "/someteststuff, 404"
//   if reflect.DeepEqual(actualResult, expectedResult) {
//     t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
//   }
// }

// Route.go
func TestParseInvalidRoute(t *testing.T) {
  testRoutesReader := strings.NewReader("On beta,Route,What it is,Page type")
  testRoutesArray := main.ParseRoutes(testRoutesReader)

  if len(testRoutesArray) != 0 {
    t.Fatalf("Route heading should not be considered a valid route")
  }
}

func TestParseValidRoute(t *testing.T) {
  testRoutesReader := strings.NewReader("✓,/search,The search form,Search form")
  actualResult := main.ParseRoutes(testRoutesReader)
  expectedResult := []string{"/search"}

  if !reflect.DeepEqual(actualResult, expectedResult) {
    t.Fatalf("Routes on beta should appear as valid route")
  }
}

func TestParseRoutesNotOnBeta(t *testing.T) {
  testRoutesReader := strings.NewReader(",/mps,Something about MPs,Test MP")
  actualResult := main.ParseRoutes(testRoutesReader)
  expectedResult := []string{}

  if !reflect.DeepEqual(actualResult, expectedResult) {
    t.Fatalf("Routes not on beta should not appear as a valid route")
  }
}

// test helper method
func contains(arr []string, route string) bool {
  for _, r := range arr {
    if r == route {
      return true
    }
  }
  return false
}

func TestManyParseRoutes(t *testing.T) {
  testRoutesReader := strings.NewReader("✓,/people/a-z,Namespace for navigation of all people,Namespace\n✓,/houses/:house/members,All members of a house ever,Paginated list\n,/mps,Something about MPs,Test MP")
  actualResult := main.ParseRoutes(testRoutesReader)

  if (!contains(actualResult, "/people/a-z") || !contains(actualResult, "/houses/1AFu55Hs/members")) && !contains(actualResult, "/mps") {
    t.Fatalf("Not all routes appearing in valid route array")
  }
}

func TestReplaceResourceId(t *testing.T){
  expectedResult := "/people/lookup?source=mnisId&id=3299"
  actualResult := main.ReplaceResourceId("/people/lookup?source=:source&id=:id")
  if actualResult[0] != expectedResult {
    t.Fatalf("Expected %s but got %s", expectedResult, actualResult[0])
  }
}

func TestNotReplaceResourceId(t *testing.T){
  expectedResult := "test/hello/:world"
  actualResult := main.ReplaceResourceId("test/hello/:world")
  if actualResult[0] != expectedResult {
    t.Fatalf("Expected %s but got %s", expectedResult, actualResult[0])
  }
}

func TestLettersReplaceResourceId(t *testing.T){
  expectedResult := 26
  actualResult := main.ReplaceResourceId("test/:letters")

  if expectedResult != len(actualResult) {
   t.Fatalf("Wrong number of results returned, got %v, expected %v", len(actualResult), expectedResult)
  }

  if actualResult[0] != "test/a" {
   t.Fatalf("Incorrect URL formed, got %v, expected %v", actualResult[0], "test/a")
  }

  if actualResult[25] != "test/z" {
   t.Fatalf("Incorrect URL formed, got %v, expected %v", actualResult[0], "test/z")
  }
}
