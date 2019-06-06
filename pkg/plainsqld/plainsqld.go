package server

import (
	"context"
	"fmt"
	"log"
	"net"

	rq "github.com/vrecan/PlainSQL/pkg/request"
	"google.golang.org/grpc"
)

//NewPlainSQLD creates a new grpc server
func NewPlainSQLD(port int) *PlainSQLD {
	return &PlainSQLD{port: port,
		srv: grpc.NewServer(),
	}
}

//PlainSQLD runs the PlainSQLD grpc API server
type PlainSQLD struct {
	port int
	srv  *grpc.Server
}

//Echo returns an echo of what you send it
func (p *PlainSQLD) Echo(ctx context.Context, req *rq.Request) (*rq.Request, error) {
	return req, nil
}

//Start initialized the server in a background thread
func (p *PlainSQLD) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", p.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		p.srv.Serve(lis)
	}()

}

//Close stop the server
func (p *PlainSQLD) Close() error {
	p.srv.Stop()
	return nil
}
