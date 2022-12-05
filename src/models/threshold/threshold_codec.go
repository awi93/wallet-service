package threshold

import (
	"google.golang.org/protobuf/proto"
)

type ThresholdCodec struct{}

func (c *ThresholdCodec) Encode(value interface{}) ([]byte, error) {
	var threshold = value.(*Threshold)
	return proto.Marshal(threshold)
}

func (c *ThresholdCodec) Decode(data []byte) (interface{}, error) {
	threshold := &Threshold{}
	err := proto.Unmarshal(data, threshold)
	return threshold, err
}
