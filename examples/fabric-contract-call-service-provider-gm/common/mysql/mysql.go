/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  mysql
 * @Version: 1.0.0
 * @Date: 2020/3/21 17:33
 */
package mysql

import (
	"bsn-irita-fabric-provider/common"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

//var MysqlDb *sql.DB
//var MysqlDbErr error
var Dbw DbWorker

const (
	Max_OpenConn     = 100
	Max_IdleConns    = 20
	Max_ConnLifeTime = 100 * time.Second
)

type DbWorker struct {
	//mysql data source name
	Dsn string
}

func Init(connString string) {
	log.Println("设置数据库连接：" + connString)
	Dbw = DbWorker{
		Dsn: connString,
	}
}

func OpenDb() (*sql.DB, error) {

	MysqlDb, MysqlDbErr := sql.Open("mysql", Dbw.Dsn)
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + Dbw.Dsn)
		return nil, MysqlDbErr
		//panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(Max_OpenConn)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(Max_IdleConns)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(Max_ConnLifeTime)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		//panic("数据库链接失败: " + MysqlDbErr.Error())
	}

	return MysqlDb, MysqlDbErr
}

func Exec(sql string, Args ...interface{}) (lastId, rows int64, err error) {

	MysqlDb, err := OpenDb()
	defer MysqlDb.Close()
	if err != nil {
		return 0, 0, err
	}
	ret, err := MysqlDb.Exec(sql, Args...)

	if err != nil {
		return 0, 0, err
	}

	//插入数据的主键id
	lastId, _ = ret.LastInsertId()
	//fmt.Println("LastInsertID:",lastInsertID)

	//影响行数
	rows, _ = ret.RowsAffected()
	//fmt.Println("RowsAffected:",rowsaffected)

	return
}

type QueryOption func(rows *sql.Rows) (interface{}, error)

func Query(option QueryOption, sql string, Args ...interface{}) (list []interface{}, err error) {

	MysqlDb, err := OpenDb()
	defer MysqlDb.Close()
	if err != nil {
		return list, err
	}

	rows, err := MysqlDb.Query(sql, Args...)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		return list, err
	}

	for rows.Next() {

		data, err := option(rows)
		if err != nil {
			return list, err
		}

		list = append(list, data)
	}

	return list, nil

}

func TableIsCreated(tableName string) bool {
	sqlstr := fmt.Sprintf("SELECT table_name FROM information_schema.TABLES WHERE table_name = '%s' ", tableName)

	exit := func(rows *sql.Rows) (interface{}, error) {
		var table_name string
		rows.Scan(&table_name)

		return table_name, nil
	}

	list, err := Query(exit, sqlstr)

	if err != nil {
		fmt.Println(err)
	}

	if len(list) > 0 {

		fmt.Println(list[0])
		//return list[0].(int64)>0,nil
		return list[0] == tableName
	} else {
		return false
	}
}

func TabIsExit(tableName string) bool {
	sqlstr := fmt.Sprintf("SELECT table_name FROM information_schema.TABLES WHERE table_name = '%s' ", tableName)

	exit := func(rows *sql.Rows) (interface{}, error) {
		var table_name string
		rows.Scan(&table_name)

		return table_name, nil
	}

	list, err := Query(exit, sqlstr)

	if err != nil {
		fmt.Println(err)
	}

	if len(list) > 0 {

		fmt.Println(list[0])
		//return list[0].(int64)>0,nil
		return list[0] == tableName
	} else {
		return false
	}
}

func CreateTable(sql string, tabName string) {

	id, rows, err := Exec(sql)
	if err != nil {
		common.Logger.Error(err)
	}
	common.Logger.Info("数据库创建:", tabName)
	common.Logger.Info(id, rows)

}
