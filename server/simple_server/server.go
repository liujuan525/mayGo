package main

import (
    "context"
    "crypto/tls"
    "crypto/x509"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "io/ioutil"
    "log"
    pb "mayGo/proto"
    "net"
    "os"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
    return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
    dir, _ := os.Getwd()
    cert, err := tls.LoadX509KeyPair(dir + "/conf/server/server.pem", dir + "/conf/server/server.key")
    if err != nil {
        log.Fatalf("tls.LoadX509KeyPair err: %v", err)
    }
    
    certPool := x509.NewCertPool()
    ca, err := ioutil.ReadFile(dir + "/conf/ca.pem")
    if err != nil {
        log.Fatalf("ioutil.ReadFile err: %v", err)
    }
    
    if ok := certPool.AppendCertsFromPEM(ca); !ok {
        log.Fatalf("certPool.AppendCertsFromPEM err")
    }
    
    c := credentials.NewTLS(&tls.Config{
        Certificates: []tls.Certificate{cert},
        ClientAuth:   tls.RequireAndVerifyClientCert,
        ClientCAs:    certPool,
    })
    
    server := grpc.NewServer(grpc.Creds(c))
    pb.RegisterSearchServiceServer(server, &SearchService{})
    
    lis, err := net.Listen("tcp", ":"+PORT)
    if err != nil {
        log.Fatalf("net.Listen err: %v", err)
    }
    
    server.Serve(lis)
}
