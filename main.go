package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores

	// TODO :: Break the DSN out into config files / CLI args
	db, err := initDB("gouser:gouser@/corkboard?parseTime=true")
	if err != nil {
		log.Fatal("Could not establish database connection:\n", err)
	}
	defer db.Close()

	app := &corkboardApp{
		DB: db,
	}

	router := initRouter(app)
	http.ListenAndServe(":3000", router)
}

func initDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func initRouter(app *corkboardApp) http.Handler {
	router := chi.NewRouter()
	router.Get("/", app.handleIndex)
	return router
}
