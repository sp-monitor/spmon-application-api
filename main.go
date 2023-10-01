package main

import (
	"net/http"
)



func main() {
	db := mysqlConnect()
	defer db.Close()
	// select all hosts and print the results
	hosts, err := db.Query("SELECT * FROM hosts")
	if err != nil {
		panic(err.Error())
	}
	defer hosts.Close()

	clientRoutes()

	err = http.ListenAndServeTLS(":3443", "./certs/cert.pem", "./certs/cert.key", nil)
	if err != nil {
		panic(err)
	}
}
