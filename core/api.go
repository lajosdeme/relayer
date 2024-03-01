package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lajosdeme/transaction-relayer/config"
	"github.com/lajosdeme/transaction-relayer/types"
)

func RunRouter() {
	r := configRouter()
	r.Run(fmt.Sprintf(":%d", config.Get().Port))
}

func configRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", ping)
	r.POST("/execute", execute)
	r.POST("/quota", quota)

	return r
}

func execute(c *gin.Context) {
	fmt.Println("new request")
	var payload types.ExecutePayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Println("failed to bind payload: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	hash, err := ExecuteRelayCall(payload)

	if err != nil {
		fmt.Println("failed to execute relay call: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	resp := types.ExecuteResponse{
		TransactionHash: hash,
	}

	c.JSON(http.StatusOK, resp)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// Dummy quota, TODO: Implement real quota calculation and getter logic
func quota(c *gin.Context) {
	q := types.Quota{
		Quota:      20000000,
		Unit:       "gas",
		TotalQuota: 20000000,
		ResetDate:  1764098470,
	}
	c.JSON(http.StatusOK, q)
}
