export GOPATH=/root/go_data_analysis
export GOMAXPROCS=1
go build -o http_go httpbenchmark
./http_go
