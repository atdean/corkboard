package main

import (
    "database/sql"
    "log"
    "runtime"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores

    db := initDB("gouser:gouser@/corkboard?parseTime=true")
    defer db.Close()
}

func initDB(dsn string) *sql.DB {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    if err = db.Ping(); err != nil {
        log.Fatal("Could not establish database connection:\n", err)
    }
    return db
}
