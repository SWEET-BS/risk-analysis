package pkg

import (
	"fmt"
	"log"
	"riskanly/conf"
	"riskanly/store"
	"testing"
)

func TestA(t *testing.T) {
	tmp :=store.Task{
		DSN: conf.DsnNewBarinSaas,
	}
	table_info :=TableInfo{
		Schema: "ads",
		IfCheck: true,
	}
	tmp.Start()
	//tmp.Sql=table_info.QueryAllTableSql()
	//query,_:=store.ExecuteSQLQuery(tmp.Engine,tmp.Sql)
	allnames ,err:=table_info.QueryAllTable(tmp.Engine)
	if err!=nil{
		log.Panic(err)
	}
	for i := 0; i < len(allnames); i++ {

		checktable := &TableInfo{
			Schema: table_info.Schema,
			Name:  allnames[i].Name,
			CheckFields: map[string]string{
				"id":          "不存在",
				"update_time": "不存在",
				"is_delete":   "不存在",
				"create_time": "不存在",
			},
			IfCheck: table_info.IfCheck,
		}
		checktable.QueryAllColumns(tmp.Engine)
		fmt.Println(checktable)
	}

}
