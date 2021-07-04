package store

import (
	"bsn-irita-fabric-provider/appchains/fabric/entity"
	"bsn-irita-fabric-provider/common"
	"bsn-irita-fabric-provider/common/mysql"
	"database/sql"
	"fmt"
	"time"
)

//存储 fabric provider 信息
func StoreProviderAppInfo(data *entity.FabricChainInfo) error {

	common.Logger.Infof("Store Provider AppInfo data to mysql,data:%v", data)

	sql := fmt.Sprintf("INSERT INTO %s (chainId,appCode,channelId,nodes,cityCode,status,createTime) values (?,?,?,?,?,?,?)", _TabName_Provider)

	data.Status = 1
	data.CreateTime = time.Now()
	creatTime := data.CreateTime.Format("2006-01-02 15:04:05")
	lastID, rows, err := mysql.Exec(
		sql,
		data.GetChainId(),
		data.AppCode,
		data.ChannelId,
		data.NodeName,
		data.CityCode,
		data.Status,
		creatTime)

	if err != nil {
		common.Logger.Errorf("save ChainId：%s err!：%s", data.GetChainId(), err.Error())
		return err
	}
	common.Logger.Infof("save success ChainId：%s", data.GetChainId())
	common.Logger.Infof("save lastID：%d, Number of rows affected：%d", lastID, rows)

	return err

}

func UpdateProviderAppInfo(data *entity.FabricChainInfo) error {

	common.Logger.Infof("Store Provider AppInfo data to mysql,data:%v", data)

	sql := fmt.Sprintf("update %s set channelId = ?,nodes =? where chainId = ?", _TabName_Provider)

	lastID, rows, err := mysql.Exec(
		sql,
		data.ChannelId,
		data.NodeName,
		data.GetChainId(),
	)

	if err != nil {
		common.Logger.Errorf("save ChainId：%s err!：%s", data.GetChainId(), err.Error())
		return err
	}
	common.Logger.Infof("save success ChainId：%s", data.GetChainId())
	common.Logger.Infof("save lastID：%d, Number of rows affected：%d", lastID, rows)

	return err
}
func DeleteProviderAppInfo(chainId string) error {
	common.Logger.Infof("Into DeleteProviderAppInfo method：")

	Deletesql := fmt.Sprintf("DELETE FROM %s WHERE ChainId = %s;", _TabName_Provider, chainId)

	lastId, rows, err := mysql.Exec(Deletesql)
	if err != nil {
		common.Logger.Errorf("DELETE err: %s", err.Error())
		return err
	}

	common.Logger.Infof("DELETE success ChainId：%s\n", chainId)
	common.Logger.Infof("DELETE lastID：%d, Number of rows affected：%d\n", lastId, rows)

	return err
}

func GetProviderAppInfo(chainId string) *entity.FabricChainInfo {

	common.Logger.Infof("Into GetProviderAppInfo method：")
	querysql := fmt.Sprintf("SELECT id,ChainId,AppCode,ChannelId,Nodes,CityCode,STATUS, CreateTime FROM %s WHERE chainId=%s;", _TabName_Provider, chainId)
	query := func(rows *sql.Rows) (interface{}, error) {

		var ChainId string
		var createTime string

		data := &entity.FabricChainInfo{}
		rows.Scan(
			&data.ID,
			&ChainId,
			&data.AppCode,
			&data.ChannelId,
			&data.NodeName,
			&data.CityCode,
			&data.Status,
			&createTime)

		DefaultTimeLoc := time.Local
		c, _ := time.ParseInLocation("2006-01-02 15:04:05", createTime, DefaultTimeLoc)
		data.CreateTime = c

		data.SetChainId(ChainId)

		return data, nil
	}

	common.Logger.Infof("start query ChainIdAppInfoList：")

	ChainInfoList, err := mysql.Query(query, querysql)
	if err != nil {
		common.Logger.Errorf("query err：%s\n", err.Error())
	}

	if len(ChainInfoList) == 0 {
		//fmt.Print(err)
		common.Logger.Infof("query ChainIdAppInfoList is 0")
		return nil
	}

	return ChainInfoList[0].(*entity.FabricChainInfo)
}

func TargetChainInfo(data *entity.CrossChainInfo) {
	common.Logger.Infof("Store Provider HubInfo data to mysql,data:%#v", data)
	insertsql := fmt.Sprintf("insert into  %s (ic_request_id,to_chainid,to_tx,source_service,tx_createtime,tx_status,error) values (?,?,?,?,?,?,?)", _TabName_cc_Tx)
	lastID, rows, err := mysql.Exec(
		insertsql,
		data.Ic_request_id,
		data.To_chainid,
		data.To_tx,
		data.Source_service,
		data.Tx_createtime.Format("2006-01-02 15:04:05"),
		data.Tx_status,
		data.Error,
	)

	if err != nil {
		common.Logger.Errorf("save Ic_request_id：%s err!：%s", data.Ic_request_id, err.Error())
	}

	common.Logger.Infof("save TargetChainInfo Ic_request_id：%s", data.Ic_request_id)
	common.Logger.Infof("save lastID：%d, Number of rows affected：%d", lastID, rows)

}

func RequestHubInfoFail(data *entity.CrossChainInfo) {
	common.Logger.Infof("Store Provider HubInfoFail data to mysql,data:%#v", data)
	updatesql := fmt.Sprintf("update %s set tx_status =?,error =? where ic_request_id=? and source_service=?;", _TabName_cc_Tx)
	lastID, rows, err := mysql.Exec(
		updatesql,
		data.Tx_status,
		data.Error,
		data.Ic_request_id,
		data.Source_service,
	)

	if err != nil {
		common.Logger.Errorf("save Ic_request_id：%s err!：%s", data.Ic_request_id, err.Error())
	}

	common.Logger.Infof("save RequestHubInfoFail Ic_request_id：%s", data.Ic_request_id)
	common.Logger.Infof("save lastID：%d, Number of rows affected：%d", lastID, rows)

}

func RequestHubInfoSucc(data *entity.CrossChainInfo) {
	common.Logger.Infof("Store Provider HubInfoSuccess data to mysql,data:%#v", data)
	updatesql := fmt.Sprintf("update %s set tx_status =?,hub_res_tx=? where ic_request_id=? and source_service=?;", _TabName_cc_Tx)
	lastID, rows, err := mysql.Exec(
		updatesql,
		data.Tx_status,
		data.Hub_res_tx,
		data.Ic_request_id,
		data.Source_service,
	)

	if err != nil {
		common.Logger.Errorf("save Ic_request_id：%s err!：%s", data.Ic_request_id, err.Error())
	}

	common.Logger.Infof("save RequestHubInfoSucc Ic_request_id：%s", data.Ic_request_id)
	common.Logger.Infof("save lastID：%d, Number of rows affected：%d", lastID, rows)
}
