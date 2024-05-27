// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: a2b01cac16
// Version Date: 2022-10-20T18:44:52Z

// Package grpc provides a gRPC client for the Vpnsvc service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"github.com/mises-id/mises-vpnsvc/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.VpnsvcServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var createorderEndpoint endpoint.Endpoint
	{
		createorderEndpoint = grpctransport.NewClient(
			conn,
			"vpnsvc.Vpnsvc",
			"CreateOrder",
			EncodeGRPCCreateOrderRequest,
			DecodeGRPCCreateOrderResponse,
			pb.CreateOrderResponse{},
			clientOptions...,
		).Endpoint()
	}

	var updateorderEndpoint endpoint.Endpoint
	{
		updateorderEndpoint = grpctransport.NewClient(
			conn,
			"vpnsvc.Vpnsvc",
			"UpdateOrder",
			EncodeGRPCUpdateOrderRequest,
			DecodeGRPCUpdateOrderResponse,
			pb.UpdateOrderResponse{},
			clientOptions...,
		).Endpoint()
	}

	var vpninfoEndpoint endpoint.Endpoint
	{
		vpninfoEndpoint = grpctransport.NewClient(
			conn,
			"vpnsvc.Vpnsvc",
			"VpnInfo",
			EncodeGRPCVpnInfoRequest,
			DecodeGRPCVpnInfoResponse,
			pb.VpnInfoResponse{},
			clientOptions...,
		).Endpoint()
	}

	var fetchordersEndpoint endpoint.Endpoint
	{
		fetchordersEndpoint = grpctransport.NewClient(
			conn,
			"vpnsvc.Vpnsvc",
			"FetchOrders",
			EncodeGRPCFetchOrdersRequest,
			DecodeGRPCFetchOrdersResponse,
			pb.FetchOrdersResponse{},
			clientOptions...,
		).Endpoint()
	}

	var fetchorderinfoEndpoint endpoint.Endpoint
	{
		fetchorderinfoEndpoint = grpctransport.NewClient(
			conn,
			"vpnsvc.Vpnsvc",
			"FetchOrderInfo",
			EncodeGRPCFetchOrderInfoRequest,
			DecodeGRPCFetchOrderInfoResponse,
			pb.FetchOrderInfoResponse{},
			clientOptions...,
		).Endpoint()
	}

	var getserverlistEndpoint endpoint.Endpoint
	{
		getserverlistEndpoint = grpctransport.NewClient(
			conn,
			"vpnsvc.Vpnsvc",
			"GetServerList",
			EncodeGRPCGetServerListRequest,
			DecodeGRPCGetServerListResponse,
			pb.GetServerListResponse{},
			clientOptions...,
		).Endpoint()
	}

	var getserverlinkEndpoint endpoint.Endpoint
	{
		getserverlinkEndpoint = grpctransport.NewClient(
			conn,
			"vpnsvc.Vpnsvc",
			"GetServerLink",
			EncodeGRPCGetServerLinkRequest,
			DecodeGRPCGetServerLinkResponse,
			pb.GetServerLinkResponse{},
			clientOptions...,
		).Endpoint()
	}

	var verifyorderfromchainEndpoint endpoint.Endpoint
	{
		verifyorderfromchainEndpoint = grpctransport.NewClient(
			conn,
			"vpnsvc.Vpnsvc",
			"VerifyOrderFromChain",
			EncodeGRPCVerifyOrderFromChainRequest,
			DecodeGRPCVerifyOrderFromChainResponse,
			pb.VerifyOrderFromChainResponse{},
			clientOptions...,
		).Endpoint()
	}

	var cleanexpiredvpnlinkEndpoint endpoint.Endpoint
	{
		cleanexpiredvpnlinkEndpoint = grpctransport.NewClient(
			conn,
			"vpnsvc.Vpnsvc",
			"CleanExpiredVpnLink",
			EncodeGRPCCleanExpiredVpnLinkRequest,
			DecodeGRPCCleanExpiredVpnLinkResponse,
			pb.CleanExpiredVpnLinkResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		CreateOrderEndpoint:          createorderEndpoint,
		UpdateOrderEndpoint:          updateorderEndpoint,
		VpnInfoEndpoint:              vpninfoEndpoint,
		FetchOrdersEndpoint:          fetchordersEndpoint,
		FetchOrderInfoEndpoint:       fetchorderinfoEndpoint,
		GetServerListEndpoint:        getserverlistEndpoint,
		GetServerLinkEndpoint:        getserverlinkEndpoint,
		VerifyOrderFromChainEndpoint: verifyorderfromchainEndpoint,
		CleanExpiredVpnLinkEndpoint:  cleanexpiredvpnlinkEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCCreateOrderResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC createorder reply to a user-domain createorder response. Primarily useful in a client.
func DecodeGRPCCreateOrderResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateOrderResponse)
	return reply, nil
}

// DecodeGRPCUpdateOrderResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC updateorder reply to a user-domain updateorder response. Primarily useful in a client.
func DecodeGRPCUpdateOrderResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UpdateOrderResponse)
	return reply, nil
}

// DecodeGRPCVpnInfoResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC vpninfo reply to a user-domain vpninfo response. Primarily useful in a client.
func DecodeGRPCVpnInfoResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.VpnInfoResponse)
	return reply, nil
}

// DecodeGRPCFetchOrdersResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC fetchorders reply to a user-domain fetchorders response. Primarily useful in a client.
func DecodeGRPCFetchOrdersResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.FetchOrdersResponse)
	return reply, nil
}

// DecodeGRPCFetchOrderInfoResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC fetchorderinfo reply to a user-domain fetchorderinfo response. Primarily useful in a client.
func DecodeGRPCFetchOrderInfoResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.FetchOrderInfoResponse)
	return reply, nil
}

// DecodeGRPCGetServerListResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getserverlist reply to a user-domain getserverlist response. Primarily useful in a client.
func DecodeGRPCGetServerListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetServerListResponse)
	return reply, nil
}

// DecodeGRPCGetServerLinkResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getserverlink reply to a user-domain getserverlink response. Primarily useful in a client.
func DecodeGRPCGetServerLinkResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetServerLinkResponse)
	return reply, nil
}

// DecodeGRPCVerifyOrderFromChainResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC verifyorderfromchain reply to a user-domain verifyorderfromchain response. Primarily useful in a client.
func DecodeGRPCVerifyOrderFromChainResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.VerifyOrderFromChainResponse)
	return reply, nil
}

// DecodeGRPCCleanExpiredVpnLinkResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC cleanexpiredvpnlink reply to a user-domain cleanexpiredvpnlink response. Primarily useful in a client.
func DecodeGRPCCleanExpiredVpnLinkResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CleanExpiredVpnLinkResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCCreateOrderRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createorder request to a gRPC createorder request. Primarily useful in a client.
func EncodeGRPCCreateOrderRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateOrderRequest)
	return req, nil
}

// EncodeGRPCUpdateOrderRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain updateorder request to a gRPC updateorder request. Primarily useful in a client.
func EncodeGRPCUpdateOrderRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateOrderRequest)
	return req, nil
}

// EncodeGRPCVpnInfoRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain vpninfo request to a gRPC vpninfo request. Primarily useful in a client.
func EncodeGRPCVpnInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.VpnInfoRequest)
	return req, nil
}

// EncodeGRPCFetchOrdersRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain fetchorders request to a gRPC fetchorders request. Primarily useful in a client.
func EncodeGRPCFetchOrdersRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.FetchOrdersRequest)
	return req, nil
}

// EncodeGRPCFetchOrderInfoRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain fetchorderinfo request to a gRPC fetchorderinfo request. Primarily useful in a client.
func EncodeGRPCFetchOrderInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.FetchOrderInfoRequest)
	return req, nil
}

// EncodeGRPCGetServerListRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getserverlist request to a gRPC getserverlist request. Primarily useful in a client.
func EncodeGRPCGetServerListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetServerListRequest)
	return req, nil
}

// EncodeGRPCGetServerLinkRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getserverlink request to a gRPC getserverlink request. Primarily useful in a client.
func EncodeGRPCGetServerLinkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetServerLinkRequest)
	return req, nil
}

// EncodeGRPCVerifyOrderFromChainRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain verifyorderfromchain request to a gRPC verifyorderfromchain request. Primarily useful in a client.
func EncodeGRPCVerifyOrderFromChainRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.VerifyOrderFromChainRequest)
	return req, nil
}

// EncodeGRPCCleanExpiredVpnLinkRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain cleanexpiredvpnlink request to a gRPC cleanexpiredvpnlink request. Primarily useful in a client.
func EncodeGRPCCleanExpiredVpnLinkRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CleanExpiredVpnLinkRequest)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
