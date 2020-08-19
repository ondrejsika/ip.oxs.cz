package handler

import (
	"fmt"
	"net/http"
	"strings"
)

// Handler for Vercel or Go server
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, GetIP(r))
	fmt.Fprintf(w, "\n")
}

// GetIP return IP address from X-Real-Ip, X-Forwarded-For or RemoteAddr
func GetIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		return r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")]
	}
	return IPAddress
}
