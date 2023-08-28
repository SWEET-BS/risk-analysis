package pkg

import (
	"fmt"
	"riskanly/conf"
)

type TableInfo struct {
	Name   string
	Schema string
	CheckFields map[string]string
}
func NewTableInfo(name string, schema string) *TableInfo {
	return &TableInfo{Name: name, Schema: schema}
}
func (t *TableInfo) QueryAllTable() string {
	return fmt.Sprintf(conf.Queryalltable, t.Schema)
}
func (t *TableInfo) QueryTableColumns() string {
	return fmt.Sprintf(conf.Querytablecolumns, t.Schema, t.Name)
}
