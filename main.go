package main

import (
	"context"

	"github.com/VladZheleznov/shopping-basket/internal/model"
	"github.com/VladZheleznov/shopping-basket/internal/repository"
	"github.com/VladZheleznov/shopping-basket/internal/server"
	"github.com/VladZheleznov/shopping-basket/internal/service"
	pb "github.com/VladZheleznov/shopping-basket/proto"
	"github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"

	"fmt"
	"net"

	"google.golang.org/grpc"
)

var (
	poolP pgxpool.Pool
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		defer log.Fatalf("error while listening port: %e", err)
	}
	fmt.Println("Server successfully started on port :50051...")
	cfg := model.Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to start service, %e", err)
	}
	conn := DBConnection(&cfg)
	fmt.Println("DB successfully connect...")
	defer func() {
		poolP.Close()
	}()
	ns := grpc.NewServer()
	newService := service.NewService(conn)
	srv := server.NewServer(newService)
	pb.RegisterCRUDServer(ns, srv)

	if err = ns.Serve(listen); err != nil {
		defer log.Fatalf("error while listening server: %e", err)
	}

}

func DBConnection(cfg *model.Config) repository.Repository {

	log.Info(cfg.PostgresDBURL)
	poolP, err := pgxpool.Connect(context.Background(), "postgres://postgres:1111@localhost:5432/test_db")
	if err != nil {
		log.Fatalf("bad connection with postgresql: %v", err)
		return nil
	}
	return &repository.PRepository{Pool: poolP}

}
