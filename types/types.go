package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Config struct {
	NodeUrl string `mapstructure:"NODE_URL"`
	Key     string `mapstructure:"KEY"`
	Mode    string `mapstructure:"MODE"`
	Port    int    `mapstructure:"PORT"`

	BaseUrl                     string `mapstructure:"BASE_URL"`
	SubscriptionContractAddress string `mapstructure:"SUBSCRIPTION_CONTRACT_ADDRESS"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRES_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}

type User struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`

	Password string `gorm:"not null"`

	VerifiedAddresses VerifiedAddresses `json:"verified_addresses"`

	Quota     Quota     `json:"quota"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type ExecutePayload struct {
	UserId  string      `json:"user_id"`
	Address string      `json:"address"`
	Tx      Transaction `json:"transaction"`
}

type Transaction struct {
	Abi            string `json:"abi"`
	Nonce          string `json:"nonce"`
	Signature      string `json:"signature"`
	ValidityTstamp string `json:"validityTimestamps"`
}

type ExecuteResponse struct {
	TransactionHash string `json:"transactionHash"`
}

type Quota struct {
	Quota      int    `json:"quota"`
	Unit       string `json:"unit"`
	TotalQuota int    `json:"totalQuota"`
	ResetDate  int    `json:"resetDate"`
}

type AuthResponse struct {
	Status string `json:"status"`
	UserId string `json:"user_id"`
}

type VerifiedAddressInput struct {
	Address string `json:"address"`
}

type VerifiedAddresses []string

func (a VerifiedAddresses) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "[]", nil
	}
	return fmt.Sprintf(`["%s"]`, strings.Join(a, `","`)), nil
}

func (a *VerifiedAddresses) Scan(src interface{}) (err error) {
	var verifiedAddresses []string
	switch src := src.(type) {
	case string:
		err = json.Unmarshal([]byte(src), &verifiedAddresses)
	case []byte:
		err = json.Unmarshal(src, &verifiedAddresses)
	default:
		return errors.New("incompatible type for description")
	}
	if err != nil {
		return err
	}
	*a = verifiedAddresses
	return nil
}
