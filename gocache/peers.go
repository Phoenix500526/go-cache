package gocache

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
type PeerPicker interface {
	// PickPeer is designed to choose the relevant PeerGetter
	// according to the key parameter.
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer
type PeerGetter interface {
	// Get can search the cache value from the relevant group
	Get(group string, key string) ([]byte, error)
}
