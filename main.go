package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "golang.org/x/net/proxy"
)

func main() {
    // Get the port from the environment variable, default to "8080" if not set
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Set up the listener on the specified port
    listenAddr := fmt.Sprintf(":%s", port)
    listener, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatalf("Error setting up listener: %v", err)
    }
    defer listener.Close()

    fmt.Printf("SOCKS5 proxy server is listening on %s\n", listenAddr)

    // Accept and handle incoming connections
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Error accepting connection: %v", err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Create a SOCKS5 proxy listener
    socks5 := proxy.New SOCKS5("tcp", "localhost:1080", nil, nil)
    if err != nil {
        log.Printf("Error creating proxy: %v", err)
        return
    }
    
    err := socks5.HandleRequest(conn)
    if err != nil {
        log.Printf("Error handling request: %v", err)
    }
}
