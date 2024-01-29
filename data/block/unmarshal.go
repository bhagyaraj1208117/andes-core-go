package block

import (
	"github.com/bhagyaraj1208117/andes-core-go/core/check"
	"github.com/bhagyaraj1208117/andes-core-go/data"
	"github.com/bhagyaraj1208117/andes-core-go/marshal"
)

// GetHeaderFromBytes will unmarshal the header bytes based on the header type
func GetHeaderFromBytes(marshaller marshal.Marshalizer, creator EmptyBlockCreator, headerBytes []byte) (data.HeaderHandler, error) {
	if check.IfNil(marshaller) {
		return nil, data.ErrNilMarshalizer
	}
	if check.IfNil(creator) {
		return nil, data.ErrNilEmptyBlockCreator
	}

	header := creator.CreateNewHeader()
	err := marshaller.Unmarshal(header, headerBytes)
	if err != nil {
		return nil, err
	}

	return header, nil
}
