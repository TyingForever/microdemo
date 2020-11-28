package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"microdemo/handler"
	pb "microdemo/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("microdemo"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMicrodemoHandler(srv.Server(), new(handler.Microdemo))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
