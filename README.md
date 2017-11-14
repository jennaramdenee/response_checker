## Response Checker
[Response Checker](https://github.com/jennaramdenee/response_checker) is a Go script designed to retrieve a list of working routes and check how these are responding, reporting back all URLs that have been visited, along with their status code.

### Requirements
Response Checker requires [Go](https://golang.org/doc/install) to be installed.

### How it Works
[Response Checker](https://github.com/jennaramdenee/response_checker) executes five main steps:
1. Retrieve a list of parliament routes
2. Extract only the routes live on Beta
3. Replace ID placeholders with real IDs
4. Make a request to each route, recording each route that has been visited, along with their status code
5. Generate a HTML report with the above results

### Usage
1. Set up this repository
```kernal
git clone https://github.com/jennaramdenee/response_checker
cd response_checker
```

2. Set base URL in `src/request.go` file

3. Build the executable
```kernal
go build .
```

4. This will generate an executable binary in your current directory, `response-checker`

5. Run the executable binary
```kernal
./response-checker
```

### Testing
Tests can be run by using the following command:
```kernal
go test -v
```

### Caveats
Currently, this script only supports the latin alphabet and not all Unicode characters.
