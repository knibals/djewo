package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalln("Le numéro de port est nécessaire")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bonjour le monde")
	})
	log.Printf("Serveur démarré sur http://localhost:%s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
