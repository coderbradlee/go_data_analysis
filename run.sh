export GOPATH=$GOPATH:/root/go_data_analysis
export GOMAXPROCS=7
go build -o http_go httpbenchmark
./http_go
