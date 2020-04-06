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
	http.ListenAndServe(":8001", nil)
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
	fmt.Printf("  User: %.2f\n", percent[0])
	fmt.Printf("  Nice: %.2f\n", percent[1])
	fmt.Printf("   Sys: %.2f\n", percent[2])
	fmt.Printf("  Intr: %.2f\n", percent[3])
	fmt.Printf("  Idle: %.2f\n", percent[4])
	fmt.Printf("States: %.2f\n", percent[5])
	fmt.Fprintf(w, "%.2f", percent[0]+percent[2])
}
