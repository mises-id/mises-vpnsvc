// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: a2b01cac16
// Version Date: 2022-10-20T18:44:52Z

package svc

// This file provides server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"context"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	// This service
	pb "github.com/mises-id/mises-vpnsvc/proto"
)

const contentType = "application/json; charset=utf-8"

var (
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = strconv.Atoi
	_ = httptransport.NewServer
	_ = ioutil.NopCloser
	_ = pb.NewVpnsvcClient
	_ = io.Copy
	_ = errors.Wrap
)

// MakeHTTPHandler returns a handler that makes a set of endpoints available
// on predefined paths.
func MakeHTTPHandler(endpoints Endpoints, responseEncoder httptransport.EncodeResponseFunc, options ...httptransport.ServerOption) http.Handler {
	if responseEncoder == nil {
		responseEncoder = EncodeHTTPGenericResponse
	}
	serverOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(headersToContext),
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerAfter(httptransport.SetContentType(contentType)),
	}
	serverOptions = append(serverOptions, options...)
	m := mux.NewRouter()

	m.Methods("POST").Path("/order/create/").Handler(httptransport.NewServer(
		endpoints.CreateOrderEndpoint,
		DecodeHTTPCreateOrderZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("POST").Path("/order/create").Handler(httptransport.NewServer(
		endpoints.CreateOrderEndpoint,
		DecodeHTTPCreateOrderOneRequest,
		responseEncoder,
		serverOptions...,
	))

	m.Methods("POST").Path("/order/update/").Handler(httptransport.NewServer(
		endpoints.UpdateOrderEndpoint,
		DecodeHTTPUpdateOrderZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("POST").Path("/order/update").Handler(httptransport.NewServer(
		endpoints.UpdateOrderEndpoint,
		DecodeHTTPUpdateOrderOneRequest,
		responseEncoder,
		serverOptions...,
	))

	m.Methods("GET").Path("/info/").Handler(httptransport.NewServer(
		endpoints.VpnInfoEndpoint,
		DecodeHTTPVpnInfoZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("GET").Path("/info").Handler(httptransport.NewServer(
		endpoints.VpnInfoEndpoint,
		DecodeHTTPVpnInfoOneRequest,
		responseEncoder,
		serverOptions...,
	))

	m.Methods("GET").Path("/orders/").Handler(httptransport.NewServer(
		endpoints.FetchOrdersEndpoint,
		DecodeHTTPFetchOrdersZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("GET").Path("/orders").Handler(httptransport.NewServer(
		endpoints.FetchOrdersEndpoint,
		DecodeHTTPFetchOrdersOneRequest,
		responseEncoder,
		serverOptions...,
	))

	m.Methods("GET").Path("/order_info/").Handler(httptransport.NewServer(
		endpoints.FetchOrderInfoEndpoint,
		DecodeHTTPFetchOrderInfoZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("GET").Path("/order_info").Handler(httptransport.NewServer(
		endpoints.FetchOrderInfoEndpoint,
		DecodeHTTPFetchOrderInfoOneRequest,
		responseEncoder,
		serverOptions...,
	))

	m.Methods("GET").Path("/server_list/").Handler(httptransport.NewServer(
		endpoints.GetServerListEndpoint,
		DecodeHTTPGetServerListZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("GET").Path("/server_list").Handler(httptransport.NewServer(
		endpoints.GetServerListEndpoint,
		DecodeHTTPGetServerListOneRequest,
		responseEncoder,
		serverOptions...,
	))

	m.Methods("GET").Path("/server_link/").Handler(httptransport.NewServer(
		endpoints.GetServerLinkEndpoint,
		DecodeHTTPGetServerLinkZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("GET").Path("/server_link").Handler(httptransport.NewServer(
		endpoints.GetServerLinkEndpoint,
		DecodeHTTPGetServerLinkOneRequest,
		responseEncoder,
		serverOptions...,
	))

	m.Methods("GET").Path("/sync_order_from_chain/").Handler(httptransport.NewServer(
		endpoints.VerifyOrderFromChainEndpoint,
		DecodeHTTPVerifyOrderFromChainZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("GET").Path("/sync_order_from_chain").Handler(httptransport.NewServer(
		endpoints.VerifyOrderFromChainEndpoint,
		DecodeHTTPVerifyOrderFromChainOneRequest,
		responseEncoder,
		serverOptions...,
	))

	m.Methods("GET").Path("/clean_expired_vpn_link/").Handler(httptransport.NewServer(
		endpoints.CleanExpiredVpnLinkEndpoint,
		DecodeHTTPCleanExpiredVpnLinkZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("GET").Path("/clean_expired_vpn_link").Handler(httptransport.NewServer(
		endpoints.CleanExpiredVpnLinkEndpoint,
		DecodeHTTPCleanExpiredVpnLinkOneRequest,
		responseEncoder,
		serverOptions...,
	))

	m.Methods("GET").Path("/get_vpn_config/").Handler(httptransport.NewServer(
		endpoints.GetVpnConfigEndpoint,
		DecodeHTTPGetVpnConfigZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	m.Methods("GET").Path("/get_vpn_config").Handler(httptransport.NewServer(
		endpoints.GetVpnConfigEndpoint,
		DecodeHTTPGetVpnConfigOneRequest,
		responseEncoder,
		serverOptions...,
	))
	return m
}

// ErrorEncoder writes the error to the ResponseWriter, by default a content
// type of application/json, a body of json with key "error" and the value
// error.Error(), and a status code of 500. If the error implements Headerer,
// the provided headers will be applied to the response. If the error
// implements json.Marshaler, and the marshaling succeeds, the JSON encoded
// form of the error will be used. If the error implements StatusCoder, the
// provided StatusCode will be used instead of 500.
func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	body, _ := json.Marshal(errorWrapper{Error: err.Error()})
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			body = jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if headerer, ok := err.(httptransport.Headerer); ok {
		for k := range headerer.Headers() {
			w.Header().Set(k, headerer.Headers().Get(k))
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(httptransport.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	w.Write(body)
}

type errorWrapper struct {
	Error string `json:"error"`
}

// httpError satisfies the Headerer and StatusCoder interfaces in
// package github.com/go-kit/kit/transport/http.
type httpError struct {
	error
	statusCode int
	headers    map[string][]string
}

func (h httpError) StatusCode() int {
	return h.statusCode
}

func (h httpError) Headers() http.Header {
	return h.headers
}

// Server Decode

// DecodeHTTPCreateOrderZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded createorder request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPCreateOrderZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.CreateOrderRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPCreateOrderOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded createorder request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPCreateOrderOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.CreateOrderRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPUpdateOrderZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded updateorder request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPUpdateOrderZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.UpdateOrderRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPUpdateOrderOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded updateorder request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPUpdateOrderOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.UpdateOrderRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPVpnInfoZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded vpninfo request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPVpnInfoZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.VpnInfoRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressVpnInfoStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressVpnInfoStr := EthAddressVpnInfoStrArr[0]
		EthAddressVpnInfo := EthAddressVpnInfoStr
		req.EthAddress = EthAddressVpnInfo
	}

	return &req, err
}

// DecodeHTTPVpnInfoOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded vpninfo request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPVpnInfoOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.VpnInfoRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressVpnInfoStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressVpnInfoStr := EthAddressVpnInfoStrArr[0]
		EthAddressVpnInfo := EthAddressVpnInfoStr
		req.EthAddress = EthAddressVpnInfo
	}

	return &req, err
}

// DecodeHTTPFetchOrdersZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded fetchorders request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPFetchOrdersZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.FetchOrdersRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressFetchOrdersStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressFetchOrdersStr := EthAddressFetchOrdersStrArr[0]
		EthAddressFetchOrders := EthAddressFetchOrdersStr
		req.EthAddress = EthAddressFetchOrders
	}

	return &req, err
}

// DecodeHTTPFetchOrdersOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded fetchorders request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPFetchOrdersOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.FetchOrdersRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressFetchOrdersStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressFetchOrdersStr := EthAddressFetchOrdersStrArr[0]
		EthAddressFetchOrders := EthAddressFetchOrdersStr
		req.EthAddress = EthAddressFetchOrders
	}

	return &req, err
}

// DecodeHTTPFetchOrderInfoZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded fetchorderinfo request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPFetchOrderInfoZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.FetchOrderInfoRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressFetchOrderInfoStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressFetchOrderInfoStr := EthAddressFetchOrderInfoStrArr[0]
		EthAddressFetchOrderInfo := EthAddressFetchOrderInfoStr
		req.EthAddress = EthAddressFetchOrderInfo
	}

	if OrderIdFetchOrderInfoStrArr, ok := queryParams["orderId"]; ok {
		OrderIdFetchOrderInfoStr := OrderIdFetchOrderInfoStrArr[0]
		OrderIdFetchOrderInfo := OrderIdFetchOrderInfoStr
		req.OrderId = OrderIdFetchOrderInfo
	}

	return &req, err
}

// DecodeHTTPFetchOrderInfoOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded fetchorderinfo request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPFetchOrderInfoOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.FetchOrderInfoRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressFetchOrderInfoStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressFetchOrderInfoStr := EthAddressFetchOrderInfoStrArr[0]
		EthAddressFetchOrderInfo := EthAddressFetchOrderInfoStr
		req.EthAddress = EthAddressFetchOrderInfo
	}

	if OrderIdFetchOrderInfoStrArr, ok := queryParams["orderId"]; ok {
		OrderIdFetchOrderInfoStr := OrderIdFetchOrderInfoStrArr[0]
		OrderIdFetchOrderInfo := OrderIdFetchOrderInfoStr
		req.OrderId = OrderIdFetchOrderInfo
	}

	return &req, err
}

// DecodeHTTPGetServerListZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded getserverlist request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPGetServerListZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.GetServerListRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressGetServerListStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressGetServerListStr := EthAddressGetServerListStrArr[0]
		EthAddressGetServerList := EthAddressGetServerListStr
		req.EthAddress = EthAddressGetServerList
	}

	return &req, err
}

// DecodeHTTPGetServerListOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded getserverlist request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPGetServerListOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.GetServerListRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressGetServerListStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressGetServerListStr := EthAddressGetServerListStrArr[0]
		EthAddressGetServerList := EthAddressGetServerListStr
		req.EthAddress = EthAddressGetServerList
	}

	return &req, err
}

// DecodeHTTPGetServerLinkZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded getserverlink request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPGetServerLinkZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.GetServerLinkRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressGetServerLinkStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressGetServerLinkStr := EthAddressGetServerLinkStrArr[0]
		EthAddressGetServerLink := EthAddressGetServerLinkStr
		req.EthAddress = EthAddressGetServerLink
	}

	if ServerGetServerLinkStrArr, ok := queryParams["server"]; ok {
		ServerGetServerLinkStr := ServerGetServerLinkStrArr[0]
		ServerGetServerLink := ServerGetServerLinkStr
		req.Server = ServerGetServerLink
	}

	return &req, err
}

// DecodeHTTPGetServerLinkOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded getserverlink request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPGetServerLinkOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.GetServerLinkRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EthAddressGetServerLinkStrArr, ok := queryParams["ethAddress"]; ok {
		EthAddressGetServerLinkStr := EthAddressGetServerLinkStrArr[0]
		EthAddressGetServerLink := EthAddressGetServerLinkStr
		req.EthAddress = EthAddressGetServerLink
	}

	if ServerGetServerLinkStrArr, ok := queryParams["server"]; ok {
		ServerGetServerLinkStr := ServerGetServerLinkStrArr[0]
		ServerGetServerLink := ServerGetServerLinkStr
		req.Server = ServerGetServerLink
	}

	return &req, err
}

// DecodeHTTPVerifyOrderFromChainZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded verifyorderfromchain request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPVerifyOrderFromChainZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.VerifyOrderFromChainRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if ChainVerifyOrderFromChainStrArr, ok := queryParams["chain"]; ok {
		ChainVerifyOrderFromChainStr := ChainVerifyOrderFromChainStrArr[0]
		ChainVerifyOrderFromChain, err := strconv.ParseUint(ChainVerifyOrderFromChainStr, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting ChainVerifyOrderFromChain from query, queryParams: %v", queryParams))
		}
		req.Chain = ChainVerifyOrderFromChain
	}

	if StartBlockVerifyOrderFromChainStrArr, ok := queryParams["startBlock"]; ok {
		StartBlockVerifyOrderFromChainStr := StartBlockVerifyOrderFromChainStrArr[0]
		StartBlockVerifyOrderFromChain, err := strconv.ParseInt(StartBlockVerifyOrderFromChainStr, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting StartBlockVerifyOrderFromChain from query, queryParams: %v", queryParams))
		}
		req.StartBlock = StartBlockVerifyOrderFromChain
	}

	return &req, err
}

// DecodeHTTPVerifyOrderFromChainOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded verifyorderfromchain request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPVerifyOrderFromChainOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.VerifyOrderFromChainRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if ChainVerifyOrderFromChainStrArr, ok := queryParams["chain"]; ok {
		ChainVerifyOrderFromChainStr := ChainVerifyOrderFromChainStrArr[0]
		ChainVerifyOrderFromChain, err := strconv.ParseUint(ChainVerifyOrderFromChainStr, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting ChainVerifyOrderFromChain from query, queryParams: %v", queryParams))
		}
		req.Chain = ChainVerifyOrderFromChain
	}

	if StartBlockVerifyOrderFromChainStrArr, ok := queryParams["startBlock"]; ok {
		StartBlockVerifyOrderFromChainStr := StartBlockVerifyOrderFromChainStrArr[0]
		StartBlockVerifyOrderFromChain, err := strconv.ParseInt(StartBlockVerifyOrderFromChainStr, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting StartBlockVerifyOrderFromChain from query, queryParams: %v", queryParams))
		}
		req.StartBlock = StartBlockVerifyOrderFromChain
	}

	return &req, err
}

// DecodeHTTPCleanExpiredVpnLinkZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded cleanexpiredvpnlink request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPCleanExpiredVpnLinkZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.CleanExpiredVpnLinkRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EndTimeCleanExpiredVpnLinkStrArr, ok := queryParams["endTime"]; ok {
		EndTimeCleanExpiredVpnLinkStr := EndTimeCleanExpiredVpnLinkStrArr[0]
		EndTimeCleanExpiredVpnLink, err := strconv.ParseInt(EndTimeCleanExpiredVpnLinkStr, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting EndTimeCleanExpiredVpnLink from query, queryParams: %v", queryParams))
		}
		req.EndTime = EndTimeCleanExpiredVpnLink
	}

	return &req, err
}

