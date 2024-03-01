package types

type Config struct {
	NodeUrl string `mapstructure:"NODE_URL"`
	Key     string `mapstructure:"KEY"`
	Mode    string `mapstructure:"MODE"`
	Port    int    `mapstructure:"PORT"`
}

type ExecutePayload struct {
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
