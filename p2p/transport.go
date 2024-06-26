package p2p

// Peer is an interface that represent the remote node.
type Peer interface {
	Close()
}

// Transport is anything that handles the communication
// betwwen the nodes in the network. This can be of the
// form (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
