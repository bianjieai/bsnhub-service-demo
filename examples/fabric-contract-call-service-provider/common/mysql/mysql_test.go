/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  mysql_test
 * @Version: 1.0.0
 * @Date: 2020/3/21 17:39
 */

package mysql

import (
	"fmt"
	"testing"
	"time"
)

const (
	USER_NAME = "root"
	PASS_WORD = "123456"
	HOST      = "192.168.1.61"
	PORT      = "3306"
	DATABASE  = "bsnflowdb"
	CHARSET   = "utf8"
)

func TestInsert2(t *testing.T) {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	fmt.Println(dbDSN)
	Init(dbDSN)

	sql := "insert INTO tb_test(name,address,allTime) values(?,?,?)"

	lastId, rows, err := Exec(sql, "小红", 233, time.Now())
	if err != nil {
		t.Fatal("Err:", err.Error())
	} else {
		fmt.Println(lastId, rows)
	}

}
