package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-osstat/memory"
)

func main() {
	go func() {
		for {
			printMemUsage()
			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", indexRoute)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8000", nil)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my webiste")
}

func printMemUsage() {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	log.Printf("memory total: %d bytes\n", memory.Total/1024/1024)
	log.Printf("memory used: %d bytes\n", memory.Used/1024/1024)
	log.Printf("memory cached: %d bytes\n", memory.Cached/1024/1024)
	log.Printf("memory free: %d bytes\n\n", memory.Free/1024/1024)
}
