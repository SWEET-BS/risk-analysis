package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// 待巡检 已检查

const (
	indexnamekey = "index_name"
	tablekey     = "source_table"
	wantkey      = "want"
	resultkey    = "result"
	filedkey     = "field"
	Dsnwtmp      = "user=postgres password=example host=192.168.200.58 port=5444 dbname=postgres sslmode=disable"
	//Dsnrtmp                = "user=postgres password=example host=192.168.200.58 port=5432 dbname=postgres sslmode=disable"
	maxIdleConnections = 10
	maxOpenConnections = 100
)

//QA 展示的 数据测试的结果  应当包括检测的表 字段 检测的规则名称 检测预期值 检测的结果
//task 任务id 任务名 处理逻辑 待质检列表
// case
// 企业表法人字段空置率大于90%
// QA 空置率，企业表， 法人， Result， 90%
// Task “select × from ” QAS

// QA 定义指标对象
// 一个任务对应多个质量（QA）结果
type QA struct {
	gorm.Model
	Qaname      string // 描述来自于任务
	IndexName   string // 完整度, xxx
	SourceTable string // table name
	Field       string // 字段名
	Want        string // 预期值
	Qaresult    string // 结果
}

// Task 1 Task ->qa1 , qa2 , qa3
type Task struct {
	gorm.Model
	Name   string
	Sql    string
	QAS    []*QA    `gorm:"_"`
	DSN    string   `gorm:"_"`
	Engine *gorm.DB `gorm:"_"`
}

// 整个生命周期按照顺序执行。不需要考虑性能问题

// core

// todo： achieve the functions

type TaskLife interface {
	// 数据库启动
	Start() error
	// 执行数据库，并且获取指标结果
	Run() ([]*QA, error)
	// 数据库停止
	Stop() error
}

func (t *Task) Start() error {
	var err error
	defer func() {
		var m any = nil
		if r := recover(); r != m {
			err = errors.New("panic occurred during database initialization")
		}
	}()
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(t.DSN), &gorm.Config{})
	if err != nil {
		return err

	}
	err = db.AutoMigrate(&QA{})
	if err != nil {
		fmt.Println("Failed to migrate table structures:", err)
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(maxIdleConnections)
	sqlDb.SetMaxOpenConns(maxOpenConnections)
	t.Engine = db

	return nil
}

func (t *Task) Run() ([]*QA, error) {
	results, err := executeSQLQuery(t.Engine, t.Sql)
	fmt.Println(results)
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		for _, row := range results {
			qa := &QA{
				Qaname:      t.Name,
				IndexName:   row[indexnamekey].(string),
				SourceTable: row[tablekey].(string),
				Field:       row[filedkey].(string),
				Want:        row[wantkey].(string),
				Qaresult:    fmt.Sprintf("%1.f%%", row[resultkey].(float64)),
			}
			t.QAS = append(t.QAS, qa)
			t.Engine.Create(qa)
		}
	}

	return t.QAS, nil
}
func (t *Task) Stop() error {
	sqlDB, err := t.Engine.DB()
	if err != nil {
		fmt.Println("Failed to get underlying *sql.DB:", err)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		fmt.Println("Failed to close database connection:", err)
		return err
	}
	return nil
}
func executeSQLQuery(db *gorm.DB, sql string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	// 使用Raw方法执行SQL查询
	result := db.Raw(sql).Scan(&results)
	if result.Error != nil {
		return nil, result.Error
	}
	return results, nil
}

// 每一个文字都有意义
/// 这里是为了调用者，描述的调用层的切片

//	type Report struct {
//		Task    []TaskLife
//		Content map[string]string
//	}
//
//	func (r *Report) report()  {
//		for k,v :=range r.Task{
//
//		}
//	}
var Tasktmp = Task{
	Name: "企业表法人字段非空值率小于90%",
	Sql:  `SELECT '90%' as want,'legal_rep' as field,'public.dm_lget_company_info' as source_table,'非空值率' as index_name,(COUNT(*) FILTER (WHERE legal_rep IS NOT NULL) * 100.0 / COUNT(*)) AS result FROM public.dm_lget_company_info;`,
	DSN:  Dsnwtmp,
}
func (t *Task) Jsontask() string{
	qas,_:=json.Marshal(t.QAS)
	message :=t.Name+string(qas)
	return message
}