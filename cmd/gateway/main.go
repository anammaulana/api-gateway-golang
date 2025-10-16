package main

import (
    "log"
    "net/http"
    "go-api-gateway/internal/router"
	 // swagger
    _ "go-api-gateway/docs"
    httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go API Gateway
// @version 1.0
// @description API Gateway for Quarkus microservices (billing, engine, printing)
// @host localhost:9099
// @BasePath /
// @schemes http
func main() {
    mux := router.SetupRouter()
  // Swagger route
    mux.Handle("/swagger/", httpSwagger.WrapHandler)

    log.Println("🚀 API Gateway running on http://localhost:9099")
    log.Println("📘 Swagger available at http://localhost:9099/swagger/index.html")
    if err := http.ListenAndServe(":9099", mux); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
