package entity

import (
	"strconv"
	"strings"
	"time"
)

type ChainBase struct {
	ChainId uint64 `json:"chainId"`
}

func (r *ChainBase) SetChainId(cid string) error {

	ccid, err := strconv.ParseUint(cid, 10, 64)
	if err != nil {
		return err
	}
	r.ChainId = ccid
	return nil
}

func (r *ChainBase) GetChainId() string {
	return strconv.FormatUint(r.ChainId, 10)
}

func (r *ChainBase) GetChainIdKey() string {
	return "chainId_" + strconv.FormatUint(r.ChainId, 10)
}

type FabricChainInfo struct {
	ChainBase

	//ID database uniqueId
	ID uint64

	AppCode string

	ChannelId string

	NodeName string

	CityCode string

	Status int

	CreateTime time.Time
}

func (f *FabricChainInfo) SetNodes(nodes []string) {
	f.NodeName = strings.Join(nodes, ";")
}

func (f *FabricChainInfo) GetNodes() []string {
	return strings.Split(f.NodeName, ";")
}

type CrossChainInfo struct {
	Funique_id     uint64
	Request_id     string
	From_chainid   string
	From_tx        string
	Hub_req_tx     string
	Ic_request_id  string
	To_chainid     string
	To_tx          string
	Hub_res_tx     string
	From_res_tx    string
	Tx_status      int
	Tx_time        time.Time
	Tx_createtime  time.Time
	Error          string
	Source_service int
}
