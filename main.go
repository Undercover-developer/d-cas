package main

import (
	"fmt"
	"log"

	"github.com/Undercover-developer/ipfs/p2p"
)

func main() {
	fmt.Println("Starting server ...")
	opts := p2p.TCPTransportOpts{
		ListenAddress:    ":4000",
		HandshakeHandler: p2p.NOPHandshakeFunc,
	}
	tr := p2p.NewTCPTransport(opts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}
