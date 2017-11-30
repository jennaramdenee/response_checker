package main_test

import (
  "io/ioutil"
  "log"
  "os"
  "testing"
  "."
)

func TestReplaceResourceId(t *testing.T){
  expectedResult := "test/mnisId/hello/g"
  actualResult := main.ReplaceResourceId("test/:source/hello/:letters")
  if actualResult != expectedResult {
    t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
  }
}

func TestNotReplaceResourceId(t *testing.T){
  expectedResult := "test/hello/world"
  actualResult := main.ReplaceResourceId("test/hello/world")
  if actualResult != expectedResult {
    t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
  }
}

func createTestFiles() (testResultFile *os.File) {
  testResultFile, err := os.Create("test_results.txt")
  if err != nil {
    log.Fatal(err)
  }
  return testResultFile
}

func readTestFiles() (testResultFileContents []byte) {
  testResultFileContents, err := ioutil.ReadFile("test_results.txt")
  if err != nil {
    log.Fatal(err)
  }
  return testResultFileContents
}

func TestRecordLinkStatusOK(t *testing.T) {
  testUrl := "/"

  testResultFile := createTestFiles()
  defer testResultFile.Close()

  main.RecordLinkStatus(testUrl, testResultFile)

  testResultFileContents := readTestFiles()

  actualResult := string(testResultFileContents)
  expectedResult := "/, 200"
  if actualResult != expectedResult {
    t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
  }
}

func TestRecordLinkStatusNotFound(t *testing.T) {
  testUrl := "/someteststuff"

  testResultFile := createTestFiles()
  defer testResultFile.Close()

  main.RecordLinkStatus(testUrl, testResultFile)

  testResultFileContents := readTestFiles()

  actualResult := string(testResultFileContents)
  expectedResult := "/someteststuff, 404"
  if actualResult != expectedResult {
    t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
  }
}

func TestParseLinks() {
  
}
