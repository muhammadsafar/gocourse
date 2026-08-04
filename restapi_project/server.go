package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming request for /orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming request for /products")
	})

	port := 3000

	//Load the TLS certificate and key files
	certFile := "../cert.pem"
	keyFile := "../key.pem"

	//Configure TLS

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//Create a custom server with TLS configuration
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil, // Use the default ServeMux
		TLSConfig: tlsConfig,
	}

	//Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	//Start the HTTPS server with TLS
	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}

	//HTTP 1.1 Server Without TLS
	// err :=  http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("Error starting server:", err)
	// }

}
