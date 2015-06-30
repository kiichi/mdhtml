package main

import (
	"flag"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	src := flag.Arg(0)
	fmt.Println(src)
	//if src == "" {
	//		src = "contents"
	//	}

	src = "contents"
	err := filepath.Walk(src, ProcDir)

	if err != nil {
		fmt.Println("ERROR: Dir Search Problem")
		return
	}

}

// 1. check timestamp or CRS/MD5 fingerpring
// 2. If changed process it
func ProcDir(path string, f os.FileInfo, err error) error {
	fmt.Println(path, filepath.Ext(path), f.Name(), f.IsDir(), f.ModTime())
	if strings.HasSuffix(f.Name(), ".md") && !f.IsDir() {
		Convert(path)
	}

	return nil
}

func Convert(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("ERROR : Cannot read file")
		return
	}
	output := blackfriday.MarkdownCommon(content)
	html := bluemonday.UGCPolicy().SanitizeBytes(output)
	fmt.Println(string(html))
}
