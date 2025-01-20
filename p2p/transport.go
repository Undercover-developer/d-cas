package p2p

//represents a remote node
type Peer interface {
	Close() error
}

//Anything that handle communication between nodes
type Transport interface {
	listenAndAccept() error
	Consume() <-chan RPC
}
