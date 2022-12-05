package dtos

type Wallet struct {
	Balance        float32 `json:"balance"`
	AboveThreshold bool    `json:"above_threshold"`
}
