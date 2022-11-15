package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	pages, _ := filepath.Glob("./templates/*.page.tmpl")
	for _, page := range pages {
		fmt.Println(filepath.Base(page))
	}
}
