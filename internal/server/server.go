package server

import (
    "fmt"
    "net/http"
    "time"
    "github.com/A-Random-Person-From-Earth/go-camp/internal/greet"
    "github.com/A-Random-Person-From-Earth/go-camp/internal/config"
)


func Start(cfg config.Config) error {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        if name == "" {
            name = cfg.Greeting
        }
        
        message := greet.Greet(name)
        
        _, err := fmt.Fprint(w, message)
        if err != nil {
            http.Error(w, "Failed to write response", http.StatusInternalServerError)
            return
        }
    })


 http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, "OK")
    })
    
    fmt.Printf("Starting server on %s\n", cfg.Port)
    
    server := &http.Server{
        Addr:         cfg.Port,
        ReadTimeout:  cfg.Timeout,
        WriteTimeout: cfg.Timeout,
    }
    
    return server.ListenAndServe()
}
