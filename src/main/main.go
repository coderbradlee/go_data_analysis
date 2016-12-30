package main
 
import (
    "database/sql"
    _"fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "net/http"
)
 
var db *sql.DB
 
func init() {
    db, _ = sql.Open("mysql", "renesola:renes0la.xx@tcp(172.18.22.202:3306)/apollo_eu_erp?charset=utf8")
    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(10)
    db.Ping()
}
 
func main() {
    startHttpServer()
}
 
func startHttpServer() {
    http.HandleFunc("/cost", cost_start)
    http.HandleFunc("/credit", credit_start)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

