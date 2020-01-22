// Package hubserver provides functionality akin to a DNS server for clients to
// find content hosts.
package hubserver

type hostRecord struct {
	// addr is the address of the host providing this record.
	addr string

	// hosts public key
	pubkey []byte

	shares []share
}

type share struct {
	// the files name
	filename string

	// the files sha512 hash
	filehash []byte
}
