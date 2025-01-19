package main

import (
	"fmt"
	"log"

	"github.com/Undercover-developer/ipfs/p2p"
)

func main() {
	fmt.Println("Hello world")
	tr := p2p.NewTCPTransport(":4000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}
