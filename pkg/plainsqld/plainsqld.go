package server

import (
	"context"
	"fmt"
	"log"
	"net"

	rq "github.com/vrecan/PlainSQL/pkg/request"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//TODO: Add testing ability to use in memory messaging

//Opts are the options for the SQLD server
type Opts struct {
	GRPCMaxMsgSize           int64  `long:"max-msg-size" description:"max msg size in bytes"`
	GRPCMaxConcurrentStreams uint32 `long:"max-concurrent-streams" description:"max number of grpc streams"`
	Port                     int32  `long:"port" default:"8091" description:"port used by server"`
	DisableTLS               bool   `long:"disable-tls" description:"Disable TLS encyrption"`
	CertFile                 string `long:"certfile" description:"Certfile location" default:"/usr/local/logrhythm/winston/winston.pem"`
	CertKey                  string `long:"certkey" description:"certkey location" default:"/usr/local/logrhythm/winston/winston.key"`
}

//NewPlainSQLD creates a new grpc server
func NewPlainSQLD(opts Opts) *PlainSQLD {
	grpcOpts := make([]grpc.ServerOption, 0)

	if !opts.DisableTLS {
		creds, err := credentials.NewServerTLSFromFile(opts.CertFile, opts.CertKey)
		if err != nil {
			panic(fmt.Sprintf("Failed to generate credentials %v", err))
		}
		grpcOpts = append(grpcOpts, grpc.Creds(creds))
	}
	grpcOpts = append(grpcOpts, grpc.MaxConcurrentStreams(uint32(32)))
	return &PlainSQLD{Opts: opts,
		srv: grpc.NewServer(grpcOpts...),
	}
}

//PlainSQLD runs the PlainSQLD grpc API server
type PlainSQLD struct {
	Opts Opts
	srv  *grpc.Server
}

//Echo returns an echo of what you send it
func (p *PlainSQLD) Echo(ctx context.Context, req *rq.Request) (*rq.Request, error) {
	return req, nil
}

//Start initialized the server in a background thread
func (p *PlainSQLD) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", p.Opts.Port))
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	log.Println("Winston listening on: ", lis.Addr().String())
	rq.RegisterPlainSQLServer(p.srv, p)

	go p.srv.Serve(lis)

}

//Close stop the server
func (p *PlainSQLD) Close() error {
	p.srv.Stop()
	return nil
}