// DecodeHTTPCleanExpiredVpnLinkOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded cleanexpiredvpnlink request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPCleanExpiredVpnLinkOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.CleanExpiredVpnLinkRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	if EndTimeCleanExpiredVpnLinkStrArr, ok := queryParams["endTime"]; ok {
		EndTimeCleanExpiredVpnLinkStr := EndTimeCleanExpiredVpnLinkStrArr[0]
		EndTimeCleanExpiredVpnLink, err := strconv.ParseInt(EndTimeCleanExpiredVpnLinkStr, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting EndTimeCleanExpiredVpnLink from query, queryParams: %v", queryParams))
		}
		req.EndTime = EndTimeCleanExpiredVpnLink
	}

	return &req, err
}

// DecodeHTTPGetVpnConfigZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded getvpnconfig request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPGetVpnConfigZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.GetVpnConfigRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// DecodeHTTPGetVpnConfigOneRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded getvpnconfig request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPGetVpnConfigOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.GetVpnConfigRequest
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	return &req, err
}

// EncodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeHTTPGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	marshaller := jsonpb.Marshaler{
		EmitDefaults: false,
		OrigName:     true,
	}

	return marshaller.Marshal(w, response.(proto.Message))
}

// Helper functions

func headersToContext(ctx context.Context, r *http.Request) context.Context {
	for k := range r.Header {
		// The key is added both in http format (k) which has had
		// http.CanonicalHeaderKey called on it in transport as well as the
		// strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		ctx = context.WithValue(ctx, k, r.Header.Get(k))
		ctx = context.WithValue(ctx, strings.ToLower(k), r.Header.Get(k))
	}

	// Tune specific change.
	// also add the request url
	ctx = context.WithValue(ctx, "request-url", r.URL.Path)
	ctx = context.WithValue(ctx, "transport", "HTTPJSON")

	return ctx
}

// encodePathParams encodes `mux.Vars()` with dot notations into JSON objects
// to be unmarshaled into non-basetype fields.
// e.g. {"book.name": "books/1"} -> {"book": {"name": "books/1"}}
func encodePathParams(vars map[string]string) map[string]string {
	var recur func(path, value string, data map[string]interface{})
	recur = func(path, value string, data map[string]interface{}) {
		parts := strings.SplitN(path, ".", 2)
		key := parts[0]
		if len(parts) == 1 {
			data[key] = value
		} else {
			if _, ok := data[key]; !ok {
				data[key] = make(map[string]interface{})
			}
			recur(parts[1], value, data[key].(map[string]interface{}))
		}
	}

	data := make(map[string]interface{})
	for key, val := range vars {
		recur(key, val, data)
	}

	ret := make(map[string]string)
	for key, val := range data {
		switch val := val.(type) {
		case string:
			ret[key] = val
		case map[string]interface{}:
			m, _ := json.Marshal(val)
			ret[key] = string(m)
		}
	}
	return ret
}
