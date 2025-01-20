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
		Decoder:          p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(opts)

	//testing rpc channel
	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Println(msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}
