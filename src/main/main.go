package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "net/http"
    "os"
    "encoding/json"
    "time"
)
type Configuration struct {
    Exec_time string
}
var db *sql.DB
var configuration Configuration

func init() {
    db, _ = sql.Open("mysql", "renesola:renes0la.xx@tcp(172.18.22.202:3306)/apollo_eu_erp?charset=utf8")
    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(10)
    db.Ping()
    file, _ := os.Open("src/main/conf.json")
    decoder := json.NewDecoder(file)
    configuration = Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
      fmt.Println("error:", err)
    }
    //fmt.Println(configuration.exec_time) // output: [UserA, UserB]
    fmt.Printf("%s\n",configuration.Exec_time)
}
 
func main() {
    ticker := time.NewTicker(time.Minute * 1)
    go func() {
        for _ = range ticker.C {
            t:=time.Now()
            // fmt.Printf("ticked at %v", time.Now())
            // fmt.Printf("%02d:%02d\n",t.Hour(), t.Minute())
            t_now:=fmt.Sprintf("%02d:%02d\n",t.Hour(), t.Minute())
            if(strings.EqualFold(configuration.Exec_time,t)){
                fmt.Printf("its time")
            }
        }
    }()
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

