package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Animal struct {
	Animal   string `json:"animal"`
	Category string `json:"category"`
}

var animalDetails = map[string]Animal{
	"1": {Animal: "Crocodile", Category: "Reptile"},
	"2": {Animal: "Tiger", Category: "Mammal"},
	"3": {Animal: "Dog", Category: "Mammal"},
	"4": {Animal: "Snake", Category: "Reptile"},
}

func getAnimalHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	animal, exists := animalDetails[id]
	if !exists {
		http.Error(w, "Animal not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(animal)
	if err != nil {
		http.Error(w, "Failed to Encode Response", http.StatusInternalServerError)
	}
}

func main() {

	port := 8080

	fmt.Printf("Server Starting on port %d\n", port)

	http.HandleFunc("/animal", getAnimalHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
