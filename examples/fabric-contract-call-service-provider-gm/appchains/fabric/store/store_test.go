package store

import (
	"bsn-irita-fabric-provider/appchains/fabric/entity"
	"fmt"
	"testing"
)

func TestInitTable(t *testing.T) {

	conn := "root:123456@tcp(192.168.1.61:3306)/bsnflowdb?charset=utf8"
	InitMysql(conn)

	data := &entity.FabricChainInfo{
		ChainBase: entity.ChainBase{ChainId: 101},
		ChannelId: "1000",
		AppCode:   "app001",
		NodeName:  "peer1;peer2",
		CityCode:  "ORG1110",
	}

	err := StoreProviderAppInfo(data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetProviderAppInfo(t *testing.T) {
	conn := "root:123456@tcp(192.168.1.61:3306)/bsnflowdb?charset=utf8"
	InitMysql(conn)

	err := GetProviderAppInfo("100")
	if err != nil {
		fmt.Print(err)
	}
}

func TestTargetChainInfo(t *testing.T) {
	conn := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"
	InitMysql(conn)

	data := &entity.CrossChainInfo{
		Request_id:     "fds",
		To_chainid:     "1111",
		To_tx:          "das",
		Source_service: 1,
	}

	TargetChainInfo(data)

}
func TestRequestHubInfo(t *testing.T) {
	conn := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"
	InitMysql(conn)

	data := &entity.CrossChainInfo{
		Tx_status:      2,
		Error:          "xcxxc",
		Request_id:     "fds",
		Source_service: 1,
	}
	RequestHubInfoFail(data)
}

func TestRequestHubInfoSucc(t *testing.T) {
	conn := "root:123456@tcp(192.168.1.60:3306)/bsnflowdb?charset=utf8"
	InitMysql(conn)

	data := &entity.CrossChainInfo{
		Tx_status:      1,
		Hub_res_tx:     "swsdadas",
		Request_id:     "fds",
		Source_service: 1,
	}
	RequestHubInfoSucc(data)
}
