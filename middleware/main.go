package main

import (
    "log"
    "net"
    "net/http"
    "io"
)

type DNSValidator struct {
    expectedIP string
}

func main() {
    validator, err := NewDNSValidator()
    if err != nil {
        log.Fatal(err)
    }

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
    host := req.Host

    ips, err := net.LookupHost(host)
    if err != nil || len(ips) == 0 {
        log.Printf("DNS not resolved for %s", host)
        http.Error(rw, "DNS not configured", http.StatusServiceUnavailable)
        return
    }

    if !contains(ips, v.expectedIP) {
        log.Printf("DNS does not point to server IP for %s", host)
        http.Error(rw, "DNS misconfigured", http.StatusServiceUnavailable)
        return
    }

    log.Printf("DNS validated for %s", host)
    rw.WriteHeader(http.StatusOK)
}

func getServerIP() (string, error) {
    resp, err := http.Get("https://api.ipify.org")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    
    ip, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    
    return string(ip), nil
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
} 