package main

import (
    "log"
    "net"
    "net/http"
    "os"
    "fmt"
)

type DNSValidator struct {
    expectedIP string
}

func main() {
    validator, err := NewDNSValidator()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Config - ACME_EMAIL: %s", os.Getenv("ACME_EMAIL"))
    log.Printf("DNS Validator started. Server IP: %s", validator.expectedIP)
    http.ListenAndServe(":8080", validator)
}

func NewDNSValidator() (*DNSValidator, error) {
    ip, err := getServerIP()
    if err != nil {
        return nil, err
    }
    return &DNSValidator{expectedIP: ip}, nil
}

func (v *DNSValidator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    log.Printf("Middleware called for request: %s %s", req.Method, req.Host)
    host := req.Host

    ips, err := net.LookupHost(host)
    if err != nil || len(ips) == 0 {
        log.Printf("DNS check failed - not resolved: %s (error: %v)", host, err)
        http.Error(rw, "DNS not configured", http.StatusServiceUnavailable)
        return
    }

    if !contains(ips, v.expectedIP) {
        log.Printf("DNS check failed - wrong IP: %s (found IPs: %v, expected: %s)", host, ips, v.expectedIP)
        http.Error(rw, "DNS misconfigured", http.StatusServiceUnavailable)
        return
    }

    log.Printf("DNS check passed: %s -> %s", host, v.expectedIP)
    rw.WriteHeader(http.StatusOK)
}

func getServerIP() (string, error) {
    if ip := os.Getenv("SERVER_IP"); ip != "" {
        return ip, nil
    }
    return "", fmt.Errorf("SERVER_IP environment variable not set")
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
} 