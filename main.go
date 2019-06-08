package main

import (
	"flag"
	"log"
	"sealion/interfaces/router"

	"github.com/sirupsen/logrus"
)

func main() {
	port := flag.Int("port", 8080, "port to listen http request")
	debugEnabled := flag.Bool("debug", false, "enable debug output in logs")

	th, err := initialize()
	if err != nil {
		log.Fatal(err)
	}
	if *debugEnabled {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if err := router.Run(*port, th); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
