 package main
 
import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
)
/**
 * 初步获取t_inventory_balance里面的数据
 */
type inventory_balance struct {
    company_id string
 	accounting_period_id string
 	item_master_id string
 	in_stock_balance int32
 	uom_id string
 	refresh_date string
}
func (u *inventory_balance) print(){
	fmt.Printf("%s||%s||%s||%d||%s||%s\n", u.company_id,u.accounting_period_id,u.item_master_id,u.in_stock_balance,u.uom_id,u.refresh_date)
}
func cost_start(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT company_id,accounting_period_id,item_master_id,in_stock_balance,uom_id,refresh_date FROM t_inventory_balance")
    defer rows.Close()
    checkErr(err)
 
    // columns, _ := rows.Columns()
    // scanArgs := make([]interface{}, len(columns))
    // values := make([]interface{}, len(columns))
    // for j := range values {
    //     scanArgs[j] = &values[j]
    // }

    var records []*inventory_balance
	for rows.Next() {
	    p := new(inventory_balance)
	    if err := rows.Scan(&p.company_id, &p.accounting_period_id, &p.item_master_id, &p.in_stock_balance, &p.uom_id, &p.refresh_date); err != nil {
	    	fmt.Println("55")
	    }
	    records = append(records, p)
	}

 	// print(records)
    copy(records) 
    pad_product_category_id(g_insert_data)
    pad_statistic_time(g_insert_data)
    pad_createAt(g_insert_data)
    pad_currency_id(g_insert_data)
    fmt.Fprintln(w, "finish")
}
func print( records []*inventory_balance) {
	for _,i:=range records{
		i.print()
	}
}
func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}
