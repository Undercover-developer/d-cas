package p2p

//represents a remote node
type Peer interface {
}

//Anything that handle communication between nodes
type Transport interface {
	listenAndAccept()
}
