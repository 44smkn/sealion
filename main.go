package main

import (
	"log"
	"flag"
	"sealion/interfaces/router"
)

func main() {
	port := flag.Int("port", 8080, "port to listen http request")
	if err := router.Run(*port); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
