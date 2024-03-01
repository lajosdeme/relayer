package core

import (
	"errors"
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

	r.POST("/verify/address", VerifyAddress)

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

	u, err := DB().GetUser(payload.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	if !utils.Contains(u.VerifiedAddresses, payload.Address) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "request to unverified address"})
		return
	}

	currentTime := time.Now().Unix()
	if u.Quota.ResetDate < int(currentTime) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "user is not subscribed"})
		return
	}

	hash, gasUsed, err := ExecuteRelayCall(payload, u.Quota.Quota)

	if err != nil {
		fmt.Println("failed to execute relay call: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	if err := DB().UpdateQuotaForUser(payload.UserId, gasUsed); err != nil {
		fmt.Println("error updating quota: ", err)
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
	resp, err := authenticateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}
	u, err := DB().GetUser(resp.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u.Quota)
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

func VerifyAddress(c *gin.Context) {
	resp, err := authenticateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	u, err := DB().GetUser(resp.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	var payload types.VerifiedAddressInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Println("failed to bind payload: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	u.VerifiedAddresses = append(u.VerifiedAddresses, payload.Address)
	if err := DB().UpdateVerifiedAddresses(u.ID, u.VerifiedAddresses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func authenticateRequest(c *gin.Context) (types.AuthResponse, error) {
	token, err := getToken(c)
	if err != nil {
		return types.AuthResponse{Status: "fail"}, err
	}

	sub, err := utils.ValidateToken(token, config.Get().TokenSecret)
	if err != nil {
		return types.AuthResponse{Status: "fail"}, err
	}

	var user types.User
	result := DB().Find(&user, "id = ?", fmt.Sprint(sub))

	if result.Error != nil {
		return types.AuthResponse{Status: "fail"}, errors.New("user doesn't exist")
	}

	return types.AuthResponse{Status: "ok", UserId: user.ID.String()}, nil
}

func getToken(c *gin.Context) (string, error) {
	var token string
	cookie, err := c.Cookie("token")

	authorizationHeader := c.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		token = fields[1]
	} else if err == nil {
		token = cookie
	}

	if token == "" {
		return "", errors.New("empty token")
	}

	return token, nil
}
