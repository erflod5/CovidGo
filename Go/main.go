package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/mackerelio/go-osstat/memory"
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
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	} else {
		fmt.Fprintf(w, "{ \"total\" : %d, \"used\": %d, \"free\": %d }", memory.Total/1024/1024, memory.Used/1024/1024, memory.Free/1024/1024)
	}
}

func printMemUsage() {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	addDataToRedis((memory.Used / 1024 / 1024), (memory.Total / 1024 / 1024))
}

func addDataToRedis(value uint64, total uint64) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.LPush("ram", value*100/total).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.LRange("ram", 0, 30).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
