package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	first "first/proto"
)

type First struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *First) Call(ctx context.Context, req *first.Request, rsp *first.Response) error {
	log.Info("Received First.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *First) Stream(ctx context.Context, req *first.StreamingRequest, stream first.First_StreamStream) error {
	log.Infof("Received First.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&first.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *First) PingPong(ctx context.Context, stream first.First_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&first.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
