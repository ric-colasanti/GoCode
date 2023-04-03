package main

import (
	"fmt"
	"go-test/directory"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func openBrowser(url string) {
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
func main() {
	router := mux.NewRouter()

	// Serve static files
	s := http.StripPrefix("/js/", http.FileServer(http.Dir("./js/")))
	router.PathPrefix("/js/").Handler(s)

	router.HandleFunc("/", home)
	router.HandleFunc("/getFile", directory.GetFileNub).Methods("GET")
	router.HandleFunc("/getDirectory", directory.GetDirectory).Methods("GET")

	fmt.Println("Listening at port:5555")
	openBrowser("http://127.0.0.1:5555")
	http.ListenAndServe(":5555", router)
}
