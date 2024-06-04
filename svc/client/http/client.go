// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: a2b01cac16
// Version Date: 2022-10-20T18:44:52Z

// Package http provides an HTTP client for the Vpnsvc service.
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gogo/protobuf/jsonpb"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"

	// This Service
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"github.com/mises-id/mises-vpnsvc/svc"
)

var (
	_ = endpoint.Chain
	_ = httptransport.NewClient
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = ioutil.NopCloser
	_ = io.EOF
)

// New returns a service backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options ...httptransport.ClientOption) (pb.VpnsvcServer, error) {

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	_ = u

	var CreateOrderZeroEndpoint endpoint.Endpoint
	{
		CreateOrderZeroEndpoint = httptransport.NewClient(
			"POST",
			copyURL(u, "/order/create/"),
			EncodeHTTPCreateOrderZeroRequest,
			DecodeHTTPCreateOrderResponse,
			options...,
		).Endpoint()
	}
	var UpdateOrderZeroEndpoint endpoint.Endpoint
	{
		UpdateOrderZeroEndpoint = httptransport.NewClient(
			"POST",
			copyURL(u, "/order/update/"),
			EncodeHTTPUpdateOrderZeroRequest,
			DecodeHTTPUpdateOrderResponse,
			options...,
		).Endpoint()
	}
	var VpnInfoZeroEndpoint endpoint.Endpoint
	{
		VpnInfoZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/info/"),
			EncodeHTTPVpnInfoZeroRequest,
			DecodeHTTPVpnInfoResponse,
			options...,
		).Endpoint()
	}
	var FetchOrdersZeroEndpoint endpoint.Endpoint
	{
		FetchOrdersZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/orders/"),
			EncodeHTTPFetchOrdersZeroRequest,
			DecodeHTTPFetchOrdersResponse,
			options...,
		).Endpoint()
	}
	var FetchOrderInfoZeroEndpoint endpoint.Endpoint
	{
		FetchOrderInfoZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/order_info/"),
			EncodeHTTPFetchOrderInfoZeroRequest,
			DecodeHTTPFetchOrderInfoResponse,
			options...,
		).Endpoint()
	}
	var GetServerListZeroEndpoint endpoint.Endpoint
	{
		GetServerListZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/server_list/"),
			EncodeHTTPGetServerListZeroRequest,
			DecodeHTTPGetServerListResponse,
			options...,
		).Endpoint()
	}
	var GetServerLinkZeroEndpoint endpoint.Endpoint
	{
		GetServerLinkZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/server_link/"),
			EncodeHTTPGetServerLinkZeroRequest,
			DecodeHTTPGetServerLinkResponse,
			options...,
		).Endpoint()
	}
	var VerifyOrderFromChainZeroEndpoint endpoint.Endpoint
	{
		VerifyOrderFromChainZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/sync_order_from_chain/"),
			EncodeHTTPVerifyOrderFromChainZeroRequest,
			DecodeHTTPVerifyOrderFromChainResponse,
			options...,
		).Endpoint()
	}
	var CleanExpiredVpnLinkZeroEndpoint endpoint.Endpoint
	{
		CleanExpiredVpnLinkZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/clean_expired_vpn_link/"),
			EncodeHTTPCleanExpiredVpnLinkZeroRequest,
			DecodeHTTPCleanExpiredVpnLinkResponse,
			options...,
		).Endpoint()
	}
	var GetVpnConfigZeroEndpoint endpoint.Endpoint
	{
		GetVpnConfigZeroEndpoint = httptransport.NewClient(
			"GET",
			copyURL(u, "/get_vpn_config/"),
			EncodeHTTPGetVpnConfigZeroRequest,
			DecodeHTTPGetVpnConfigResponse,
			options...,
		).Endpoint()
	}

	return svc.Endpoints{
		CreateOrderEndpoint:          CreateOrderZeroEndpoint,
		UpdateOrderEndpoint:          UpdateOrderZeroEndpoint,
		VpnInfoEndpoint:              VpnInfoZeroEndpoint,
		FetchOrdersEndpoint:          FetchOrdersZeroEndpoint,
		FetchOrderInfoEndpoint:       FetchOrderInfoZeroEndpoint,
		GetServerListEndpoint:        GetServerListZeroEndpoint,
		GetServerLinkEndpoint:        GetServerLinkZeroEndpoint,
		VerifyOrderFromChainEndpoint: VerifyOrderFromChainZeroEndpoint,
		CleanExpiredVpnLinkEndpoint:  CleanExpiredVpnLinkZeroEndpoint,
		GetVpnConfigEndpoint:         GetVpnConfigZeroEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

// CtxValuesToSend configures the http client to pull the specified keys out of
// the context and add them to the http request as headers.  Note that keys
// will have net/http.CanonicalHeaderKey called on them before being send over
// the wire and that is the form they will be available in the server context.
func CtxValuesToSend(keys ...string) httptransport.ClientOption {
	return httptransport.ClientBefore(func(ctx context.Context, r *http.Request) context.Context {
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				r.Header.Set(k, v)
			}
		}
		return ctx
	})
}

