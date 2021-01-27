package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "log"
    "mayGo/pkg/gtls"
    pb "mayGo/proto"
    "os"
    "time"
)

const PORT = "9001"

func main() {
    dir, _ := os.Getwd()
    tlsClient := gtls.Client{
        ServerName: "mayGo",
        CaFile:     dir + "/conf/ca.pem",
        CertFile:   dir + "/conf/client/client.pem",
        KeyFile:    dir + "/conf/client/client.key",
    }
    
    c, err := tlsClient.GetCredentialsByCA()
    if err != nil {
        log.Fatalf("GetTLSCredentialsByCA err: %v", err)
    }
    
    conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
    if err != nil {
        log.Fatalf("grpc.Dial err: %v", err)
    }
    defer conn.Close()
    
    ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5*time.Second)))
    defer cancel()
    
    client := pb.NewSearchServiceClient(conn)
    resp, err := client.Search(ctx, &pb.SearchRequest{
        Request: "gRPC",
    })
    if err != nil {
        statusErr, ok := status.FromError(err)
        if ok {
            if statusErr.Code() == codes.DeadlineExceeded {
                log.Fatalln("client.Search err: deadline")
            }
        }
        
        log.Fatalf("client.Search err: %v", err)
    }
    
    log.Printf("resp: %s", resp.GetResponse())
}
