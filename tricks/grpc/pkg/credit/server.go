package credit

import (
	"context"

	"github.com/triumphpc/go-study/tricks/grpc/pkg/api/proto"
)

type CreditServer struct {
	proto.CreditServiceServer
}

func (c *CreditServer) Credit(ctx context.Context, req *proto.CreditRequest) (*proto.CreditResponse, error) {
	var response proto.CreditResponse

	response.Confirmation = "Hello world"

	return &response, nil
}
