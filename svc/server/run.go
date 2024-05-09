// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: a2b01cac16
// Version Date: 2022-10-20T18:44:52Z

package server

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"

	// 3d Party
	"google.golang.org/grpc"

	// This Service
	"github.com/mises-id/mises-vpnsvc/handlers"
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"github.com/mises-id/mises-vpnsvc/svc"
)

var DefaultConfig svc.Config

func init() {
	flag.StringVar(&DefaultConfig.DebugAddr, "debug.addr", ":5060", "Debug and metrics listen address")
	flag.StringVar(&DefaultConfig.HTTPAddr, "http.addr", ":5050", "HTTP listen address")
	flag.StringVar(&DefaultConfig.GRPCAddr, "grpc.addr", ":5040", "gRPC (HTTP) listen address")

	// Use environment variables, if set. Flags have priority over Env vars.
	if addr := os.Getenv("DEBUG_ADDR"); addr != "" {
		DefaultConfig.DebugAddr = addr
	}
	if port := os.Getenv("PORT"); port != "" {
		DefaultConfig.HTTPAddr = fmt.Sprintf(":%s", port)
	}
	if addr := os.Getenv("HTTP_ADDR"); addr != "" {
		DefaultConfig.HTTPAddr = addr
	}
	if addr := os.Getenv("GRPC_ADDR"); addr != "" {
		DefaultConfig.GRPCAddr = addr
	}
}

func NewEndpoints(service pb.VpnsvcServer) svc.Endpoints {
	// Business domain.

	// Wrap Service with middlewares. See handlers/middlewares.go
	service = handlers.WrapService(service)

	// Endpoint domain.
	var (
		createorderEndpoint    = svc.MakeCreateOrderEndpoint(service)
		updateorderEndpoint    = svc.MakeUpdateOrderEndpoint(service)
		vpninfoEndpoint        = svc.MakeVpnInfoEndpoint(service)
		fetchordersEndpoint    = svc.MakeFetchOrdersEndpoint(service)
		fetchorderinfoEndpoint = svc.MakeFetchOrderInfoEndpoint(service)
	)

	endpoints := svc.Endpoints{
		CreateOrderEndpoint:    createorderEndpoint,
		UpdateOrderEndpoint:    updateorderEndpoint,
		VpnInfoEndpoint:        vpninfoEndpoint,
		FetchOrdersEndpoint:    fetchordersEndpoint,
		FetchOrderInfoEndpoint: fetchorderinfoEndpoint,
	}

	// Wrap selected Endpoints with middlewares. See handlers/middlewares.go
	endpoints = handlers.WrapEndpoints(endpoints)

	return endpoints
}

// Run starts a new http server, gRPC server, and a debug server with the
// passed config and logger
func Run(cfg svc.Config) {
	service := handlers.NewService()
	endpoints := NewEndpoints(service)

	if cfg.GenericHTTPResponseEncoder == nil {
		cfg.GenericHTTPResponseEncoder = svc.EncodeHTTPGenericResponse
	}

	// Mechanical domain.
	errc := make(chan error)

	// Interrupt handler.
	go handlers.InterruptHandler(errc)

	// Debug listener.
	go func() {
		log.Println("transport", "debug", "addr", cfg.DebugAddr)

		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

		errc <- http.ListenAndServe(cfg.DebugAddr, m)
	}()

	// HTTP transport.
	go func() {
		log.Println("transport", "HTTP", "addr", cfg.HTTPAddr)
		h := svc.MakeHTTPHandler(endpoints, cfg.GenericHTTPResponseEncoder)
		errc <- http.ListenAndServe(cfg.HTTPAddr, h)
	}()

	// gRPC transport.
	go func() {
		log.Println("transport", "gRPC", "addr", cfg.GRPCAddr)
		ln, err := net.Listen("tcp", cfg.GRPCAddr)
		if err != nil {
			errc <- err
			return
		}

		srv := svc.MakeGRPCServer(endpoints)
		s := grpc.NewServer()
		pb.RegisterVpnsvcServer(s, srv)

		errc <- s.Serve(ln)
	}()

	// Run!
	log.Println("exit", <-errc)
}
