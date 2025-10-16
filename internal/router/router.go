package router

import (
	"net/http"

	"go-api-gateway/internal/config"
	"go-api-gateway/internal/proxy"
)

// SetupRouter inisialisasi route utama gateway
func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// 🩺 Health check endpoint
	mux.HandleFunc("/health", HealthHandler)

	// 🔁 Proxy ke service Quarkus
	services := config.LoadServices()
	for _, s := range services {
		mux.Handle(s.Prefix+"/", proxy.NewReverseProxy(s.Target, s.Prefix, s.BasePath))
	}

	return mux
}

// @Summary Health Check
// @Description Check if the API Gateway is running
// @Tags System
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
