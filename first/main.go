package main

import (
	"first/handler"
	pb "first/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("first"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterFirstHandler(srv.Server(), new(handler.First))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
