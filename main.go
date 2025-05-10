package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

type HaikuResponse struct {
	Haiku string `json:"haiku"`
}

var haikuListEn []string
var haikuListPl []string

func loadHaikuList(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var data []string
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getLanguageFromRequest(r *http.Request) string {
	langHeader := r.Header.Get("Accept-Language")
	fmt.Println("Accept-Language header:", langHeader)
	if strings.HasPrefix(langHeader, "pl") {
		return "pl"
	}
	return "en"
}

func haikuHandler(w http.ResponseWriter, r *http.Request) {
	lang := getLanguageFromRequest(r)
	var haikuList []string
	if lang == "pl" {
		haikuList = haikuListPl
	} else {
		haikuList = haikuListEn
	}
	// Check if the haiku list is empty
	if len(haikuList) == 0 {
		http.Error(w, "Haiku list is empty", http.StatusInternalServerError)
		return
	}
	// Randomly select a haiku from the list
	text := haikuList[rand.Intn(len(haikuList))]

	haiku := HaikuResponse{
		Haiku: text,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(haiku)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	var readingError error
	haikuListPl, readingError = loadHaikuList("haiku_list.json")
	if readingError != nil {
		// Handle error
		panic("Error loading haiku list: " + readingError.Error())
	}
	haikuListEn, readingError = loadHaikuList("haiku_list_en.json")
	if readingError != nil {
		// Handle error
		panic("Error loading haiku list: " + readingError.Error())
	}
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/haiku", haikuHandler)
	fmt.Println("Server is running on port 8080...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		// Handle error
		panic("Error starting server: " + err.Error())
	}
}
