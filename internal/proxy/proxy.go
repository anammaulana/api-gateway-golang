

package proxy

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "path"
    "strings"
)

func NewReverseProxy(target *url.URL, prefix string, basePath string) http.HandlerFunc {
    proxy := httputil.NewSingleHostReverseProxy(target)

    proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
        log.Printf("❌ Proxy error to %s: %v", target, err)
        http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
    }

    return func(w http.ResponseWriter, r *http.Request) {
        // Bersihkan prefix dari URL client
        cleanPath := strings.TrimPrefix(r.URL.Path, prefix)

        // Gabungkan basePath (di service) + sisa path client
        r.URL.Path = path.Join(basePath, cleanPath)

        log.Printf("➡️  %s %s → %s%s", r.Method, prefix, target, r.URL.Path)
        proxy.ServeHTTP(w, r)
    }
}
