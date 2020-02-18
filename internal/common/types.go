// package common coontains functionality that is common to all parts of the
// system.
package common

type Record interface {
	// Addr returns the address of the content provider.
	Addr() string

	// Public key of the content provider.
	PubKey() []byte

	// AllShares returns a slice of all the available Shares (content).
	AllShares() []Share

	// Valid reports if the Record is still valid by making sure it is not
	// past it's day to die (which is like time to live, but an absolute
	// time and date).
	Valid() bool
}

// Share is a shareable resource; typically a filesystem file.
type Share interface {
	// FileName returns the file name of the share.
	FileName() string

	// FileHash returns the sha512 hash of the share.
	FileHash() []byte

	// Size returns the file size.
	Size() int
}
