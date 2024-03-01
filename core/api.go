package core

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lajosdeme/transaction-relayer/config"
	"github.com/lajosdeme/transaction-relayer/types"
	"github.com/lajosdeme/transaction-relayer/utils"

	"github.com/google/uuid"
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

	r.POST("/register", Register)
	r.POST("/login", Login)

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

func Register(c *gin.Context) {
	var payload *types.SignUpInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	var user types.User
	res := DB().Find(&user, "email = ?", strings.ToLower(payload.Email))
	if res.Error == nil && user.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "email already registered"})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "passwords do not match"})
		return
	}

	now := time.Now()
	newUser := types.User{
		ID:        uuid.New(),
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := DB().Create(newUser); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "user with email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	fmt.Println("new user registered: ", newUser.ID.String())

	// JWT generation
	token, err := utils.GenerateToken(config.Get().TokenExpiresIn, newUser.ID, config.Get().TokenSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	c.SetCookie("token", token, config.Get().TokenMaxAge*60, "/", config.Get().BaseUrl, false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"user_id": newUser.ID.String(),
		"token":   token,
	})
}

func Login(c *gin.Context) {
	fmt.Println("login request received")
	var payload types.SignInInput

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	var user types.User

	res := DB().Find(&user, "email = ?", strings.ToLower(payload.Email))
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "user doesn't exist"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		fmt.Println("password verify error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "invalid password"})
		return
	}

	// JWT generation
	token, err := utils.GenerateToken(config.Get().TokenExpiresIn, user.ID, config.Get().TokenSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	c.SetCookie("token", token, config.Get().TokenMaxAge*60, "/", config.Get().BaseUrl, false, true)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"token":  token,
	})
}
