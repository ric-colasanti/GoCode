package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "./"
	current_directory, err := readDir(path)
	if err != nil {
		panic(err)
	}
	directoryJSON, err := json.Marshal(current_directory)
	fmt.Println(string(directoryJSON))
}

type directory struct {
	Home            string
	Files           []string
	Sub_Directories []string
}

func readDir(path string) (directory, error) {
	abs, _ := filepath.Abs(path)
	current_directory := directory{Home: abs, Files: []string{}, Sub_Directories: []string{}}
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
		defer file.Close()
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

	return current_directory, nil
}
