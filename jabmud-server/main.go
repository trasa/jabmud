package main

import (
	"flag"
	"github.com/trasa/jabmud/jab"
)

func main() {
	jabberHostPtr := flag.String("jabberhost", "192.168.99.100", "ejabber host ip address")
	jabberPortPtr := flag.String("jabberport", "5275", "ejabber host port")
	flag.Parse()

	go connectHttpServer()
	jab.ConnectComponent(*jabberHostPtr, *jabberPortPtr)
}
