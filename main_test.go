package main_test

import (
  "io/ioutil"
  "log"
  "os"
  "reflect"
  "strings"
  "testing"
  "."
)

func readTestFiles(name string) (testResultFileContents []byte) {
  testResultFileContents, err := ioutil.ReadFile(name)
  if err != nil {
    log.Fatal(err)
  }
  return testResultFileContents
}

func createTestOutput(contents string){
  testOutput := []byte(contents)
  err := ioutil.WriteFile("output.txt", testOutput, 0644)
  if err != nil {
    log.Fatal(err)
  }
}

func removeFiles(name string){
  err := os.Remove(name)
  if err != nil {
    log.Fatal(err)
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

func TestRecordLinkStatusOK(t *testing.T) {
  removeFiles("results.txt")
  testUrl := []string{"/"}

  main.RecordLinkStatus(testUrl)

  testResultFileContents := readTestFiles("results.txt")

  actualResult := string(testResultFileContents)
  expectedResult := "/, 200"
  if reflect.DeepEqual(actualResult, expectedResult) {
    t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
  }
}

func TestRecordLinkStatusNotFound(t *testing.T) {
  removeFiles("results.txt")
  testUrls := []string {"/someteststuff", "/someotherstuff"}

  main.RecordLinkStatus(testUrls)

  testResultFileContents := readTestFiles("results.txt")

  actualResult := string(testResultFileContents)
  expectedResult := "/someteststuff, 404"
  if reflect.DeepEqual(actualResult, expectedResult) {
    t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
  }
}

func TestParseLinksNotRoute(t *testing.T) {
  removeFiles("results.txt")
  createTestOutput("On beta,Route,What it is,Page type")
  main.ParseLinks()

  testFileContents := readTestFiles("results.txt")

  body := string(testFileContents)

  if strings.Contains(body, "Route") {
    t.Fatalf("Route should not appear in results.txt file")
  }
}

func TestParseLinks(t *testing.T) {
  removeFiles("results.txt")
  createTestOutput("✓,/search,The search form,Search form")
  main.ParseLinks()

  testFileContents := readTestFiles("results.txt")

  body := string(testFileContents)

  if !strings.Contains(body, "/search") {
    t.Fatalf("/search should appear in results.txt file")
  }
}

func TestParseLinksNotOnBeta(t *testing.T) {
  removeFiles("results.txt")
  createTestOutput(",/mps,Something about MPs,Test MP")
  main.ParseLinks()

  testFileContents := readTestFiles("results.txt")

  body := string(testFileContents)

  if strings.Contains(body, "/mps") {
    t.Fatalf("/mps should not appear in results.txt file")
  }
}

func TestManyParseLinks(t *testing.T) {
  removeFiles("results.txt")
  createTestOutput("✓,/people/a-z,Namespace for navigation of all people,Namespace\n✓,/houses/:house/members,All members of a house ever,Paginated list\n,/mps,Something about MPs,Test MP")
  main.ParseLinks()

  testFileContents := readTestFiles("results.txt")

  body := string(testFileContents)

  if (!strings.Contains(body, "/people/a-z") || !strings.Contains(body, "/houses/1AFu55Hs/members")) && !strings.Contains(body, "/mps") {
    t.Fatalf("Not all routes appearing in results.txt file")
  }
}
