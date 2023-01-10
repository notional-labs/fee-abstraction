package types

// NewCapability returns a reference to a new Capability to be used as an
// actual capability.
func NewCapability(index uint64) *Capability {
	return &Capability{Index: index}
}
