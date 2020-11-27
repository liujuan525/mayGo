package main

import (
    "fmt"
    "log"
    "mayGo/routers"
    "net/http"
    
    "mayGo/pkg/setting"
)

func main() {
    router := routers.InitRouter()
    
    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }
    
    if err := s.ListenAndServe(); err != nil {
        log.Println(err)
    }
}
