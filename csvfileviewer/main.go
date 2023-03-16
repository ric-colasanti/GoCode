package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ric-colasanti/GoCode/csvfileviewer/getdirectory"
)

// func main() {

// }

func Home(w http.ResponseWriter, r *http.Request) {
	abs, _ := os.Getwd()
	fmt.Println("home", abs)
	http.ServeFile(w, r, "./index.html")
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		fmt.Println("Linux", url)
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Println(err)
	}

}

func CSVFileToMap(filePath string) (returnMap []map[string]string) {

	csvfile, err := os.Open(filePath)
	if err != nil {
		return nil
	}

	defer csvfile.Close()
	csvfile.Seek(0, 0)
	reader := csv.NewReader(csvfile)

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil
	}

	header := []string{} // holds first row (header)
	for lineNum, record := range rawCSVdata {

		// for first row, build the header slice
		if lineNum == 0 {
			for i := 0; i < len(record); i++ {
				header = append(header, strings.TrimSpace(record[i]))
			}
		} else {
			// for each cell, map[string]string k=header v=value
			line := map[string]string{}
			for i := 0; i < len(record); i++ {
				line[header[i]] = record[i]
			}
			returnMap = append(returnMap, line)
		}
	}
	return returnMap
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	dir, ok := r.URL.Query()["directory"]
	if ok {
		data := CSVFileToMap(dir[0])
		if data == nil {
			fmt.Println("Error")
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}

func GetDirectory(w http.ResponseWriter, r *http.Request) {
	dir, ok := r.URL.Query()["directory"]
	if ok {
		current_directory, err := getdirectory.ReadDir(dir[0])
		if err != nil {
			fmt.Println("Error")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(current_directory)
	}
}

func main() {
	router := mux.NewRouter()

	// Serve static files
	s := http.StripPrefix("/js/", http.FileServer(http.Dir("./js/")))
	router.PathPrefix("/js/").Handler(s)

	router.HandleFunc("/", Home)
	router.HandleFunc("/getdirectory", GetDirectory).Methods("GET")
	router.HandleFunc("/getfile", GetFile).Methods("GET")

	fmt.Println("Listening at port:5555")
	openbrowser("http://127.0.0.1:5555")
	http.ListenAndServe(":5555", router)
}
