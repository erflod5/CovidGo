package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mackerelio/go-osstat/memory"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	http.HandleFunc("/ram", ramRoute)
	http.HandleFunc("/cpu", cpuRoute)
	http.ListenAndServe(":8002", nil)
}

func ramRoute(w http.ResponseWriter, r *http.Request) {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	} else {
		fmt.Fprintf(w, "%d,%d,%d", memory.Total/1024/1024, memory.Used/1024/1024, memory.Free/1024/1024)
	}
}

func cpuRoute(w http.ResponseWriter, r *http.Request) {
	percent, _ := cpu.Percent(time.Second, true)
	fmt.Fprintf(w, "%.2f", percent[0])
}
