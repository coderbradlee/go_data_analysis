 package main
 
import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "log"
)
/**
 * 初步获取t_inventory_balance里面的数据
 */
type t_account struct {
      account_id string
      name string
      display_name string
      note string 
}
type type_t_account []*t_account
var g_t_account type_t_account


func (u *t_account) print(){
    fmt.Printf("%s||%s||%s||%s\n", u.account_id,u.name,u.display_name,u.note)
}

func start(w http.ResponseWriter, r *http.Request) {
    //首先获取公司id，让后按公司id进行下面的工作
    rows, err := db.Query(`SELECT 
       account_id,
      name,
      display_name,
      note FROM t_account`)
    defer rows.Close()
    checkErr(err)
    
    for rows.Next() {
        p := new(t_account)
        if err := rows.Scan(&p.account_id,&p.name,&p.display_name,&p.note); err != nil {
            log.Printf("sql error")
        }
        g_t_account = append(g_t_account, p)
    }
    // print_t_account()
    // fmt.Println("------------------------------------")
    // str:="管理费用general and administrative expenses\\办公费Office Expense\\小件物品Small Object"
    // split(str)
    // fmt.Println(str)
    deal_with_split()
    print_t_account()
    fmt.Fprintln(w, "finish")
}
func deal_with_split() {
  for _,i:=range g_t_account{
        i.name=split(i.name)
        i.display_name=split(i.display_name)
        i.note=split(i.note)
    }
}
func split(strs string) string{
    str:=[]rune(strs)
    before:=0
    for i:=1;i<len(str);i++{
      // fmt.Println(str[before])
      // fmt.Println(str[i])
      if str[i]!=
      if str[before]<128&&str[i]>128{
        str = append(str[:i], append([]rune{35}, str[i:]...)...)
        //before=i
        i++
      }else if str[before]>128&&str[i]<128{
        str = append(str[:i], append([]rune{35}, str[i:]...)...)
        //before=i
        i++
      }
      before=i
    }
    strs=string(str)
    // fmt.Println(strs)
    return strs
}
func print_t_account() {
    for _,i:=range g_t_account{
        i.print()
    }
}
func checkErr(err error) {
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
}