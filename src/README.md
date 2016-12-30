
2016.12.30
t_cost_statistics 缺少类似t_purchase_cost的表，以下几个字段无法取得：
total_stock_quantity int32
total_cost_amount int32
last_month_unit_price int32
currency_id string

下面先做credit相关
curl -X GET 'http://localhost:9090/cost'
curl -X GET 'http://localhost:9090/credit'
按过期天数排序
//////////////////////////////////////////////////
export GOPATH=/root/go_data_analysis
cd src
go get github.com/go-sql-driver/mysql
ab -c 100 -n 1000 'http://localhost:9090/pool'

2016.12.29
原型位置：file:///E:/renesola/svn/apollo/1.0/design/Demo/ERP/WEB-INF/Inventory/index_fm.html



t_credit_black_list_detail
t_credit_black_list
t_cost_statistics

t_inventory_balance
t_inventory_balance_detail


