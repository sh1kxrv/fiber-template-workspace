package main

import (
	"net"
	"service/internal/boot"
	"service/internal/config"
	"service/internal/rpc"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	if err := boot.BootViper("config"); err != nil {
		logrus.Fatalf("Failed to initialize config: %s", err.Error())
	}

	cfg := config.InitConfig()

	boot.BootLogrus()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	logrus.Infof("listening on %s", listener.Addr())

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			rpc.AuthInterceptor(cfg.ServiceToken),
		),
	)

	register := rpc.NewRPCRegister(s)
	register.Register()

	s.Serve(listener)
}
