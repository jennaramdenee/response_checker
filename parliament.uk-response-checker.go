package main

import (
  "bufio"
  // "bytes"
  "encoding/csv"
  // "encoding/json"
  "fmt"
  "io"
  // "io/ioutil"
  "log"
  // "net/http"
  "os"
  // "regexp"
  // "strings"
  // "unicode/utf8"
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
  // 1. Create a new file, result.txt (if it doesn't already exist)
  // _, err := os.Create("results.txt")
  // checkError(err)

  // 2. Get http response with links
  // linksResponse, err := http.Get("https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv")
  // checkError(err)
  // defer linksResponse.Body.Close()

  // body, err := ioutil.ReadAll(linksResponse.Body)
  // bodyString := string(body)

  // // 3. Create a new file, output.csv (if it doesn't already exist) to write results to
  // outputFile, err := os.Create("output.txt")
  // checkError(err)
  // defer outputFile.Close()

  // Replace carriage return with new line
  // var r = regexp.MustCompile("\r")
  // s := r.ReplaceAllString(bodyString, "\n")

  // Write response to outputFile
  // writer := bufio.NewWriter(outputFile)
  // fmt.Fprintf(writer, "%v", s)


  // Create new file for lines
  linesFile, err := os.Create("somelines.txt")
  checkError(err)

  // 4. Open output file and parse each comma separated value
  outputFile, err := os.Open("output.txt")
  checkError(err)
  defer outputFile.Close()

  csvReader := csv.NewReader(outputFile)

  // Create an array containing 4 comma separated values
  value := ""

  for {
    separatedValues, err := csvReader.Read()
    checkError(err)
    if err == io.EOF {
      break
    }

    length := len(separatedValues)
    writer := bufio.NewWriter(linesFile)

    for i, word := range separatedValues {
      fmt.Printf("number: %v Word: %v\n", i, word)
      if i % 4 == 0 && string(word) != "" {
        value = separatedValues[i + 1]
        fmt.Fprintf(writer, "%v\n", value)
        value = ""
        writer.Flush()
      } else if (i + 1) == length {
        break
      }
    }
  }
}


resourceIdMap := map[string]int{
  ":source": "mnisId",
  ":id": "3299",
  ":letters": "g",
  ":person": "TyNGhslR",
  ":contact-point",
  ":constituency": "3WLS0fFd",
  ":party": "DIifZMjq",
  ":house": "1AFu55Hs",
  ":parliament": "b0t56VVL",
  ":postcode": "SW1A 0AA",
  ":medium": "3UJ7otWM",
  ":resource": "S70cUJGM",
}

func replaceResourceId(word string) {
  if
}

// Code for writing request and response headers and body
// client := &http.Client{}
//
// request, err := http.NewRequest("GET", "https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv", nil)
// response, err := client.Do(request)
//
// request.Write(outputFile)
// response.Write(outputFile)
