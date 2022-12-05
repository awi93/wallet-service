package deposit

import (
	"google.golang.org/protobuf/proto"
)

type DepositCodec struct{}

func (c *DepositCodec) Encode(value interface{}) ([]byte, error) {
	var deposit = value.(*Deposit)
	return proto.Marshal(deposit)
}

func (c *DepositCodec) Decode(data []byte) (interface{}, error) {
	deposit := &Deposit{}
	err := proto.Unmarshal(data, deposit)
	return deposit, err
}
