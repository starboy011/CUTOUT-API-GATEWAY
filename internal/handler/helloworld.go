package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve database connection details from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")
	fmt.Print("Rishav", dbUser, dbPassword, dbName, dbHost, dbPort, sslMode)
	// Check if any environment variables are missing
	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" || sslMode == "" {
		http.Error(w, "Missing required environment variables", http.StatusInternalServerError)
		return
	}

	// PostgreSQL connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to connect to database: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error connecting to database: %v", err), http.StatusInternalServerError)
		return
	}

	// Query the users table
	rows, err := db.Query("SELECT firstname, lastname, shopid, shopname FROM users")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	// Iterate through the result set
	var users []string
	for rows.Next() {
		var firstname, lastname, shopid, shopname string
		if err := rows.Scan(&firstname, &lastname, &shopid, &shopname); err != nil {
			http.Error(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}
		users = append(users, fmt.Sprintf("Firstname: %s, Lastname: %s, ShopID: %s, ShopName: %s", firstname, lastname, shopid, shopname))
	}
	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error iterating rows: %v", err), http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "text/plain")
	for _, user := range users {
		fmt.Fprintln(w, user)
	}
}
