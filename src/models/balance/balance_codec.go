package balance

import (
	"google.golang.org/protobuf/proto"
)

type BalanceCodec struct{}

func (c *BalanceCodec) Encode(value interface{}) ([]byte, error) {
	var balance = value.(*Balance)
	return proto.Marshal(balance)
}

func (c *BalanceCodec) Decode(data []byte) (interface{}, error) {
	balance := &Balance{}
	err := proto.Unmarshal(data, balance)
	return balance, err
}
