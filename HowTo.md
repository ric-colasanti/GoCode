# Go: How to

## Useful links

[How to write go code from go](https://go.dev/doc/code)

[Video using Go with vscode](https://www.youtube.com/watch?v=1MXIGYrMk80)

[How to use net/http](https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go)

## Create a module
1. ```go mod init name-of-prog```
2. Running ```go mod init``` will create a go.mod file, which is used to manage the module's dependencies and versioning. This file specifies the module's name, version, and dependencies, as well as the Go version that the module is compatible with.
3. When the code is compiled it will have the programme name of the init name. So here it would be name-of-prog

## Adding external packages
1. In the upper directory 
2. ```go get github.com/gorilla/mux```
3. That package will then be added to the ```go.mod``` 

## Go Packages
1. These are the main method of organizing coe into 'modules'
2. All functions in those packages that need to be exposed to the main or other packages must have the first letter of their names capitalized.

## Defer
Defer is a keyword in Golang that is used to delay the execution of a function or statement until the end of the current function. This can be useful for ensuring that resources are cleaned up even if an error occurs, or for performing other tasks that need to be done after the function has finished executing.
```
func openFile() (file *os.File, err error) {
    file, err = os.Open("filename.txt")
    if err != nil {
        return nil, err
    }

    defer file.Close()

    // Do something with the file

    return file, nil
}
```

## Gorilla

The gorilla/mux package is a popular HTTP router and URL matcher for building Go web servers. It implements the http.Handler interface, so it can be used with any net/http server.

- Routing: Gorilla Mux can match requests based on URL host, path, path prefix, schemes, header and query values, HTTP methods or using custom matchers.

```
router := mux.NewRouter()

	// Serve static files
	s := http.StripPrefix("/js/", http.FileServer(http.Dir("./js/")))
	router.PathPrefix("/js/").Handler(s)

	router.HandleFunc("/", home)
	router.HandleFunc("/getFile", directory.GetFile).Methods("GET")
	router.HandleFunc("/getDirectory", directory.GetDirectory).Methods("GET")

	http.ListenAndServe(":5555", router)
```

## net/http 
The net/http package in Golang provides a set of functions and types for creating and using HTTP servers and clients. It is the standard library package for working with HTTP in Go.

```
s := http.StripPrefix("/js/", http.FileServer(http.Dir("./js/")))
http.ListenAndServe(":5555", router)
```
Documentation [net/http](https://pkg.go.dev/net/http)

## Create empty 2D slice
```
dataSlice:= [][]float64{}
```




