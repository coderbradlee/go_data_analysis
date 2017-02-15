export GOPATH=$GOPATH:./
export GOMAXPROCS=8
go build -o http_go httpbenchmark
./http_go
