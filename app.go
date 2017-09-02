package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// How to try: PORT=3000 go run app.go version.go
func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Couldn't start the service, port is not set")
	}

	router := httprouter.New()
	router.GET("/", info)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

func info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	info := struct {
		Version   string `json:"version"`
		Commit    string `json:"commit"`
		BuildTime string `json:"buildTime"`
	}{
		Version:   version,
		Commit:    commit,
		BuildTime: buildTime,
	}

	err := json.NewEncoder(w).Encode(info)

	if err != nil {
		log.Printf("Info handler, error during encode: %v", err)
		http.Error(w, "Something went wrong", http.StatusServiceUnavailable)
	}
}
