package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	port := flag.Int("port", 443, "the port on which to run keewee")
	host := flag.String("host", "", "the host for which to retrieve certificates, e.g. example.com")
	flag.Parse()

	if *host == "" {
		fmt.Println("No host speficied")
		os.Exit(1)
	}

	fmt.Printf("starting keewee on port %d with host %s", *port, *host)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("/static")))

	m := autocert.Manager{
		Cache:      autocert.DirCache("/cache"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(*host),
	}

	s := &http.Server{
		Addr:      fmt.Sprintf(":%d", *port),
		Handler:   mux,
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}

	s.ListenAndServeTLS("", "")
}
