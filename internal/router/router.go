// package router

// import (
//     "net/http"
//     "go-api-gateway/internal/config"
//     "go-api-gateway/internal/proxy"
// )

// func SetupRouter() *http.ServeMux {
//     mux := http.NewServeMux()
//     services := config.LoadServices()

//     for _, s := range services {
//         mux.HandleFunc(s.Prefix+"/", proxy.NewReverseProxy(s.Target, s.Prefix))
//     }

//     // Health check
//     mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
//         w.WriteHeader(http.StatusOK)
//         w.Write([]byte("OK"))
//     })

//     return mux
// }


package router

import (
    "net/http"
    "go-api-gateway/internal/config"
    "go-api-gateway/internal/proxy"
)

func SetupRouter() *http.ServeMux {
    mux := http.NewServeMux()
    services := config.LoadServices()

    for _, s := range services {
        mux.HandleFunc(s.Prefix+"/", proxy.NewReverseProxy(s.Target, s.Prefix, s.BasePath))
    }

    // Health check endpoint
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    return mux
}
