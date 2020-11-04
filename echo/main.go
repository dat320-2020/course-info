// A simple TCP client and server
package main

import "flag"

func main() {
	var (
		server   = flag.Bool("server", false, "Start echo server if true; otherwise start echo client")
		endpoint = flag.String("endpoint", "localhost:12101", "Endpoint on which the server runs or to which the client connects")
	)
	flag.Parse()
	if *server {
		serverLoop(*endpoint)
	} else {
		clientLoop(*endpoint)
	}
}
