package http

import (
	"avito-flats/internal/domain/entities"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("postgresql", "user:pg4afl@tcp(localhost:5432)/avito-flats")
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, email VARCHAR(255) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL, status INT NOT NULL)")
	if err != nil {
		log.Fatal("Cannot create table: ", err)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func register(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error hashing password"})
		return
	}

	user.Password = hashedPassword

	// Assign status
	if user.Type != 1 {
		user.Type = 0
	}

	query := "INSERT INTO users (email, password, status) VALUES (?, ?, ?)"
	res, err := db.Exec(query, user.Email, user.Password, user.Type)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error inserting user into database"})
		return
	}

	userID, err := res.LastInsertId()
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error retrieving inserted user ID"})
		return
	}

	user.ID = int(userID)
	user.Password = ""

	respondWithJSON(w, http.StatusCreated, user)
}

func login(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	var storedUser entities.User
	err = db.QueryRow("SELECT id, email, password, status FROM users WHERE email = ?", user.Email).Scan(&storedUser.ID, &storedUser.Email, &storedUser.Password, &storedUser.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
		} else {
			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error checking user in database"})
		}
		return
	}

	if !checkPasswordHash(user.Password, storedUser.Password) {
		respondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
		return
	}

	storedUser.Password = ""
	respondWithJSON(w, http.StatusOK, storedUser)
}

func main() {
	initDB()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server running on port 8000")
	log.Fatal(srv.ListenAndServe())
}
