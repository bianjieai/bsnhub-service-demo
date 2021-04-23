package server

import (
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/common"
	"github.com/gin-gonic/gin"
	"github.com/irisnet/service-sdk-go/service"
	"io/ioutil"
	"net/http"
)

var (
	callBack service.RespondCallback
)

func SetTestCallBack(cb service.RespondCallback) {
	callBack = cb
}

// HTTPService represents an HTTP service
type HTTPService struct {
	Router       *gin.Engine
	ChainManager *ChainManager
}

// NewHTTPService constructs a new HTTPService instance
func NewHTTPService(
	chainManager *ChainManager,
) *HTTPService {
	srv := HTTPService{
		Router:       gin.Default(),
		ChainManager: chainManager,
	}

	srv.createRouter()

	return &srv
}

func (srv *HTTPService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.Router.ServeHTTP(w, r)
}

func (srv *HTTPService) createRouter() {
	r := gin.Default()
	api := r.Group("/api/v0")
	fiscobcos := api.Group("/fiscobcos")
	{
		fiscobcos.POST("/chains", srv.AddChain)
		fiscobcos.GET("/chains", srv.GetChains)
		fiscobcos.POST("/delete/:chainid", srv.DeleteChain)
	}

	r.GET("/health", srv.ShowHealth)
	r.POST("/test", srv.TestCallBack)

	srv.Router = r
}

func (srv *HTTPService) TestCallBack(c *gin.Context) {

	common.Logger.Infof("Into TestCallBack")

	var bodyBytes []byte

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		common.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, "invalid JSON payload")
		return
	}
	common.Logger.Infof("Test Data Is : %s", string(bodyBytes))

	output, result := callBack("testreqCtxId", "testreqId", string(bodyBytes))

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

	chainID, err := srv.ChainManager.AddChain(bodyBytes)
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, AddChainResult{ChainID: chainID})
}

func (srv *HTTPService) DeleteChain(c *gin.Context) {
	chainID := c.Param("chainid")
	if err := srv.ChainManager.DeleteChain(chainID); err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}

	onSuccess(c, nil)
}

func (srv *HTTPService) GetChains(c *gin.Context) {
	chains, err := srv.ChainManager.GetChains()
	if err != nil {
		onError(c, http.StatusInternalServerError, err.Error())
		return
	}
	onSuccess(c, chains)
}

// ShowHealth returns the health state
func (srv *HTTPService) ShowHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": true})
}

func onError(c *gin.Context, code int, msg string) {
	common.Logger.Errorf(msg)

	c.JSON(code, ErrorResponse{
		Code:  CODE_ERROR,
		Error: msg,
	})
}

func onSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code:   CODE_SUCCESS,
		Result: result,
	})
}
