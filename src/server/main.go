package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "net/http"
)
 
var db *sql.DB
 
func init() {
    db, _ = sql.Open("mysql", "renesola:renes0la.xx@tcp(172.18.22.202:3306)/apollo_eu_erp?charset=utf8")
    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(1000)
    db.Ping()
}
 
func main() {
	init()
    startHttpServer()
}
 
func startHttpServer() {
    http.HandleFunc("/pool", pool)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
 
func pool(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM t_inventory_balance limit 1")
    defer rows.Close()
    checkErr(err)
 
    columns, _ := rows.Columns()
    scanArgs := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
    for j := range values {
        scanArgs[j] = &values[j]
    }
 
    record := make(map[string]string)
    for rows.Next() {
        //将行数据保存到record字典
        err = rows.Scan(scanArgs...)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
    }
 
    fmt.Println(record)
    fmt.Fprintln(w, "finish")
}
 
func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}