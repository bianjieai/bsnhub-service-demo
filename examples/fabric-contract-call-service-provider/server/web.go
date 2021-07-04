package server

import (
	"bsn-irita-fabric-provider/appchains"
	"bsn-irita-fabric-provider/common"
	"github.com/gin-gonic/gin"

	"io/ioutil"
	"net/http"
)

type HTTPService struct {
	Router   *gin.Engine
	AppChain appchains.AppChainHandlerI
}

func NewHTTPService(
	appChain appchains.AppChainHandlerI,
) *HTTPService {
	srv := HTTPService{
		Router:   gin.Default(),
		AppChain: appChain,
	}

	srv.createRouter()

	return &srv
}

func (srv *HTTPService) createRouter() {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	api := r.Group("/api/v0")
	fabric := api.Group("/fabric")
	{
		fabric.POST("/registerChain", srv.AddChain)
		fabric.POST("/removeAppChain/:chainId", srv.DeleteChain)
		fabric.POST("/updateAppChain", srv.UpdateChain)
	}

	r.GET("/health", srv.ShowHealth)
	r.POST("/test", srv.TestSP)

	srv.Router = r
}

func (srv *HTTPService) TestSP(c *gin.Context) {
	common.Logger.Infoln("Into testSP")

	var bodyBytes []byte

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		common.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}
	common.Logger.Infof("bodyBytes：%s", string(bodyBytes))

	output, result := srv.AppChain.Callback("", "", string(bodyBytes))

	common.Logger.Infof("output:%s", output)
	common.Logger.Infof("result:%s", result)

	onSuccess(c, output)
}

func (srv *HTTPService) AddChain(c *gin.Context) {
	var bodyBytes []byte

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		common.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}

	common.Logger.Infof("AddChain data is %s", string(bodyBytes))

	err = srv.AppChain.RegisterChain(bodyBytes)
	if err != nil {
		common.Logger.Errorf(err.Error())
		onError(c, err)
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) DeleteChain(c *gin.Context) {
	chainID := c.Param("chainId")

	err := srv.AppChain.DeleteChain(chainID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		onError(c, err)
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) UpdateChain(c *gin.Context) {
	var bodyBytes []byte

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		common.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}

	common.Logger.Infof("AddChain data is %s", string(bodyBytes))

	err = srv.AppChain.UpdateChain(bodyBytes)
	if err != nil {
		common.Logger.Errorf(err.Error())
		onError(c, err)
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) ShowHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": true})
}
func onError(c *gin.Context, err error) {
	common.Logger.Error(err.Error())

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:  CODE_ERROR,
		Error: err.Error(),
	})
}

func onSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code:   CODE_SUCCESS,
		Msg:    "success",
		Result: result,
	})
}
