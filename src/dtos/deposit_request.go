package dtos

type DepositRequest struct {
	WalletId int32   `json:"wallet_id"`
	Amount   float32 `json:"amount"`
}
