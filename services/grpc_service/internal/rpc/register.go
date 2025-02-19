package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RPCRegister struct {
	g *grpc.Server
}

func NewRPCRegister(g *grpc.Server) *RPCRegister {
	return &RPCRegister{g: g}
}

func (r *RPCRegister) Register() {
	// cache := cache.NewMemoryCache(context.Background())

	// service := NewAuthService(cache)
	// authServer := &authServer{
	// 	authService: service,
	// }

	// pb.RegisterAuthenticationServiceServer(r.g, authServer)

	reflection.Register(r.g)
}
