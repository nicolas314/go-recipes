// Demonstrates: HTTPS server listening on port 4443
// Generating (self-signed) server key and certificate is left to the
// reader
// This server will require incoming user to authenticate using a
// client-side certificate issued by any authority. It will respond by
// printing the client certificate Common Name.
package main

import (
    "fmt"
    "net/http"
    "crypto/tls"
)

func Hello(w http.ResponseWriter, req * http.Request) {
    w.Header().Set("Content-type", "text/plain")
    client_cert := req.TLS.PeerCertificates[0]
    fmt.Fprintf(w, "You are: %s\n", client_cert.Subject.CommonName)
}

func main() {
    http.HandleFunc("/", Hello)
    t := tls.Config {
            ClientAuth: tls.RequireAnyClientCert,
         }
    s := &http.Server {
            Addr:       ":4443",
            TLSConfig:  &t,
         }
    err := s.ListenAndServeTLS("server.crt", "server.key")
    if err!=nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Listening on 4443...")
}
