package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"sync"
	"time"
)

const authToken = "supersecrettoken"

type Device struct{
	ID string `json:"id"`
	Status string `json:"status"`
	LastSeen time.Time 	`json:"last_seen"`
	Recovered bool `json: "recovered"`


}

var (
	deviceStore = make(map[string]*Device)
	storeMutex sync.Mutex
)


func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer "+authToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}


func updateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var d Device
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	d.LastSeen = time.Now()
	d.Recovered = false

	storeMutex.Lock()
	deviceStore[d.ID] = &d
	storeMutex.Unlock()

	fmt.Fprintf(w, "Device %s updated successfully", d.ID)
}


func statusHandler(w http.ResponseWriter, r *http.Request) {
	storeMutex.Lock()
	defer storeMutex.Unlock()

	devices := make([]*Device, 0, len(deviceStore))
	for _, d := range deviceStore {
		devices = append(devices, d)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}

func monitorAvailability() {
	for {
		time.Sleep(5 * time.Second)
		storeMutex.Lock()
		for _, d := range deviceStore {
			if time.Since(d.LastSeen) > 10*time.Second && !d.Recovered {
				d.Status = "recovered"
				d.Recovered = true
				fmt.Printf("Device %s marked as recovered due to inactivity.\n", d.ID)
			}
		}
		storeMutex.Unlock()
	}
}

func main() {
	go monitorAvailability()

	http.HandleFunc("/device/update", authMiddleware(updateDeviceHandler))
	http.HandleFunc("/device/status", statusHandler)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil))

}