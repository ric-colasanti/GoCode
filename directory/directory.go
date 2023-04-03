package directory

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type directoryStruct struct {
	Home            string
	Files           []string
	Sub_Directories []string
}

type dataStruct struct {
	Header []string
	Data   [][]float64
}

func readDir(path string) (directoryStruct, error) {
	abs, _ := filepath.Abs(path)
	current_directory := directoryStruct{Home: abs, Files: []string{}, Sub_Directories: []string{}}
	file, err := os.Open(path)
	if err != nil {
		return current_directory, err
	}
	defer file.Close()
	names, _ := file.Readdirnames(0)
	for _, name := range names {
		filePath := fmt.Sprintf("%v/%v", path, name)
		file, err := os.Open(filePath)
		if err != nil {
			return current_directory, err
		}
		fileInfo, err := file.Stat()
		if err != nil {
			return current_directory, err
		}
		if fileInfo.IsDir() {
			current_directory.Sub_Directories = append(current_directory.Sub_Directories, fileInfo.Name())
		} else {
			current_directory.Files = append(current_directory.Files, fileInfo.Name())
		}
	}
	defer file.Close()
	return current_directory, err
}

func GetDirectory(w http.ResponseWriter, r *http.Request) {
	dir, ok := r.URL.Query()["directory"]
	if ok {
		current_directory, err := readDir(dir[0])
		if err != nil {
			fmt.Println("Error")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(current_directory)
	}
}

func getDataNub(filePath string) (data dataStruct) {
	dataSlice := [][]float64{}
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	lineCounter := 0
	reader := csv.NewReader(csvFile)
	header, _ := reader.Read()
	rawCSVdata, _ := reader.Read()
	for (rawCSVdata != nil) && (lineCounter < 5) {
		newRow := []float64{}
		for _, element := range rawCSVdata {
			element, _ := strconv.ParseFloat(element, 64)
			newRow = append(newRow, element)
		}
		dataSlice = append(dataSlice, newRow)
		rawCSVdata, _ = reader.Read()
		lineCounter++
	}
	return dataStruct{header, dataSlice}
}

func GetFileNub(w http.ResponseWriter, r *http.Request) {
	dir, ok := r.URL.Query()["directory"]
	if ok {
		data := getDataNub(dir[0])
		if data.Data == nil {
			fmt.Println("Error")
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
