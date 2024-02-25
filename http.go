package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	fileName string = "./data.json"
	PORT     string = ":5000"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/getall", getAll)
	http.HandleFunc("/getone/", getOne)
	fmt.Println("server running on port ", PORT)
	http.ListenAndServe(PORT, nil)
}
func getOne(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}
	var data string = ""
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(w, "server side error")
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data += scanner.Text()
	}
	var result []map[string]interface{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		http.Error(w, "server side error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	if id < 0 || id >= len(result) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(result[id])
	if err != nil {
		http.Error(w, "server side error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.Write(jsonData)
}
func getAll(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(w, "server side error")
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintf(w, scanner.Text())
	}
}
