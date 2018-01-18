package main_test

import (
  "reflect"
  "strings"
  "testing"
  "."
)

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
