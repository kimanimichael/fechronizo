package main

import (
	"encoding/json"
	"log"
	"net/http"
	
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Server side error", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json response from: %v", payload)
		/* send an HTTP response header with an internal server error status code*/
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	/* write the data itself*/
	w.Write(dat)
}