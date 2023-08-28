package conf

//pkg.check package
const (
		Queryalltable = `
					 SELECT table_name
                     FROM information_schema.tables
                     WHERE table_schema = '%s'
                     AND table_type = 'BASE TABLE';
                     `
        Querytablecolumns=`
					SELECT column_name, data_type
					FROM information_schema.columns
					WHERE table_schema = '%s'
					AND table_name = '%s';
					`
)
//store.store  package
const (
	Indexnamekey = "index_name"
	Tablekey     = "source_table"
	Wantkey      = "want"
	Resultkey    = "result"
	Filedkey     = "field"
	//Dsnwtmp      = "user=postgres password=example host=192.168.200.58 port=5444 dbname=postgres sslmode=disable"
	//DsnDm                = "user=postgres password=example host=192.168.200.58 port=5432 dbname=postgres sslmode=disable"
	Dsnwtmp ="不做上传"
	MaxIdleConnections = 10
	MaxOpenConnections = 100
)
//pkg.bot package
const (
	Sceret     = "数据巡检："
	//WebhookURL = "https://oapi.dingtalk.com/robot/send?access_token=2282b1d0eeded882853e3b82f358f35573a2f411f04e9fefb8a7818bde574db2"
	//WebhookURL="https://oapi.dingtalk.com/robot/send?access_token=e6bd5b1004a9deae492ce606213d8cca400ab49cfbab4689fd21fb934bf2f94d"
	WebhookURL ="不做上传"
)

