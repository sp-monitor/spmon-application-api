package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//define the db variable
// var db *sql.DB

func mysqlConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/spmon")
	if err != nil {
		panic(err.Error())
	}

	createTables(db)
	return db
}


func createTables(db *sql.DB) {
	// Check if the hosts table exists
	_, err := db.Query("SELECT 1 FROM hosts LIMIT 1")
	if err != nil {
		// Create the hosts table
		_, err := db.Query(`CREATE TABLE hosts 
		(
			id INT AUTO_INCREMENT PRIMARY KEY, 
			name VARCHAR(255) NOT NULL, 
			ip VARCHAR(255) NOT NULL,
			os VARCHAR(255) NOT NULL 
	  )`)
		if err != nil {
			panic(err.Error())
		}

		// Create the modules table
		_, err = db.Query(`CREATE TABLE modules
		(
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description LONGTEXT
		)`)
		if err != nil {
			panic(err.Error())
		}

		// Create the data table
		_, err = db.Query(`CREATE TABLE data
		(
			id INT AUTO_INCREMENT PRIMARY KEY,
			host_id INT NOT NULL,
			module_id INT NOT NULL,
			data JSON NOT NULL,
			datetime DATETIME NOT NULL,
			FOREIGN KEY (host_id) REFERENCES hosts(id),
			FOREIGN KEY (module_id) REFERENCES modules(id)
		)`)
		if err != nil {
			panic(err.Error())
		}
	}
}