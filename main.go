package main

import (
	"flag"
	"sealion/interfaces/router"
)

func main() {
	port := flag.Int("port", 8080, "port to listen http request")
	router.Run(*port)
}
