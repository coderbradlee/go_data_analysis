export GOPATH=/root/go_data_analysis
cd src
go get github.com/go-sql-driver/mysql
ab -c 100 -n 1000 'http://localhost:9090/pool'