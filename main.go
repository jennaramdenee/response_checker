package main

import (
  "bufio"
  "encoding/csv"
  "fmt"
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "regexp"
  "strings"
)

const baseUrl = "https://beta.parliament.uk"

func checkError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

type Link struct {
  url   string
  code  int
}

func main(){
  RetrieveLinks()
  ParseLinks()
}

func RetrieveLinks() {
  // 2. Get http response with links
  linksResponse, err := http.Get("https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv")
  checkError(err)
  defer linksResponse.Body.Close()

  body, err := ioutil.ReadAll(linksResponse.Body)
  bodyString := string(body)

  // // 3. Create a new file, output.csv (if it doesn't already exist) to write results to
  outputFile, err := os.Create("output.txt")
  checkError(err)
  defer outputFile.Close()

  // Replace carriage return with new line
  var r = regexp.MustCompile("\r")
  s := r.ReplaceAllString(bodyString, "\n")

  // Write response to outputFile
  writer := bufio.NewWriter(outputFile)
  fmt.Fprintf(writer, "%v", s)
}

func ParseLinks() {
  // Create a new file, result.txt (if it doesn't already exist)
  resultFile, err := os.Create("results.txt")
  checkError(err)
  defer resultFile.Close()

  // Open output file and parse each comma separated value
  outputFile, err := os.Open("output.txt")
  checkError(err)
  defer outputFile.Close()

  csvReader := csv.NewReader(outputFile)

  value := ""

  for {
    separatedValues, err := csvReader.Read()
    checkError(err)

    // Logic to separate out each link
    for i, word := range separatedValues {
      if i % 4 == 0 && string(word) != "" && string(separatedValues[i+1]) != "Route" {
        value = ReplaceResourceId(separatedValues[i + 1])
        // Make call to link and record response code
        RecordLinkStatus(value, resultFile)
        value = ""
      }
    }
    if err == io.EOF {
      break
    }
  }
}

func RecordLinkStatus(url string, resultFile *os.File){
  // Create new Link object
  link := Link{url: url}

  // Visit link
  response, err := http.Get(baseUrl + link.url)
  checkError(err)
  defer response.Body.Close()

  // Get response code
  link.code = response.StatusCode

  // Write Response to file
  writer := bufio.NewWriter(resultFile)
  fmt.Fprintf(writer, "%v, %v\n", link.url, link.code)

  writer.Flush()
}

func ReplaceResourceId(word string) string {
  resourceIdMap := map[string]string {
    ":source": "mnisId",
    ":id": "3299",
    ":letters": "g",
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

  // Replace with valid ids
  for id, value := range resourceIdMap {
    if strings.Contains(word, id) {
      word = strings.Replace(word, id, value, -1)
    }
  }
  return word
}
