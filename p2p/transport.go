package p2p

// Peer is an interface that represents the remote node
type Peer interface {}

// Trabsport is anything that handles the communication between nodes
type Transport interface {
	ListenAndAccept() error
}
