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
  // // 1. Create a new file, result.txt (if it doesn't already exist)
  // _, err := os.Create("results.txt")
  // checkError(err)
  //
  // // 2. Get http response with links
  // linksResponse, err := http.Get("https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv")
  // checkError(err)
  // defer linksResponse.Body.Close()
  //
  // // 3. Create a new file, output.csv (if it doesn't already exist) to write results to
  // outputFile, err := os.Create("output.txt")
  // checkError(err)
  //
  // io.Copy(outputFile, linksResponse.Body)

  // Create new file for lines
  linesFile, err := os.Create("somelines.txt")
  checkError(err)

  // 4. Open output file and parse each comma separated value
  outputFile, err := os.Open("output.txt")
  checkError(err)
  defer outputFile.Close()

  reader := csv.NewReader(bufio.NewReader(outputFile))

  // Create an array containing 4 comma separated values
  value := ""
  // line := []string{}

  for {
    separatedValues, err := reader.Read()
    if err == io.EOF {
      break
    }

    length := len(separatedValues)
    writer := bufio.NewWriter(linesFile)

    for i, word := range separatedValues {
      fmt.Println("Index: ", i, "Word: ", word)
      if i == 6 {
        fmt.Println("hello", word)
      }
      if i % 4 == 0 && string(word) != "" {
        value = separatedValues[i + 1]
        fmt.Fprintf(writer, "%v\n", value)
        value = ""
        writer.Flush()
        // contixnue
      } else if (i + 1) == length {
        break
      }
    }
  }
  // manipulateLines()
}

// Do some stuff with that array ***
// func manipulateLines() {
//   // Open resultFile
//   resultFile, err := os.Open("results.txt")
//   checkError(err)
//   defer resultFile.Close()
//
//   // Go through somelines.txt and do all that crap
//   content, err := ioutil.ReadFile("somelines.txt")
//   checkError(err)
//
//   var v interface{}
//   if err := json.Unmarshal([]byte(content), &v); err != nil {
//       fmt.Println(err)
//       fmt.Println("hello")
//       return
//   }
//   // fmt.Println(string(v))
//   fmt.Printf("%#v\n", v)


  // lines := strings.Split(string(content), "\n")
  // fmt.Println(lines)
  // linesFile, err := os.Open("somelines.txt")
  // checkError(err)
  // defer linesFile.Close()

  // Read each line
  // scanner := bufio.NewScanner(linesFile)
  //
  // reader := bufio.NewReader(linesFile)
  // lines, _, err := reader.ReadLine()
  // for i, line range lines {
  //   fmt.Println(lines)
  // }

  // For each line




  // for scanner.Scan(){
  //   line := scanner.Text()
  //   fmt.Println(reflect.TypeOf(line))
  //   break

    // // Check if arr[1] is empty
    // if string(line[0]) != "" {
    //   // If not, get URL from [2]
    //   url := string(line[1])
    //   // 4a. Create new Link object
    //   link := Link{url: url}
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

  // }

// }
//





// Code for writing request and response headers and body
// client := &http.Client{}
//
// request, err := http.NewRequest("GET", "https://raw.githubusercontent.com/ukparliament/ontologies/master/urls.csv", nil)
// response, err := client.Do(request)
//
// request.Write(outputFile)
// response.Write(outputFile)
