package fabric

import (
	"bsn-irita-fabric-provider/appchains/fabric/config"
	"bsn-irita-fabric-provider/appchains/fabric/entity"
	logging "bsn-irita-fabric-provider/common"
	"bsn-irita-fabric-provider/errors"
	"fmt"
	"github.com/BSNDA/fabric-sdk-go-gm/pkg/client/channel"
	"github.com/BSNDA/fabric-sdk-go-gm/pkg/fabsdk"
	"time"
)

func NewFabricChain(sdkConf *config.FabricConfig, info *entity.FabricChainInfo) (*FabricChain, error) {
	fabric := &FabricChain{
		chainInfo: info,
		config:    sdkConf,
	}
	sdk, err := fabric.fabSdk()
	if err != nil {
		logging.Logger.Errorf("fabric sdk init failed %s", err)
		return nil, errors.New("fabric sdk init failed %s", err)

	}
	fabric.sdk = sdk

	channelProvider := fabric.sdk.ChannelContext(fabric.chainInfo.ChannelId,
		fabsdk.WithOrg(fabric.config.OrgName),
		fabsdk.WithUser(fabric.config.MspUserName),
	)

	client, err := channel.New(channelProvider)
	if err != nil {
		logging.Logger.Errorf("fabric channel client init failed %s", err)
		return nil, errors.New("fabric channel client init failed %s", err)
	}
	fabric.channelClient = client
	fabric.LastCallTime = time.Now()

	return fabric, nil

}

type FabricChain struct {
	sdk           *fabsdk.FabricSDK
	channelClient *channel.Client

	config    *config.FabricConfig
	chainInfo *entity.FabricChainInfo

	LastCallTime time.Time
}

func (f *FabricChain) Close() {
	if f.sdk != nil {
		f.sdk.Close()
	}

}

func (f *FabricChain) fabSdk() (*fabsdk.FabricSDK, error) {
	//todo dynamic generation fabric SDK

	conf := f.config.GetSdkConfig(f.chainInfo.ChannelId, f.chainInfo.GetNodes())
	c, err := conf()
	if err != nil {
		ps, _ := c[0].Lookup("channels")
		logging.Logger.Infof("New Fabric SDK Channels is  %s", ps)
	}

	sdk, err := fabsdk.New(conf)

	if err != nil {
		logging.Logger.Errorf("New Fabric SDK has error %s", err.Error())
	}

	return sdk, err
}

func (f *FabricChain) Invoke(chaincode string, args []string) (*entity.FabricRespone, error) {
	fabRes, err := f.channelClient.Execute(f.getChannelRequest(chaincode, args))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("call fabric Invoke failed :%s", err))
	}

	return entity.NewFabricRes(fabRes), nil
}

func (f *FabricChain) Query(chaincode string, args []string) (*entity.FabricRespone, error) {

	fabRes, err := f.channelClient.Query(f.getChannelRequest(chaincode, args))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("call fabric Query failed :%s", err))
	}

	return entity.NewFabricRes(fabRes), nil
}

func (f *FabricChain) getChannelRequest(chaincode string, args []string) channel.Request {

	f.LastCallTime = time.Now()

	fcn := args[0]

	var byteArgs [][]byte

	for i := 1; i < len(args); i++ {
		byteArgs = append(byteArgs, []byte(args[i]))
	}
	request := channel.Request{
		ChaincodeID: chaincode,
		Fcn:         fcn,
		Args:        byteArgs,
	}

	return request
}
