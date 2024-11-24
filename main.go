package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SwanHtetAungPhyo/user_service/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
}

func main() {
	// Initialize database
	database.DB_INIT()

	// Migrate models
	database.Migration(&User{})

	// Start server
	ServerStart()
}

func ServerStart() {
	mux := http.NewServeMux()
	user := &User{
		Model: gorm.Model{ID: 1},
		Name:  "Swan",
	}

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		body, err := json.Marshal(user)
		if err != nil {
			log.Printf("Failed to marshal user: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	})

	log.Println("Server running on port 5003")
	http.ListenAndServe(":5003", mux)
}
