package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	second "second/proto"
)

type Second struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Second) Call(ctx context.Context, req *second.Request, rsp *second.Response) error {
	log.Info("Received Second.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Second) Stream(ctx context.Context, req *second.StreamingRequest, stream second.Second_StreamStream) error {
	log.Infof("Received Second.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&second.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Second) PingPong(ctx context.Context, stream second.Second_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&second.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
