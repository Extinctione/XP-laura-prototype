package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

var users = map[string]User{}

func main() {
	// Serve static frontend
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// API endpoints
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/jobs", jobsHandler)
	http.HandleFunc("/users", usersHandler) // new endpoint

	// Start server
	http.ListenAndServe(":8080", nil)
}

// Signup endpoint
func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil || u.ID == "" || u.Name == "" || u.Email == "" || u.Role == "" {
		http.Error(w, "bad request or missing fields", http.StatusBadRequest)
		return
	}

	users[u.ID] = u
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// Jobs endpoint
func jobsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		jobs := []string{"Tutor", "Designer", "Dog Walker", "Errand Helper"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jobs)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

// Users endpoint
func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
