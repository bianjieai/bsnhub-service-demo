package mysql

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type table struct {
	funique_id    int64
	request_id    string     //请求唯一id
	from_chainid  string     //起始链ID
	from_tx       string     //起始链交易ID
	hub_req_tx    string     //HUB请求交易ID
	ic_request_id string     //HUB请求交易ID
	to_chainid    string     //目标链ID
	to_tx         *string    //目标链交易ID
	hub_res_tx    string     //HUB响应交易ID
	from_res_tx   *string    //向起始链响应数据的交易ID
	tx_status     int        //交易状态
	tx_time       *time.Time //交易完成时间
	tx_createtime time.Time  //交易创建时间
	error         string     //交易异常
}

// NewDB create a instance of the mysql db
func NewDB(mysqlConfig Config, options ...Option) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		mysqlConfig.DBUserName, mysqlConfig.DBUserPassphrase, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	var cfg = config{
		maxIdleConns: 10,
		maxOpenConns: 10,
		maxLifetime:  time.Hour,
		debug:        false,
	}
	for _, optionFn := range options {
		if err := optionFn(&cfg); err != nil {
			panic(err)
		}
	}

	//SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(cfg.maxIdleConns)

	//SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(cfg.maxOpenConns)

	//SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(cfg.maxLifetime)

	DB = db
}

func insert(field string, value string) error {

	stmt, err := DB.Prepare("INSERT tb_irita_crosschain_tx SET " + field + "=?" + ", source_service = 1 ")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(value)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println("insert succ:", id)

	return nil
}

func update(field, value, icResID string) error {
	cmd := fmt.Sprintf("UPDATE tb_irita_crosschain_tx SET %s=? where ic_request_id='%s' and source_service = 1", field, icResID)
	stmt, err := DB.Prepare(cmd)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(value)
	if err != nil {
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("update succ:", row)

	return nil
}
func updateTime(field, icResID string) error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	cmd := fmt.Sprintf("UPDATE tb_irita_crosschain_tx SET %s='%s' where ic_request_id='%s' and source_service = 1 ", field, nowTime, icResID)
	stmt, err := DB.Prepare(cmd)
	if err != nil {
		return err
	}

	res, err := stmt.Exec()
	if err != nil {
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("update succ:", row)

	return nil
}

func query(field, value string) error {
	var table table
	sqlStr := "select funique_id,request_id,from_chainid,from_tx,hub_req_tx,to_chainid,to_tx,hub_res_tx,from_res_tx,tx_status,tx_time,tx_createtime from tb_irita_crosschain_tx where " + field + "=?"
	err := DB.QueryRow(sqlStr, value).Scan(&table.funique_id, &table.request_id, &table.from_chainid, &table.from_tx, &table.hub_req_tx, &table.to_chainid, &table.to_tx, &table.hub_res_tx, &table.from_res_tx, &table.tx_status, &table.tx_time, &table.tx_createtime)
	if err != nil {
		return err
	}
	fmt.Printf("table:%v\n", table)
	return nil
}

func Close() {
	_ = DB.Close()
}
