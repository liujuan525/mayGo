package main

import (
    "context"
    "github.com/grpc-ecosystem/go-grpc-middleware"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "log"
    "mayGo/interceptors"
    "mayGo/pkg/gtls"
    pb "mayGo/proto"
    "net"
    "os"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
    if ctx.Err() == context.Canceled {
        return nil, status.Errorf(codes.Canceled, "searchService.Search canceled")
    }
    
    return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
    dir, _ := os.Getwd()
    tlsServer := gtls.Server{
        CaFile:   dir + "/conf/ca.pem",
        CertFile: dir + "/conf/server/server.pem",
        KeyFile:  dir + "/conf/server/server.key",
    }
    c, err := tlsServer.GetCredentialsByCA()
    if err != nil {
        log.Fatalf("GetTLSCredentialsByCA err: %v", err)
    }
    
    opts := []grpc.ServerOption{
        grpc.Creds(c),
        grpc_middleware.WithUnaryServerChain(
            interceptors.RecoveryInterceptor,
            interceptors.LoggingInterceptor,
        ),
    }
    
    server := grpc.NewServer(opts...)
    pb.RegisterSearchServiceServer(server, &SearchService{})
    
    lis, err := net.Listen("tcp", ":"+PORT)
    if err != nil {
        log.Fatalf("net.Listen err: %v", err)
    }
    
    server.Serve(lis)
}
