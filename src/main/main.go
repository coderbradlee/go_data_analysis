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
    "strings"
    "io/ioutil"
)
type Configuration struct {
    Exec_time string
    Port string
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
    //fmt.Printf("%s\n",configuration.Exec_time)
}
 
func main() {
    ticker := time.NewTicker(time.Minute * 1)
    go func() {
        for _ = range ticker.C {
            t:=time.Now()
            // fmt.Printf("ticked at %v", time.Now())
            // fmt.Printf("%02d:%02d\n",t.Hour(), t.Minute())
            t_conf:=fmt.Sprintf("%s",configuration.Exec_time)
            t_now:=fmt.Sprintf("%02d:%02d",t.Hour(), t.Minute())
            fmt.Printf("%s\n",t_conf)
            fmt.Printf("%s\n",t_now)
            if(strings.EqualFold(t_conf,t_now)){
                fmt.Printf("its time\n")
                request_credit()
                }
        }
    }()
    startHttpServer()
}
func request_credit() {
    httpClient := &http.Client{}
    port:=fmt.Sprintf("%s",configuration.Port)
    var endPoint string = "http://localhost"+port+"/credit"

    req, err := http.NewRequest("GET", endPoint, nil)
    if err != nil {
        log.Fatalf("Error Occured. %+v", err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    // use httpClient to send request
    response, err := httpClient.Do(req)
    if err != nil && response == nil {
        log.Fatalf("Error sending request to API endpoint. %+v", err)
    } else {
        // Close the connection to reuse it
        defer response.Body.Close()

        // Let's check if the work actually is done
        // We have seen inconsistencies even when we get 200 OK response
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatalf("Couldn't parse response body. %+v", err)
        }

        log.Println("Response Body:", string(body))
    }
}
func startHttpServer() {
    port:=fmt.Sprintf("%s",configuration.Port)
    http.HandleFunc("/cost", cost_start)
    http.HandleFunc("/credit", credit_start)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

