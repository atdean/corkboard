package main

import (
	"database/sql"
	"net/http"
)

type corkboardApp struct {
	DB *sql.DB
}

func (app *corkboardApp) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
