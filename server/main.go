package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("data.json")
	if err != nil {
		http.Error(w, "Unable to read the JSON file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var data interface{} 
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
