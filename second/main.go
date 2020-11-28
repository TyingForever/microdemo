package main

import (
    "github.com/micro/micro/v3/service"
    "github.com/micro/micro/v3/service/logger"
    "microdemo/second/handler"
    pb "microdemo/second/proto"
)

func main() {
    // Create service
    srv := service.New(
        service.Name("second"),
        service.Version("latest"),
    )

    // Register handler
    pb.RegisterSecondHandler(srv.Server(), new(handler.Second))

    // Run service
    if err := srv.Run(); err != nil {
        logger.Fatal(err)
    }
}
