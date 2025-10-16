package config

import "net/url"

// type Service struct {
//     Prefix string
//     Target *url.URL
// }

type Service struct {
    Prefix   string   // Prefix di Gateway
    Target   *url.URL // Host tujuan (tanpa path)
    BasePath string   // Base path di backend service
}


func LoadServices() []Service {
    billing, _ := url.Parse("http://localhost:8002")
    engine, _ := url.Parse("http://localhost:8001/")
    printing, _ := url.Parse("http://localhost:8003")

    // return []Service{
    //     {Prefix: "/billing", Target: billing},
    //     {Prefix: "/engine", Target: engine},
    //     {Prefix: "/printing", Target: printing},
    // }

	 return []Service{
        {Prefix: "/api/dlpd", Target: billing, BasePath: "/api/dlpd"},
        {Prefix: "/api/flow", Target: engine, BasePath: "/api/flow"},
        {Prefix: "/printing", Target: printing, BasePath: "/api/print"},
    }
}
