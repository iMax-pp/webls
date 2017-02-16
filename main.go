package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var addr = flag.String("addr", ":8080", "http service address")
var repo = flag.String("repo", ".", "repository to expose")
var ext = flag.String("ext", ".*", "file extension to filter")

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(lsRepo))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func lsRepo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ls(*repo)); err != nil {
		log.Fatal(err)
	}
}

func ls(rep string) []string {
	var result []string
	files, err := ioutil.ReadDir(rep)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			// recurse on directory
			result = append(result, ls(rep+"/"+file.Name())...)
		} else if *ext == ".*" || filepath.Ext(file.Name()) == *ext {
			// only return files of given extension
			result = append(result, filepath.Clean(rep+"/"+file.Name()))
		}
	}
	return result
}
