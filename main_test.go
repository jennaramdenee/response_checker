package main_test

import (
  // "fmt"
  "io/ioutil"
  "log"
  "os"
  "reflect"
  "strings"
  "testing"
  "."
)

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

func createTestResultFiles() (testResultFile *os.File) {
  testResultFile, err := os.Create("test_results.txt")
  if err != nil {
    log.Fatal(err)
  }
  return testResultFile
}

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

func TestRecordLinkStatusOK(t *testing.T) {
  removeFiles("results.txt")
  testUrl := []string{"/"}

  testResultFile := createTestResultFiles()
  defer testResultFile.Close()

  main.RecordLinkStatus(testUrl)

  testResultFileContents := readTestFiles("test_results.txt")

  actualResult := string(testResultFileContents)
  expectedResult := "/, 200"
  if reflect.DeepEqual(actualResult, expectedResult) {
    t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
  }
}

func TestRecordLinkStatusNotFound(t *testing.T) {
  removeFiles("results.txt")
  testUrls := []string {"/someteststuff", "/someotherstuff"}

  testResultFile := createTestResultFiles()
  defer testResultFile.Close()

  main.RecordLinkStatus(testUrls)

  testResultFileContents := readTestFiles("test_results.txt")

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
