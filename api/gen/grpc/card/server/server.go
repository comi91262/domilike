// Code generated by goa v3.11.0, DO NOT EDIT.
//
// card gRPC server
//
// Command:
// $ goa gen github.com/comi91262/domilike/design

package server

import (
	"context"

	card "github.com/comi91262/domilike/gen/card"
	cardpb "github.com/comi91262/domilike/gen/grpc/card/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the cardpb.CardServer interface.
type Server struct {
	GetH goagrpc.UnaryHandler
	cardpb.UnimplementedCardServer
}

// New instantiates the server struct with the card service endpoints.
func New(e *card.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		GetH: NewGetHandler(e.Get, uh),
	}
}

// NewGetHandler creates a gRPC handler which serves the "card" service "get"
// endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetRequest, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in cardpb.CardServer interface.
func (s *Server) Get(ctx context.Context, message *cardpb.GetRequest) (*cardpb.GetResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "card")
	resp, err := s.GetH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*cardpb.GetResponse), nil
}