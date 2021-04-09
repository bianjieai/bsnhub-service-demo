package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/common"
)

// HTTPService represents an HTTP service
type HTTPService struct {
	Router        *gin.Engine
	ChainManager  *ChainManager
}

// NewHTTPService constructs a new HTTPService instance
func NewHTTPService(
	chainManager *ChainManager,
) *HTTPService {
	srv := HTTPService{
		Router:        gin.Default(),
		ChainManager:  chainManager,
	}

	srv.createRouter()

	return &srv
}

func (srv *HTTPService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.Router.ServeHTTP(w, r)
}

func (srv *HTTPService) createRouter() {
	r := gin.Default()

	r.POST("/chains", srv.AddChain)
	r.GET("/chains", srv.GetChains)
	r.POST("/delete/:chainid", srv.DeleteChain)

	r.GET("/health", srv.ShowHealth)

	srv.Router = r
}

func (srv *HTTPService) AddChain(c *gin.Context) {
	var req AddChainRequest
	if err := c.BindJSON(&req); err != nil {
		onError(c, http.StatusBadRequest, "invalid JSON payload")
		return
	}
	chainID, err := srv.ChainManager.AddChain([]byte(req.ChainParams))
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
