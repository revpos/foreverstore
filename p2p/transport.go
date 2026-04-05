// Package p2p
package p2p

// import "net"

// Peer is an interface that represents the reomte node.
type Peer interface {
	// net.Conn
	Close() error
}

// Transport is anything that handles the communincation
// between the nodes in the network. This can be of the
// form (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
