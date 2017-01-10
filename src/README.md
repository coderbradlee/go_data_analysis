
E:/renesola/svn/apollo/1.0/design/Demo/ERP/WEB-INF/Inventory/index_fm.html
信用和成本相关的页面
E:/renesola/svn/apollo/1.0/design/Demo/ERP/WEB-INF/Inventory/index_om.html
之前销售top10的页面

2017.01.03
go build -o src/main/go_data_analysis main 

t_commercial_invoice找到订单-->payment term 30days-->推导出deadline，用今天时间减去deadline

t_credit_black_list_detail 写入信用相关的表
t_credit_black_list 写入信用相关的表
customer_credit_flow 信用相关数据来源表

信用相关的页面E:\renesola\svn\apollo\1.0\design\Demo\Newshop\WEB-INF\financialmanager


////////////////////////////////////////////////////////////////////////////////
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



t_credit_black_list_detail 写入信用相关的表
t_credit_black_list 写入信用相关的表
t_cost_statistics 库存成本统计表
t_sales_statistics 销售统计表
t_sales_statistics_detail 销售统计明细表


t_inventory_balance 成本相关数据来源表
t_inventory_balance_detail 成本相关数据来源表


