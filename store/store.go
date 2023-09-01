package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"riskanly/conf"
)

// 待巡检 已检查

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
	Name          string
	Sql           string
	QAS           []*QA    `gorm:"_"`
	DSN           string   `gorm:"_"`
	Engine        *gorm.DB `gorm:"_"`
	IfCreateTable bool     `gorm:"_"`
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
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(conf.MaxIdleConnections)
	sqlDb.SetMaxOpenConns(conf.MaxOpenConnections)
	t.Engine = db
	if t.IfCreateTable == true {
		err = t.Engine.AutoMigrate(&QA{})
		if err != nil {
			fmt.Println("Failed to migrate table structures:", err)
			return err
		}
	}
	return nil
}

func (t *Task) Run() ([]*QA, error) {
	results, err := ExecuteSQLQuery(t.Engine, t.Sql)
	fmt.Println(results)
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		for _, row := range results {
			qa := &QA{
				Qaname:      t.Name,
				IndexName:   row[conf.Indexnamekey].(string),
				SourceTable: row[conf.Tablekey].(string),
				Field:       row[conf.Filedkey].(string),
				Want:        row[conf.Wantkey].(string),
				Qaresult:    fmt.Sprintf("%1.f%%", row[conf.Resultkey].(float64)),
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
func ExecuteSQLQuery(db *gorm.DB, sql string) ([]map[string]interface{}, error) {
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

func (t *Task) Jsontask() string {
	qas, _ := json.Marshal(t.QAS)
	message := t.Name + string(qas)
	return message
}

var Tasktmp = Task{
	Name: "企业表法人字段非空值率小于90%",
	Sql:  `SELECT '90%' as want,'legal_rep' as field,'public.dm_lget_company_info' as source_table,'非空值率' as index_name,(COUNT(*) FILTER (WHERE legal_rep IS NOT NULL) * 100.0 / COUNT(*)) AS result FROM public.dm_lget_company_info;`,
	DSN:  conf.Dsnwtmp,
}
// 查询某个条件下是否存在值
var Taskindex = Task{
	Name: "指标规则id存在数据",
	Sql: `SELECT t2.rule_id , COUNT(t1.*) AS count
FROM "index".inx_general t1
right JOIN "index".inx_regular_program t2 ON t1.rule_id  = t2.rule_id 
GROUP BY t2.rule_id  `,
	DSN: conf.DsnNewBarinSaas,
}
func (t *Task) CheckCount() string{
	var buf bytes.Buffer
	str,_ :=ExecuteSQLQuery(t.Engine,t.Sql)
	for _, m := range str {
		if m["count"].(int64)==int64(0){
			str1,_:=convertMapToString(m)
			buf.WriteString(str1+" ")
		}
	}
	return buf.String()
}

func convertMapToString(data map[string]interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
