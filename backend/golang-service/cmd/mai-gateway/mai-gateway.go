package main

import (
	"log"
	"manso.live/backend/golang-service/server/mai-gateway/server"
	"os"
)

var (
	Version   string
	BuildTime string
)

func main() {
	st, err := os.Lstat(os.Args[0])
	if err != nil {
		log.Fatalf("os.Lastat error: %v", err)
	}

	log.Printf("===========================================================")
	log.Printf("|  Server Name : %-40s |", st.Name())
	log.Printf("|  Build Time  : %-40s |", BuildTime)
	log.Printf("|  Version     : %-39s |", Version)
	log.Printf("===========================================================")

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
