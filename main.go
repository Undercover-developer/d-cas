package main

import (
	"fmt"
	"log"

	"github.com/Undercover-developer/ipfs/p2p"
)

func OnPeer(peer p2p.Peer) error {
	//testing peering logic
	fmt.Println("peering...")
	return nil
}

func main() {
	fmt.Println("Starting server ...")
	opts := p2p.TCPTransportOpts{
		ListenAddress:    ":4000",
		HandshakeHandler: p2p.NOPHandshakeFunc,
		Decoder:          p2p.DefaultDecoder{},
		OnPeer:           OnPeer,
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
