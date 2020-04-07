package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
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

	http.ListenAndServe(":8001", nil)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	resp, error := http.Get("http://localhost:8002/ram")
	if error != nil {
		fmt.Println("Error")
		fmt.Fprintf(w, "{ \"total\" : %s, \"used\": %s, \"free\": %s }", "1", "1", "1")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading response. ", err)
		}
		s := strings.Split(string(body), ",")
		fmt.Fprintf(w, "{ \"total\" : %s, \"used\": %s, \"free\": %s }", s[0], s[1], s[2])
	}
}

func printMemUsage() {
	resp, error := http.Get("http://localhost:8002/ram")
	if error != nil {
		fmt.Println("Error")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading response. ", err)
		}
		s := strings.Split(string(body), ",")
		i, err := strconv.ParseUint(s[0], 10, 64)
		j, err := strconv.ParseUint(s[1], 10, 64)
		addDataToRedis(j, i)
	}
}

func addDataToRedis(value uint64, total uint64) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.187:7001",
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
