package entity

import (
	"time"
)

type RegisterData struct {
	ChainId   uint64   `json:"chainId"`
	ChannelId string   `json:"channelId"`
	Nodes     []string `json:"nodes"`
	AppCode   string   `json:"appCode"`
}
type DeleteChain struct {
	ChainId uint64 `json:"chainId"`
}

func (r *RegisterData) ToStoreDate(cityCode string) *FabricChainInfo {
	provider := &FabricChainInfo{
		ChainBase:  ChainBase{ChainId: r.ChainId},
		AppCode:    r.AppCode,
		ChannelId:  r.ChannelId,
		CityCode:   cityCode,
		Status:     0,
		CreateTime: time.Now(),
	}
	provider.SetNodes(r.Nodes)
	return provider
}
func (r *RegisterData) UpdateStoreDate() *FabricChainInfo {
	data := &FabricChainInfo{}
	data.ChainId = r.ChainId
	data.AppCode = r.AppCode
	data.ChannelId = r.ChannelId
	data.SetNodes(r.Nodes)
	return data
}
