package pkg

import (
	"fmt"
	"gorm.io/gorm"
	"riskanly/conf"
	"riskanly/store"
	"strings"
)

type TableInfo struct {
	Name        string
	Schema      string
	CheckFields map[string]string
	IfCheck     bool
}

func NewTableInfo(name string, schema string) *TableInfo {
	return &TableInfo{Name: name, Schema: schema}
}
func (t *TableInfo) QueryAllTableSql() string {
	return fmt.Sprintf(conf.Queryalltable, t.Schema)
}
func (t *TableInfo) QueryTableColumnsSql() string {
	var keys []string
	var andin string
	if t.IfCheck && t.CheckFields != nil {
		for k := range t.CheckFields {
			keys = append(keys, k)
		}
		andin = fmt.Sprintf("and column_name in ('%s')", strings.Join(keys, "','"))
	}
	return fmt.Sprintf(conf.Querytablecolumns, t.Schema, t.Name, andin)
}
func (t *TableInfo) QueryAllTable(db *gorm.DB) ([]*TableInfo, error) {
	alltbales := make([]*TableInfo, 0)
	alltable, err := store.ExecuteSQLQuery(db, t.QueryAllTableSql())
	for _, v := range alltable {
		tablename := &TableInfo{
			Name: v[conf.Keytable].(string),
		}
		alltbales = append(alltbales, tablename)
	}
	if err != nil {
		return nil, err
	}
	return alltbales, nil
}
func (t *TableInfo) QueryAllColumns(db *gorm.DB) {
	//fmt.Println(t.QueryTableColumnsSql())
	//待检测字段已经存在t中
	columns, err := store.ExecuteSQLQuery(db, t.QueryTableColumnsSql())
	if err != nil {
		return
	}
	//fmt.Println(columns)
	for _,v:=range columns{
		if _,ok:=t.CheckFields[v[conf.KeyColumn].(string)];ok{
			t.CheckFields[v[conf.KeyColumn].(string)]="存在"
		}
	}
}
