package dtos

type Wallet struct {
	WalletId       int     `json:"wallet_id"`
	Balance        float32 `json:"balance"`
	AboveThreshold bool    `json:"above_threshold"`
}
