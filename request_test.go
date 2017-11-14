package main_test

import (
  "reflect"
  "strings"
  "testing"
  "."
)

// TODO: Does not currently implement tests for RetrieveRouteList method
// Request.go
func TestRecordRouteStatusOK(t *testing.T) {
  testRoutes := []string{"/"}
  actualResult := main.RecordRouteStatus(testRoutes)
  expectedResult := []main.Route{ main.Route{Url: "/", Code: 200} }

  if !reflect.DeepEqual(actualResult, expectedResult) {
    t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
  }
}

func TestRecordRouteStatusNotFound(t *testing.T) {
  testRoutes := []string {"/someteststuff", "/someotherstuff"}
  actualResult := main.RecordRouteStatus(testRoutes)

  expectedResult := []main.Route {
    main.Route{Url: "/someteststuff", Code: 404},
    main.Route{Url: "/someotherstuff", Code: 404},
  }

  if !reflect.DeepEqual(actualResult, expectedResult) {
    t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
  }
}

func TestRecordRouteStatusRedirect(t *testing.T) {
  testRoutes := []string {"/people/lookup?source=mnisId&id=3299"}
  actualResult := main.RecordRouteStatus(testRoutes)

  expectedResult := []main.Route {
    main.Route{Url: "/people/lookup?source=mnisId&id=3299", Code: 302},
  }

  if !reflect.DeepEqual(actualResult, expectedResult) {
    t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
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