// HTTP Client Decode

// DecodeHTTPCreateOrderResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded CreateOrderResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPCreateOrderResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.CreateOrderResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPUpdateOrderResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded UpdateOrderResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPUpdateOrderResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.UpdateOrderResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPVpnInfoResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded VpnInfoResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPVpnInfoResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.VpnInfoResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPFetchOrdersResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded FetchOrdersResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPFetchOrdersResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.FetchOrdersResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPFetchOrderInfoResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded FetchOrderInfoResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPFetchOrderInfoResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.FetchOrderInfoResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPGetServerListResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded GetServerListResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPGetServerListResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.GetServerListResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPGetServerLinkResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded GetServerLinkResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPGetServerLinkResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.GetServerLinkResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPVerifyOrderFromChainResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded VerifyOrderFromChainResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPVerifyOrderFromChainResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.VerifyOrderFromChainResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPCleanExpiredVpnLinkResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded CleanExpiredVpnLinkResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPCleanExpiredVpnLinkResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.CleanExpiredVpnLinkResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// DecodeHTTPGetVpnConfigResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded GetVpnConfigResponse response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPGetVpnConfigResponse(_ context.Context, r *http.Response) (interface{}, error) {
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err == io.EOF {
		return nil, errors.New("response http body empty")
	}
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.GetVpnConfigResponse
	if err = jsonpb.UnmarshalString(string(buf), &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// HTTP Client Encode

// EncodeHTTPCreateOrderZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a createorder request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPCreateOrderZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.CreateOrderRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"order",
		"create",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	r.URL.RawQuery = values.Encode()
	// Set the body parameters
	var buf bytes.Buffer
	toRet := request.(*pb.CreateOrderRequest)

	toRet.EthAddress = req.EthAddress

	toRet.ChainId = req.ChainId

	toRet.PlanId = req.PlanId

	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(toRet); err != nil {
		return errors.Wrapf(err, "couldn't encode body as json %v", toRet)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// EncodeHTTPCreateOrderOneRequest is a transport/http.EncodeRequestFunc
// that encodes a createorder request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPCreateOrderOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.CreateOrderRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"order",
		"create",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	r.URL.RawQuery = values.Encode()
	// Set the body parameters
	var buf bytes.Buffer
	toRet := request.(*pb.CreateOrderRequest)

	toRet.EthAddress = req.EthAddress

	toRet.ChainId = req.ChainId

	toRet.PlanId = req.PlanId

	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(toRet); err != nil {
		return errors.Wrapf(err, "couldn't encode body as json %v", toRet)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// EncodeHTTPUpdateOrderZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a updateorder request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPUpdateOrderZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.UpdateOrderRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"order",
		"update",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	r.URL.RawQuery = values.Encode()
	// Set the body parameters
	var buf bytes.Buffer
	toRet := request.(*pb.UpdateOrderRequest)

	toRet.EthAddress = req.EthAddress

	toRet.OrderId = req.OrderId

	toRet.TxnHash = req.TxnHash

	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(toRet); err != nil {
		return errors.Wrapf(err, "couldn't encode body as json %v", toRet)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// EncodeHTTPUpdateOrderOneRequest is a transport/http.EncodeRequestFunc
// that encodes a updateorder request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPUpdateOrderOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.UpdateOrderRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"order",
		"update",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	r.URL.RawQuery = values.Encode()
	// Set the body parameters
	var buf bytes.Buffer
	toRet := request.(*pb.UpdateOrderRequest)

	toRet.EthAddress = req.EthAddress

	toRet.OrderId = req.OrderId

	toRet.TxnHash = req.TxnHash

	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(toRet); err != nil {
		return errors.Wrapf(err, "couldn't encode body as json %v", toRet)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// EncodeHTTPVpnInfoZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a vpninfo request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPVpnInfoZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.VpnInfoRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"info",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPVpnInfoOneRequest is a transport/http.EncodeRequestFunc
// that encodes a vpninfo request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPVpnInfoOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.VpnInfoRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"info",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPFetchOrdersZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a fetchorders request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPFetchOrdersZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.FetchOrdersRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"orders",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPFetchOrdersOneRequest is a transport/http.EncodeRequestFunc
// that encodes a fetchorders request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPFetchOrdersOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.FetchOrdersRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"orders",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPFetchOrderInfoZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a fetchorderinfo request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPFetchOrderInfoZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.FetchOrderInfoRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"order_info",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	values.Add("orderId", fmt.Sprint(req.OrderId))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPFetchOrderInfoOneRequest is a transport/http.EncodeRequestFunc
// that encodes a fetchorderinfo request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPFetchOrderInfoOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.FetchOrderInfoRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"order_info",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	values.Add("orderId", fmt.Sprint(req.OrderId))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPGetServerListZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a getserverlist request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPGetServerListZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.GetServerListRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"server_list",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPGetServerListOneRequest is a transport/http.EncodeRequestFunc
// that encodes a getserverlist request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPGetServerListOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.GetServerListRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"server_list",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPGetServerLinkZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a getserverlink request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPGetServerLinkZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.GetServerLinkRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"server_link",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	values.Add("server", fmt.Sprint(req.Server))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPGetServerLinkOneRequest is a transport/http.EncodeRequestFunc
// that encodes a getserverlink request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPGetServerLinkOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.GetServerLinkRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"server_link",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("ethAddress", fmt.Sprint(req.EthAddress))

	values.Add("server", fmt.Sprint(req.Server))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPVerifyOrderFromChainZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a verifyorderfromchain request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPVerifyOrderFromChainZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.VerifyOrderFromChainRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"sync_order_from_chain",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("chain", fmt.Sprint(req.Chain))

	values.Add("startBlock", fmt.Sprint(req.StartBlock))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPVerifyOrderFromChainOneRequest is a transport/http.EncodeRequestFunc
// that encodes a verifyorderfromchain request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPVerifyOrderFromChainOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.VerifyOrderFromChainRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"sync_order_from_chain",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("chain", fmt.Sprint(req.Chain))

	values.Add("startBlock", fmt.Sprint(req.StartBlock))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPCleanExpiredVpnLinkZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a cleanexpiredvpnlink request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPCleanExpiredVpnLinkZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.CleanExpiredVpnLinkRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"clean_expired_vpn_link",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("endTime", fmt.Sprint(req.EndTime))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPCleanExpiredVpnLinkOneRequest is a transport/http.EncodeRequestFunc
// that encodes a cleanexpiredvpnlink request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPCleanExpiredVpnLinkOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.CleanExpiredVpnLinkRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"clean_expired_vpn_link",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("endTime", fmt.Sprint(req.EndTime))

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPGetVpnConfigZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a getvpnconfig request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPGetVpnConfigZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.GetVpnConfigRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"get_vpn_config",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	r.URL.RawQuery = values.Encode()
	return nil
}

// EncodeHTTPGetVpnConfigOneRequest is a transport/http.EncodeRequestFunc
// that encodes a getvpnconfig request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPGetVpnConfigOneRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.GetVpnConfigRequest)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"get_vpn_config",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	r.URL.RawQuery = values.Encode()
	return nil
}

func errorDecoder(buf []byte) error {
	var w errorWrapper
	if err := json.Unmarshal(buf, &w); err != nil {
		const size = 8196
		if len(buf) > size {
			buf = buf[:size]
		}
		return fmt.Errorf("response body '%s': cannot parse non-json request body", buf)
	}

	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}
