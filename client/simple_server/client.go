package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "log"
    pb "mayGo/proto"
    "os"
)

const PORT = "9001"

func main() {
    dir, _ := os.Getwd()
    c, err := credentials.NewClientTLSFromFile(dir + "/conf/server.pem", "mayGo")
    if err != nil {
        log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
    }
    
    conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
    if err != nil {
        log.Fatalf("grpc.Dial err: %v", err)
    }
    defer conn.Close()
    
    client := pb.NewSearchServiceClient(conn)
    resp, err := client.Search(context.Background(), &pb.SearchRequest{
        Request: "gRPC",
    })
    if err != nil {
        log.Fatalf("client.Search err: %v", err)
    }
    
    log.Printf("resp: %s", resp.GetResponse())
}
