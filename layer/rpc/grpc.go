package rpc

import (
	"context"
	"net"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type GRPC struct {
	host     string
	port     int
	listener net.Listener
	server   *grpc.Server
	client   *grpc.ClientConn
}

func (g *GRPC) CreateServer() error {
	var err error
	g.listener, err = net.Listen("tcp", g.get_address())
	g.server = grpc.NewServer()
	return err
}

func (g *GRPC) OpenConnection(ctx context.Context) error {
	g.client = grpc.DialContext(ctx, g.get_address(), grpc.WithInsecure())
}

func (g *GRPC) SetPort(_port_ int) {
	g.port = _port_
}

func (g *GRPC) SetHost(_host_ string) {
	g.host = _host_
}

func (g *GRPC) GetPort() string {
	return g.port
}

func (g *GRPC) GetHost() string {
	return g.host
}

func (g *GRPC) get_address() string {
	var concatenate_string strings.Builder

	concatenate_string.WriteString(g.host)
	concatenate_string.WriteString(":")
	concatenate_string.WriteString(strconv.Itoa(g.port))
	return concatenate_string.String()
}
