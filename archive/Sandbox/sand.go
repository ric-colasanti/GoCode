package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
)

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

func Home(w http.ResponseWriter, r *http.Request) {
	abs, _ := os.Getwd()
	fmt.Println("home", abs)
	http.ServeFile(w, r, "./index.html")
}

func GetData() (data [][]float64) {
	csvfile, err := os.Open("TestSheet.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(csvfile)
	rawCSVdata, err := reader.Read()
	rawCSVdata, err = reader.Read()
	for rawCSVdata != nil {
		d1, _ := strconv.ParseFloat(rawCSVdata[1], 64)
		d2, _ := strconv.ParseFloat(rawCSVdata[4], 64)
		new_row := []float64{d1, d2}
		data = append(data, new_row)
		rawCSVdata, err = reader.Read()
	}
	fmt.Println(data)
	return
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	data := GetData()
	if data == nil {
		fmt.Println("Error")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	}
}

func main() {
	router := mux.NewRouter()

	// Serve static files
	s := http.StripPrefix("/js/", http.FileServer(http.Dir("./js/")))
	router.PathPrefix("/js/").Handler(s)

	router.HandleFunc("/", Home)
	router.HandleFunc("/getfile", GetFile).Methods("GET")

	fmt.Println("Listening at port:5555")
	openbrowser("http://127.0.0.1:5555")
	http.ListenAndServe(":5555", router)

}
