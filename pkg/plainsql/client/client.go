package client

import (
	"context"
	"log"

	"crypto/tls"

	"github.com/vrecan/PlainSQL/pkg/request"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

//Client is a GRPC client for requests
type Client struct {
	Opts   ClientOpts
	Client request.PlainSQLClient
	conn   *grpc.ClientConn
}

//NewClient builds and returns a new Client that wraps a GRPC connections\
func NewClient(opts ClientOpts) *Client {
	return &Client{Opts: opts}
}

//Init starts dials the grpc connection
func (c *Client) Init() (err error) {
	c.conn, err = NewGRPCClient(c.Opts)
	if err != nil {
		return err
	}
	c.Client = request.NewPlainSQLClient(c.conn)
	return nil
}

//Echo call to get response from server
func (c *Client) Echo(req request.Request) (resp *request.Request, err error) {
	log.Println("Client conn: ", c.Client)
	resp, err = c.Client.Echo(context.Background(), &req)
	return resp, err
}

//Close connection to grpc client
func (c *Client) Close() error {
	return c.conn.Close()
}

// ClientOpts to add to your flags or env config
type ClientOpts struct {
	ServerHost string `env:"companysettings.host"       short:"S" long:"server"      default:"127.0.0.1:8091" description:"URL of the PlainSQLD server"`
	SkipVerify bool   `env:"companysettings.skipVerify" short:"k" long:"skip-verify" description:"Skip hostname verification"`
}

// NewGRPCClient that connects to the configurator server
func NewGRPCClient(opts ClientOpts) (conn *grpc.ClientConn, err error) {
	t, err := createTLSConfig(opts.SkipVerify)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create TLS config")
	}

	creds := credentials.NewTLS(t)
	conn, err = grpc.Dial(opts.ServerHost,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to dial server [%s]", opts.ServerHost)
	}
	return conn, nil
}

func createTLSConfig(skipVerify bool) (config *tls.Config, err error) {
	config = &tls.Config{
		InsecureSkipVerify: skipVerify,
	}
	return config, nil
}

// IsNotFound returns true if the given error has an IsNotFound function that also returns true
func IsNotFound(err error) bool {
	sErr, ok := status.FromError(err)
	if !ok {
		return false
	}
	return sErr.Code() == codes.NotFound
}
