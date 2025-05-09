package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/hoppscotch/proxyscotch/libproxy"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// hostPtr := flag.String("host", "localhost:9159", "the hostname that the server should listen on.")
	// tokenPtr := flag.String("token", "", "the Proxy Access Token used to restrict access to the server.")
	// allowedOriginsPtr := flag.String("allowed-origins", "*", "a comma separated list of allowed origins.")
	// bannedOutputsPtr := flag.String("banned-outputs", "", "a comma separated list of banned outputs.")
	// bannedDestsPtr := flag.String("banned-dests", "", "a comma separated list of banned proxy destinations.")

	// flag.Parse()

	finished := make(chan bool)
	//libproxy.Initialize(*tokenPtr, *hostPtr, *allowedOriginsPtr, *bannedOutputsPtr, *bannedDestsPtr, onProxyStateChangeServer, false, finished)
	libproxy.Initialize("", "0.0.0.0:9159", "*", "", "", onProxyStateChangeServer, false, finished)

	<-finished
}

func onProxyStateChangeServer(status string, isListening bool) {
	log.Printf("[ready=%v] %s", isListening, status)
}
