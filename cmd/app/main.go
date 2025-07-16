package main

import (
    "fmt"
    "github.com/A-Random-Person-From-Earth/go-camp/internal/config"  
    "github.com/A-Random-Person-From-Earth/go-camp/internal/server" 
)

func main() {
    
    cfg, err := config.Load()
    if err != nil {
        fmt.Printf("Failed to load configuration: %v\n", err)
        return
    }
    
    
    err = server.Start(cfg)
    if err != nil {
        fmt.Printf("Server failed to start: %v\n", err)
    }
}
