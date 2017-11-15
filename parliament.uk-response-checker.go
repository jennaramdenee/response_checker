package main

import (
  // "bufio"
  // "encoding/csv"
  // "fmt"
  // "io"
  // "io/ioutil"
  "log"
  "net/http"
  // "strings"
  "os"
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
  _, err := os.Create("results.txt")
  checkError(err)

  // 2. Get http response with links
  linksResponse, err := http.Get("https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv")
  checkError(err)
  defer linksResponse.Body.Close()

  // // 3. Create a new file, output.csv (if it doesn't already exist) to write results to
  outputFile, err := os.Create("output.csv")
  checkError(err)

  io.Copy(outputFile, linksResponse.Body)

  // // 3. Read each line
  // // scanner := bufio.NewScanner(outputFile)
  //
  // // 4. Open file and parse each comma separated value
  // outputFile, err = os.Open("output.csv")
  // reader := csv.NewReader(bufio.NewReader(outputFile))
  //
  // // 5. For each line
  // for {
  //   separatedValue, err := reader.Read()
  //   if err == io.EOF {
  //     break
  //   }
  //
  //   fmt.Println("HELLO")
  //   fmt.Println(line[1:4])
  // }

  // 4a. Create new Link object
  //   link := Link{url: scanner.Text()}
  //
  //   // 4b. Visit link
  //   response, err := http.Get(baseUrl + link.url)
  //   if err != nil {
  //     log.Fatal(err)
  //   }
  //   defer response.Body.Close()
  //
  //   // 4c. Get response code
  //   link.code = response.StatusCode
  //
  //   // 4d. Write Response to file
  //   writer := bufio.NewWriter(resultFile)
  //   fmt.Fprintf(writer, "%v, %v\n", link.url, link.code)
  //
  //   writer.Flush()
  // }
  // if err := scanner.Err(); err != nil {
  //   fmt.Fprintln(os.Stderr, "Issues reading input:", err)
  // }

}



// Code for writing request and response headers and body
// client := &http.Client{}
//
// request, err := http.NewRequest("GET", "https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv", nil)
// response, err := client.Do(request)
//
// request.Write(outputFile)
// response.Write(outputFile)
