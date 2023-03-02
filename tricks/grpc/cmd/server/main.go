package main

import (
	"fmt"
	"log"
	"net"

	"github.com/triumphpc/go-study/tricks/grpc/pkg/api/proto"
	"github.com/triumphpc/go-study/tricks/grpc/pkg/credit"
	"google.golang.org/grpc"
)

func main() {
	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()
	// регистрируем сервис
	proto.RegisterCreditServiceServer(s, &credit.CreditServer{})

	fmt.Println("сервер gRPC начал работу")
	// получаем запрос gRpc
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
