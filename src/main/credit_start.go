 package main
 
import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
)
/**
 * 初步获取t_inventory_balance里面的数据
 */
type credit_black_list struct {
      credit_black_list_id string
      company_id string
      customer_master_id string
      status int
      total_overdue_days int
      total_overdue_amount float32
      currency_id string
      statistic_beginning_date string
      statistic_ending_date string
      accounting_period_id string
      sort_no int
      createAt string
      createBy string
      dr int
      data_version int
}
var g_credit_black_list credit_black_list
type customer_credit_flow struct {
      company_id string
      customer_master_id string
      transaction_no string
      effective_date string
      expire_date string
      transaction_amount float32
      currency_id string
      credit_days int
      bank_receipt_voucher_id string
      ar_accountant_id string
      balance float32
      transaction_date string
}
type commercial_invoice struct {
       company_id string
       invoice_no string
       invoice_date string
       sales_order_id string
}
func (u *customer_credit_flow) print(){
	fmt.Printf("%s||%s||%s||%s\n", u.company_id,u.customer_master_id,u.currency_id,u.transaction_date)
}
func credit_start(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query(`SELECT 
       company_id
       invoice_no
       invoice_date
       sales_order_id FROM t_commercial_invoice`)
    defer rows.Close()
    checkErr(err)
 
    var records []*commercial_invoice
    for rows.Next() {
        p := new(commercial_invoice)
        if err := rows.Scan(&p.company_id,
          &p.invoice_no,
          &p.invoice_date,
          &p.sales_order_id); err != nil {
            log.Println("sql error")
        }
        records = append(records, p)
    }

    print_flow(records)
    fmt.Fprintln(w, "finish")
}
func credit_start1(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query(`SELECT 
      company_id,
      customer_master_id,
      transaction_no,
      effective_date,
      expire_date,
      transaction_amount,
      currency_id,
      credit_days,
      bank_receipt_voucher_id,
      ar_accountant_id,
      balance,
      transaction_date FROM t_customer_credit_flow`)
    defer rows.Close()
    checkErr(err)
 
    // columns, _ := rows.Columns()
    // scanArgs := make([]interface{}, len(columns))
    // values := make([]interface{}, len(columns))
    // for j := range values {
    //     scanArgs[j] = &values[j]
    // }

    var records []*customer_credit_flow
	for rows.Next() {
	    p := new(customer_credit_flow)
	    if err := rows.Scan(&p.company_id,
          &p.customer_master_id,
          &p.transaction_no,
          &p.effective_date,
          &p.expire_date,
          &p.transaction_amount,
          &p.currency_id,
          &p.credit_days,
          &p.bank_receipt_voucher_id,
          &p.ar_accountant_id,
          &p.balance,
          &p.transaction_date); err != nil {
	    	fmt.Println("55")
	    }
	    records = append(records, p)
	}

 	print_flow(records)
    // copy(records) 
    // pad_product_category_id(g_insert_data)
    // pad_statistic_time(g_insert_data)
    // pad_createAt(g_insert_data)
    // pad_currency_id(g_insert_data)
    fmt.Fprintln(w, "finish")
}
func print_flow( records []*customer_credit_flow) {
	for _,i:=range records{
		i.print()
	}
}
func print_invoice( records []*commercial_invoice) {
    for _,i:=range records{
        i.print()
    }
}
