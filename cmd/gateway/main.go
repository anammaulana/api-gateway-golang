package main

import (
    "log"
    "net/http"
    "go-api-gateway/internal/router"
)

func main() {
    mux := router.SetupRouter()

    log.Println("🚀 API Gateway running on http://localhost:9099")
    if err := http.ListenAndServe(":9099", mux); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
