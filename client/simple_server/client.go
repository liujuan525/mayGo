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
    "os"
)

const PORT = "9001"

func main() {
    dir, _ := os.Getwd()
    cert, err := tls.LoadX509KeyPair(dir + "/conf/client/client.pem", dir + "/conf/client/client.key")
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
        ServerName:   "mayGo",
        RootCAs:      certPool,
    })
    
    
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
