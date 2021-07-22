package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

func main() {
	server.RunGRPCServer(func(server *grpc.Server) {
		svc := ports.NewGrpcServer(application)
		trainer.RegisterTrainerServiceServer(server, svc)
	})
}
