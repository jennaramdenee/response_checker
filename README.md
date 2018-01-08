## Response Checker
Response Checker is a Go script designed to retrieve a list of working routes and check how these are responding, reporting back all URLs that have been visited, along with their status code.

### Requirements
Response Checker requires [Go](https://golang.org/doc/install) to be installed.

### Usage
1. Set up this repository
```kernal
git clone https://github.com/jennaramdenee/response_checker
cd response_checker
```

2. Set base URL in `src/main.go` file

3. Run the script
```kernal
go run main.go
```

4. Two files will be created in your directory
  * `output.txt` - list of all routes
  * `results.txt` - list of URLs that have been visited, along with their status code

### Testing
Tests can be run by using the following command:
```kernal
go test -v
```
