package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var addr = flag.String("addr", ":8080", "http service address")
var repo = flag.String("repo", ".", "repository to expose")
var ext = flag.String("ext", ".*", "file extension to filter")

func main() {
	flag.Parse()
	http.Handle("/files", http.HandlerFunc(lsFiles))
	http.Handle("/folders", http.HandlerFunc(lsFolders))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func lsFiles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ls(*repo, false)); err != nil {
		log.Fatal(err)
	}
}

func lsFolders(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ls(*repo, true)); err != nil {
		log.Fatal(err)
	}
}

func ls(rep string, onlyFolders bool) []string {
	var result []string
	files, err := ioutil.ReadDir(rep)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			// ignore hidden files
			continue
		}
		if file.IsDir() {
			if onlyFolders {
				result = append(result, filepath.Clean(rep+"/"+file.Name()))
			}
			// recurse on directory
			result = append(result, ls(rep+"/"+file.Name(), onlyFolders)...)
		} else if !onlyFolders && (*ext == ".*" || filepath.Ext(file.Name()) == *ext) {
			// only return files of given extension
			result = append(result, filepath.Clean(rep+"/"+file.Name()))
		}
	}
	return result
}
