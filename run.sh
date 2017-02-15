export GOPATH=$GOPATH:/root/go_data_analysis
export GOMAXPROCS=8
go build -o http_go httpbenchmark
./http_go
