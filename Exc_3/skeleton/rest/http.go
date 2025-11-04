package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ordersystem/repository"
)

type Server struct {
	store *repository.Store
	mux   *http.ServeMux
}

func NewServer(store *repository.Store) *Server {
	s := &Server{store: store, mux: http.NewServeMux()}
	s.routes()
	return s
}

func (s *Server) routes() {
	// Serve everything in ./frontend at site root:
	//  - http://localhost:3000/ -> frontend/index.html
	//  - http://localhost:3000/app.js -> frontend/app.js, etc.
	s.mux.Handle("/", http.FileServer(http.Dir("frontend")))

	// API endpoints
	s.mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	s.mux.HandleFunc("/drinks", s.handleDrinks)
	s.mux.HandleFunc("/orders", s.handleOrders)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) handleDrinks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	ds, err := s.store.Drinks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, ds)
}

func (s *Server) handleOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		o, err := s.store.TotalledOrders()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, o)
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		drinkID64, _ := strconv.ParseUint(r.Form.Get("drink_id"), 10, 64)
		qty, _ := strconv.Atoi(r.Form.Get("qty"))
		o, err := s.store.CreateOrder(uint(drinkID64), qty)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, o)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
