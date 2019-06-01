package main

import (
	"log"
	"flag"
	"sealion/interfaces/router"
	"github.com/sirupsen/logrus"
)

func main() {
	port := flag.Int("port", 8080, "port to listen http request")
	debugEnabled := flag.Bool("debug", false, "enable debug output in logs")
	if *debugEnabled {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if err := router.Run(*port); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
