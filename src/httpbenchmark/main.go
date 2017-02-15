package main
 
import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "encoding/json"
    "time"
    "strings"
    "flag"
    "runtime"
)
type Configuration struct {
    Exec_time string
    Port string
    Log_name string
}
var configuration Configuration
func init() {
    file, _ := os.Open("src/mainproject/conf.json")
    decoder := json.NewDecoder(file)
    configuration = Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
      fmt.Println("error:", err)
    }
    log_init()
    //fmt.Println(configuration.exec_time) // output: [UserA, UserB]
    //fmt.Printf("%s\n",configuration.Exec_time)
}
func log_init() {
    log_name:=fmt.Sprintf("%s",configuration.Log_name)
    logFileName := flag.String("log", log_name, "Log file name")
    runtime.GOMAXPROCS(runtime.NumCPU())
    flag.Parse()

    //set logfile Stdout
    logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if logErr != nil {
        fmt.Println("Fail to find", *logFile, "cServer start Failed")
        os.Exit(1)
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
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
            // log.Printf("%s\n",t_conf)
            // log.Printf("%s\n",t_now)
            if(strings.EqualFold(t_conf,t_now)){
                log.Printf("its time\n")
                request_credit()
                }
        }
    }()
    startHttpServer()
}

func startHttpServer() {
    port:=fmt.Sprintf("%s",configuration.Port)
    http.HandleFunc("/test", test_start)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
func test_start(w http.ResponseWriter, r *http.Request) {
    
    fmt.Fprintln(w, "finish")
}
