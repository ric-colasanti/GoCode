package main

import (
	"fmt"

	"github.com/ric-colasanti/GoCode/csvfileviewer/getdirectory"
)

func main() {
	fmt.Println(getdirectory.ReadDir("./"))
}
