package common

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Conn *grpc.ClientConn
)

// grpc连接
func Grpc_conn() *grpc.ClientConn {
	conn, err := grpc.Dial("127.0.0.1:8008", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: #{err}")
	}
	Conn = conn
	return conn
}

func GetConn() *grpc.ClientConn {
	return Conn
}
