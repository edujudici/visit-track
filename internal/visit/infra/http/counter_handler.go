package http

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	counter     int
	counterLock sync.Mutex
)

func VisitHandler(w http.ResponseWriter, r *http.Request) {
	counterLock.Lock()
	counter++
	counterLock.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Visita registrada!"))
}

func TotalHandler(w http.ResponseWriter, r *http.Request) {
	counterLock.Lock()
	total := counter
	counterLock.Unlock()

	response := map[string]int{"total_visits": total}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
