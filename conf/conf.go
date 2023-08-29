package conf

// pkg.check package
const (
	KeyColumn     = "column_name"
	Keytable      = "table_name"
	Queryalltable = `
					 SELECT table_name
                     FROM information_schema.tables
                     WHERE table_schema = '%s'
                     AND table_type = 'BASE TABLE';
                     `
	Querytablecolumns = `
					SELECT column_name
					FROM information_schema.columns
					WHERE table_schema = '%s'
					AND table_name = '%s'  %s
					`
)

// store.store  package
const (
	Indexnamekey       = "index_name"
	Tablekey           = "source_table"
	Wantkey            = "want"
	Resultkey          = "result"
	Filedkey           = "field"
	Dsnwtmp            = "不做上传"
	MaxIdleConnections = 10
	MaxOpenConnections = 100
)

// pkg.bot package
const (
	Sceret     = "数据巡检："
	WebhookURL = "不做上传"
)
