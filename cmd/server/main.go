package main

import (
	"fmt"
	"log"

	db "github.com/98y7tbnb97t/users-service/internal/database"
	"github.com/98y7tbnb97t/users-service/internal/transport/grpc"
	"github.com/98y7tbnb97t/users-service/internal/user"
)

func main() {
	db.InitDB()
	repo := user.NewRepository(db.DB)
	svc := user.NewService(repo)

	fmt.Println("server run on port: 50051")
	if err := grpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC server stopped with error: %v", err)
	}
	fmt.Println("server run on port: 50051")
}
