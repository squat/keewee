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
	insecure := flag.Bool("insecure", false, "whether or not to disable TLS")
	flag.Parse()

	if *host == "" && !*insecure {
		fmt.Println("No host speficied")
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("/static")))

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: mux,
	}

	fmt.Printf("starting keewee on port %d", *port)
	if *insecure {
		s.ListenAndServe()
	} else {
		fmt.Printf("using host %s", *host)
		m := autocert.Manager{
			Cache:      autocert.DirCache("/cache"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(*host),
		}
		s.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
		s.ListenAndServeTLS("", "")
	}
}
